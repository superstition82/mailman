package store

import (
	"database/sql"
	"pocketmail/server/profile"
)

// Store provides database access to all raw objects.
type Store struct {
	Profile *profile.Profile
	db      *sql.DB
}

// New creates a new instance of Store.
func New(db *sql.DB, profile *profile.Profile) *Store {
	return &Store{
		Profile: profile,
		db:      db,
	}
}
