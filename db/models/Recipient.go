package models

import "gorm.io/gorm"

type Status string

const (
	PENDING Status = "PENDING"
	VALID   Status = "VALID"
	INVALID Status = "INVALID"
)

type Recipient struct {
	gorm.Model

	Email  string
	Status Status
}
