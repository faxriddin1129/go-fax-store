package routes

import (
	"github.com/gin-gonic/gin"
	"microservice/internal/controllers"
)

func MainRoutes(r *gin.Engine) {

	r.GET("/", controllers.Welcome)

}
