package controllers

import (
	"github.com/gin-gonic/gin"
	"gofax-store/pkg/utils"
)

func Welcome(c *gin.Context) {
	utils.RespondJson(c, nil, 200, "Welcome to the GOFAX STORE 2")
}
