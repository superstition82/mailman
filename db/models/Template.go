package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Template struct {
	gorm.Model
	ID uint64 `gorm:"primary_key"`

	Subject string
	Body    string
	Date    time.Time
}
