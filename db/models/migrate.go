package models

import "gorm.io/gorm"

// Migrate : migrate models
func Migrate(db *gorm.DB) {
	db.AutoMigrate(&SMTP{})
	db.AutoMigrate(&Template{})
	db.AutoMigrate(&Recipient{})
}
