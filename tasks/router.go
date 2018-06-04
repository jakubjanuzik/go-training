package tasks

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/location"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func getLocationUrl(context *gin.Context, id uint) string {
	return fmt.Sprintf("%v%v/%v", location.Get(context), context.Request.URL, id)
}

func createTask(context *gin.Context) {
	var task Task
	context.Bind(&task)
	db, _ := context.MustGet("db").(*gorm.DB)
	db.Save(&task)
	context.Header("Location", getLocationUrl(context, task.ID))
	context.JSON(http.StatusCreated, gin.H{"resourceId": task.ID})
}

func getTasks(context *gin.Context) {
	var tasks []Task
	db, _ := context.MustGet("db").(*gorm.DB)

	db.Find(&tasks)
	context.JSON(http.StatusOK, tasks)
}

func getTask(context *gin.Context) {
	context.JSON(http.StatusOK, context.MustGet("task").(Task))
}

func getTaskMiddleWare(db *gorm.DB) gin.HandlerFunc {
	return func(context *gin.Context) {
		var task Task
		id := context.Param("id")
		db.First(&task, id)
		if task.ID == 0 {
			context.AbortWithStatus(http.StatusNotFound)
		} else {
			context.Set("task", task)
			context.Next()
			// context.JSON(http.StatusOK, task)
		}
	}
}

func updateTask(context *gin.Context) {
	db, _ := context.MustGet("db").(*gorm.DB)
	task := context.MustGet("task").(Task)

	context.Bind(&task)
	db.Save(&task)
	context.JSON(http.StatusOK, gin.H{"resourceId": task.ID})
}

func deleteTask(context *gin.Context) {
	task := context.MustGet("task").(Task)
	db, _ := context.MustGet("db").(*gorm.DB)

	db.Delete(&task)
	context.JSON(http.StatusNoContent, gin.H{"resourceId": task.ID})
}

func Start() {
	router, db := initialize()

	api := router.Group("/api/v1/")
	api.POST("tasks", createTask)
	api.GET("tasks", getTasks)
	api.GET("tasks/:id", getTaskMiddleWare(db), getTask)
	api.PUT("tasks/:id", getTaskMiddleWare(db), updateTask)
	api.DELETE("tasks/:id", getTaskMiddleWare(db), deleteTask)
	router.Run()
}
