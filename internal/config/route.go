package config

import (
	"github.com/gin-gonic/gin"
	"gofax-store/routes"
)

func RegisterRoutes(r *gin.Engine) {
	routes.MainRoutes(r)
}
