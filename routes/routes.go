package routes

import (
	"course-go/controllers"

	"github.com/gin-gonic/gin"
)

func Serve(r *gin.Engine) {

	//Create group routes
	articleGroup := r.Group("/api/v1/articles")

	// Define controllers
	articleController := controllers.Articles{}

	articleGroup.GET("/", articleController.FindAll)
	articleGroup.GET("/:id", articleController.FindOne)
	articleGroup.POST("", articleController.Create)
}
