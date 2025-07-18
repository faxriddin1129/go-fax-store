package models

import (
	"gorm.io/gorm"
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
