package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"microservice/internal/models"
	"microservice/pkg/env"
	"microservice/pkg/utils"
	"path/filepath"
	"strings"
	"time"
)

func StoreFile(c *gin.Context) {

	file, err := c.FormFile("file")
	if err != nil {
		utils.RespondJson(c, nil, 422, "File is required")
		return
	}

	storageDir, err := utils.FileGetPath()
	if err != nil {
		utils.RespondJson(c, err.Error(), 500, "Error retrieving the path.")
		return
	}

	const maxSize = 10 * 1024 * 1024
	allowedExtensions := []string{"jpg", "jpeg", "pdf", "webp", "png", "xlsx", "svg", "doc", "docx"}

	if file.Size > maxSize {
		utils.RespondJson(c, nil, 413, "File size must not exceed 10MB.")
		return
	}

	sizeInMB := float64(file.Size) / (1024 * 1024)
	formattedSize := fmt.Sprintf("%.2f MB", sizeInMB)

	ext := strings.ToLower(filepath.Ext(file.Filename))
	ext = strings.TrimPrefix(ext, ".")
	allowed := utils.InArray(allowedExtensions, ext)
	if !allowed {
		utils.RespondJson(c, nil, 415, "Unsupported file format")
		return
	}

	fullPath := filepath.Join(storageDir, file.Filename)

	err = c.SaveUploadedFile(file, fullPath)
	if err != nil {
		utils.RespondJson(c, err.Error(), 500, "Failed to save the file")
		return
	}

	now := time.Now()
	year := now.Format("2006")
	month := now.Format("01")
	day := now.Format("02")
	name := file.Filename

	loadUrl := env.GetEnv("BASE_URL")
	url := loadUrl + "/" + fullPath

	ip := c.ClientIP()
	device := c.GetHeader("User-Agent")
	newFile := models.Files{
		Year:   year,
		Month:  month,
		Day:    day,
		Ip:     ip,
		Device: device,
		Name:   name,
		Url:    url,
		Size:   formattedSize,
		Format: ext,
	}

	err = utils.DB.Create(&newFile).Error
	if err != nil {
		utils.RespondJson(c, nil, 500, "Internal server error. File failed save")
		return
	}

	utils.RespondJson(c, newFile, 200, "Successfully saved")
}
