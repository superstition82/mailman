package db

import (
	"context"
	"fmt"
	"probemail/db/models"
	"probemail/server/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type DB struct {
	DBInstance *gorm.DB
	config     *config.Config
}

func NewDB(config *config.Config) *DB {
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
	sqliteDB, err := gorm.Open("sqlite3", db.config.DSN+"?cache=shared&_foreign_keys=0&_journal_mode=WAL")
	if err != nil {
		return fmt.Errorf("failed to open db with dsn: %s, err: %w", db.config.DSN, err)
	}
	sqliteDB.LogMode(true)
	db.DBInstance = sqliteDB

	models.Migrate(sqliteDB)

	return nil
}
