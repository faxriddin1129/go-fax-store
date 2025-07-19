package controllers

import (
	"github.com/gin-gonic/gin"
	"microservice/internal/models"
	"microservice/pkg/utils"
)

func GetAllFiles(c *gin.Context) {

	year := c.Query("year")
	month := c.Query("month")
	day := c.Query("day")

	data := models.GetAllFiles(year, month, day)

	utils.RespondJson(c, data, 200, "Success")
	return
}
