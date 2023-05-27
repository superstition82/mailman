package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

// SMTP : SMTP Model
type SMTP struct {
	gorm.Model
	ID   uint64 `gorm:"primary_key"` // auto increment uint default true
	Date time.Time

	Host     string
	Port     string
	Email    string
	Password string
}
