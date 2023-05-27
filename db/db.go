package db

import (
	"context"
	"fmt"
	"probemail/db/models"
	"probemail/util"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DB struct {
	DBInstance *gorm.DB
	config     *util.Config
}

func NewDB(config *util.Config) *DB {
	db := &DB{
		config: config,
	}
	return db
}

func (db *DB) Open(context context.Context) error {
	// Ensure a DSN is set before attempting to open the database.
	if db.config.DSN == "" {
		return fmt.Errorf("DSN required")
	}

	// Connect to the database without foreign_key.
	sqliteDB, err := gorm.Open(sqlite.Open(db.config.DSN+"?cache=shared&_foreign_keys=0&_journal_mode=WAL"), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to open db with dsn: %s, err: %w", db.config.DSN, err)
	}
	db.DBInstance = sqliteDB
	models.Migrate(sqliteDB)

	return nil
}
