package models

import "gorm.io/gorm"

type Template struct {
	gorm.Model

	Subject string
	Body    string
}
