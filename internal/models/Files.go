package models

import (
	"gorm.io/gorm"
	"log"
	"microservice/pkg/utils"
)

type Files struct {
	gorm.Model
	Year   string `json:"Year" gorm:"type:varchar(255)"`
	Month  string `json:"Month" gorm:"type:varchar(255)"`
	Day    string `json:"Day" gorm:"type:varchar(255)"`
	Ip     string `json:"Ip" gorm:"type:varchar(255)"`
	Device string `json:"Device"`
	Url    string `json:"Url" gorm:"type:varchar(255)"`
	Slag   string `json:"Slag" gorm:"type:varchar(255)"`
}

func (Files) TableName() string {
	return "files"
}

func GetAllFiles(year, month, day string) []Files {
	var data []Files
	query := utils.DB

	if year != "" {
		query = query.Where("year = ?", year)
	}

	if month != "" {
		query = query.Where("month = ?", month)
	}

	if day != "" {
		query = query.Where("day = ?", day)
	}

	result := query.Find(&data)
	if result.Error != nil {
		log.Printf("Error fetching files: %v", result.Error)
		return nil
	}

	return data
}

type TreeResponse struct {
	Years []YearNode `json:"years"`
}

type YearNode struct {
	Year   string      `json:"year"`
	Months []MonthNode `json:"months"`
}

type MonthNode struct {
	Month string    `json:"month"`
	Days  []DayNode `json:"days"`
}

type DayNode struct {
	Day   string `json:"day"`
	Count int    `json:"count"`
}

type GroupedResult struct {
	Year  string `json:"year"`
	Month string `json:"month"`
	Day   string `json:"day"`
	Count int    `json:"count"`
}

func GetFilesTree() interface{} {

	results := utils.DB.Model(&Files{}).
		Select("year, month, day, COUNT(*) as count").
		Group("year, month, day").
		Order("CAST(year AS UNSIGNED), CAST(month AS UNSIGNED), CAST(day AS UNSIGNED)")

	return results
}
