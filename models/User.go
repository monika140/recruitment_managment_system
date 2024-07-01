package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username        string  `json:"name"`
	Email           string  `json:"email" gorm:"unique"`
	Password        string  `json:"password"`
	Address         string  `json:"address"`
	UserType        string  `json:"user_type"`
	PasswordHash    string  `json:"-"`
	Profile         Profile `json:"profile"`
	ProfileHeadline string  `json:"profile_headline"`
}

// Function to auto-migrate the user model.
// func MigrateDB(db *gorm.DB) {
// 	db.AutoMigrate(&User{})
// }
