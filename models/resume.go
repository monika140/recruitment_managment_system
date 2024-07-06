package models

import "gorm.io/gorm"

type Resume struct {
	gorm.Model
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	JobID uint   `json:"job_id"`
}
