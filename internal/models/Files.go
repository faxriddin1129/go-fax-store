package models

import (
	"gorm.io/gorm"
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
	Name   string `json:"Name" gorm:"type:varchar(255)"`
	Count  int    `json:"Count"`
	Size   string `json:"Size"`
}

func (Files) TableName() string {
	return "files"
}

func GetAllFiles(year, month, day string) interface{} {
	var data []Files
	query := utils.DB

	queryType := ""

	if year == "" && month == "" && day == "" {
		query = query.Select("year, COUNT(DISTINCT month) as count").Group("year").Order("year")
		queryType = "year"
	} else if year != "" && month == "" && day == "" {
		queryType = "month"
		query = query.Where("year = ?", year).Select("month, COUNT(DISTINCT day) as count").Group("month").Order("CAST(month AS INTEGER)")
	} else if year != "" && month != "" && day == "" {
		queryType = "day"
		query = query.Where("year = ? AND month = ?", year, month).Select("day, COUNT(*) as count").Group("day").Order("day::INTEGER")
	} else {
		queryType = "file"
		query = query.Where("year = ? AND month = ? AND day = ?", year, month, day)
	}

	result := query.Find(&data)
	if result.Error != nil {
		return []Files{}
	}

	sendData := make([]map[string]interface{}, 0)

	if queryType == "year" {
		for _, d := range data {
			sendData = append(sendData, map[string]interface{}{
				"Year":  d.Year,
				"Count": d.Count,
				"Type":  "Year",
			})
		}
	}

	if queryType == "month" {
		for _, d := range data {
			sendData = append(sendData, map[string]interface{}{
				"Month": d.Month,
				"Count": d.Count,
				"Type":  "Month",
			})
		}
	}

	if queryType == "day" {
		for _, d := range data {
			sendData = append(sendData, map[string]interface{}{
				"Day":   d.Day,
				"Count": d.Count,
				"Type":  "Day",
			})
		}
	}

	if queryType == "file" {
		for _, d := range data {
			sendData = append(sendData, map[string]interface{}{
				"CreatedAt": d.CreatedAt,
				"UpdatedAt": d.UpdatedAt,
				"Year":      d.Year,
				"Month":     d.Month,
				"Day":       d.Day,
				"Ip":        d.Ip,
				"Device":    d.Device,
				"Url":       d.Url,
				"Name":      d.Name,
				"Count":     d.Count,
				"Size":      d.Size,
				"Type":      "File",
			})
		}
	}

	return sendData
}
