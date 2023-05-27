package db

import (
	"database/sql"
	"probemail/util"
)

// Store provides database access to all raw objects.
type Store struct {
	Config *util.Config
	db     *sql.DB
}

// NewStore creates a new instance of Store.
func NewStore(db *sql.DB, config *util.Config) *Store {
	return &Store{
		Config: config,
		db:     db,
	}
}
