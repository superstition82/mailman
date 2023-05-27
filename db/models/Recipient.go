package models

import (
	"github.com/jinzhu/gorm"
)

type Status string

const (
	PENDING Status = "PENDING"
	VALID   Status = "VALID"
	INVALID Status = "INVALID"
)

type Recipient struct {
	gorm.Model
	ID uint64 `gorm:"primary_key"`

	Email  string
	Status Status
}
