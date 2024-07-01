package models

import (
	"gorm.io/gorm"
)

type Application struct {
	gorm.Model
	JobID  uint `json:"job_id"`
	UserID uint `json:"user_id"`
}
