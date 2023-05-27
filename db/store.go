package db

import (
	"database/sql"
	_config "probemail/server/config"
)

// Store provides database access to all raw objects.
type Store struct {
	Config *_config.Config
	db     *sql.DB
}

// NewStore creates a new instance of Store.
func NewStore(db *sql.DB, config *_config.Config) *Store {
	return &Store{
		Config: config,
		db:     db,
	}
}
