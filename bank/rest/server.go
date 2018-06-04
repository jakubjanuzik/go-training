package rest

import (
	"github.com/gin-contrib/location"
	"github.com/gin-gonic/gin"
	"github.com/jakubjanuzik/bank/account"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	cors "gopkg.in/gin-contrib/cors.v1"

	"github.com/jinzhu/gorm"
)

var Api *gin.RouterGroup

func InitDb() *gorm.DB {
	db, err := gorm.Open("sqlite3", "bank.db")
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&account.Account{})
	return db
}

func DbMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Set("db", db)
		context.Next()
	}
}

func initialize() (*gin.Engine, *gorm.DB) {
	db := InitDb()
	engine := gin.Default()
	engine.Use(cors.Default(), DbMiddleware(db), location.Default())
	return engine, db
}
