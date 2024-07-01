package models

import (
	"gorm.io/gorm"
)

type Profile struct {
	gorm.Model
	UserID     uint   `json:"user_id"`
	ResumeFile string `json:"resume_file"`
	Skills     string `json:"skills"`
	Education  string `json:"education"`
	Experience string `json:"experience"`
	Phone      string `json:"phone"`
}
