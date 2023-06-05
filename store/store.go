package store

import (
	"database/sql"
	"mails/server/config"
)

// Store provides database access to all raw objects.
type Store struct {
	config *config.Config
	db     *sql.DB
}

// New creates a new instance of Store.
func New(db *sql.DB, config *config.Config) *Store {
	return &Store{
		config: config,
		db:     db,
	}
}
