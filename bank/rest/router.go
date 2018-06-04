package rest

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/location"
	"github.com/gin-gonic/gin"
	"github.com/jakubjanuzik/bank/generator"
	"github.com/jakubjanuzik/bank/repository"
	"github.com/jakubjanuzik/bank/service"
	"github.com/jinzhu/gorm"
)

type requestBody struct {
	Number          string `json:"number"`
	Funds           int
	DispositionType string
}

func getLocationUrl(context *gin.Context, id int) string {
	return fmt.Sprintf("%v%v/%v", location.Get(context), context.Request.URL, id)
}

func createAccount(context *gin.Context) {
	service := context.MustGet("service").(service.AccountService)
	number := service.CreateAccount()

	context.JSON(http.StatusCreated, gin.H{"accountNumber": number})
}

func getAccount(context *gin.Context) {
	service := context.MustGet("service").(service.AccountService)
	number := context.Param("number")

	account, _ := service.Repository.GetByNumber(number)
	context.JSON(http.StatusOK, account)
}

func handleDisposition(context *gin.Context) {
	fmt.Printf("INITIAL\n")
	reqBody := requestBody{}
	service := context.MustGet("service").(service.AccountService)

	context.Bind(&reqBody)
	fmt.Printf("%v", reqBody)
	accNumber := reqBody.Number
	funds := reqBody.Funds
	dispositionType := reqBody.DispositionType

	if dispositionType == "deposit" {
		fmt.Printf("DEPOSIT\n")
		service.DepositFunds(accNumber, funds)
		context.JSON(http.StatusOK, gin.H{})
	} else if dispositionType == "withdraw" {
		fmt.Printf("WITHDRAW\n")
		service.WithdrawFunds(accNumber, funds)
		context.JSON(http.StatusOK, gin.H{})
	} else {
		fmt.Printf("UNKNOWN\n")
		context.JSON(http.StatusBadRequest, gin.H{})
	}

}

func accountServiceMiddleware(db *gorm.DB, accountService service.AccountService) gin.HandlerFunc {
	return func(context *gin.Context) {

		context.Set("service", accountService)
		context.Next()
	}
}

func Start() {
	router, db := initialize()
	accRepo := repository.GormAccountRepository{DB: db}
	generator := generator.GormAccountRepository{DB: db}
	userRepo := repository.UserRepository{DB: db}

	accountService := service.AccountService{Repository: &accRepo, Generator: &generator, UserRepository: userRepo}

	api := router.Group("/api/v1/")
	api.POST("accounts", accountServiceMiddleware(db, accountService), createAccount)
	api.GET("accounts/:number", accountServiceMiddleware(db, accountService), getAccount)
	api.POST("dispositions", accountServiceMiddleware(db, accountService), handleDisposition)
	router.Run()
}
