package models

import "gorm.io/gorm"

// SMTP : SMTP Model
type SMTP struct {
	gorm.Model
	Host     string
	Port     string
	Email    string
	Password string
}
