package routes

import (
	"github.com/gin-gonic/gin"
	"gofax-store/internal/controllers"
)

func MainRoutes(r *gin.Engine) {

	r.GET("/", controllers.Welcome)
	r.POST("/store", controllers.StoreFile)
	r.GET("/files", controllers.GetAllFiles)
	r.Static("/storage", "./storage")

}
