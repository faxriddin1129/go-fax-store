package routes

import (
	"github.com/gin-gonic/gin"
	"microservice/internal/controllers"
)

func MainRoutes(r *gin.Engine) {

	r.GET("/", controllers.Welcome)
	r.POST("/store", controllers.StoreFile)
	r.GET("/files", controllers.GetAllFiles)
	r.GET("/tree", controllers.GetTree)

}
