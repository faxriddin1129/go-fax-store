package controllers

import (
	"github.com/gin-gonic/gin"
	"gofax-store/internal/models"
	"gofax-store/pkg/utils"
)

func GetAllFiles(c *gin.Context) {

	year := c.Query("year")
	month := c.Query("month")
	day := c.Query("day")

	data := models.GetAllFiles(year, month, day)

	utils.RespondJson(c, data, 200, "Success")
	return
}
