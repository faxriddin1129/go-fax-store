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

	if year == "" || month == "" || day == "" {
		utils.RespondJson(c, nil, 422, "Year, month and day are required")
		return
	}

	data := models.GetAllFiles(year, month, day)

	utils.RespondJson(c, data, 200, "Success")
	return
}

func GetTree(c *gin.Context) {
	data := models.GetFilesTree()
	utils.RespondJson(c, data, 200, "Files tree structure")
	return
}
