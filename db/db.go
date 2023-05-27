package db

import (
	"context"
	"database/sql"
	"embed"
	"errors"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"

	"probemail/util"
)

//go:embed migration
var migrationFS embed.FS

type DB struct {
	DBInstance *sql.DB
	config     *util.Config
}

func NewDB(config *util.Config) *DB {
	db := &DB{
		config: config,
	}
	return db
}

func (db *DB) Open(ctx context.Context) (err error) {
	// Ensure a DSN is set before attempting to open the database.
	if db.config.DSN == "" {
		return fmt.Errorf("DSN required")
	}

	// Connect to the database without foreign_key.
	sqliteDB, err := sql.Open("sqlite3", db.config.DSN+"?cache=shared&_foreign_keys=0&_journal_mode=WAL")
	if err != nil {
		return fmt.Errorf("failed to open db with dsn: %s, err: %w", db.config.DSN, err)
	}
	db.DBInstance = sqliteDB

	if db.config.Mode == "prod" {
		// If db file not exists, we should migrate the database.

	} else {
		// In non-prod mode, we should always migrate the database.
		if _, err := os.Stat(db.config.DSN); errors.Is(err, os.ErrNotExist) {
			if err := db.applyLatestSchema(ctx); err != nil {
				return fmt.Errorf("failed to apply latest schema: %w", err)
			}
		}
	}

	return nil
}

const (
	latestSchemaFileName = "LATEST__SCHEMA.sql"
)

func (db *DB) applyLatestSchema(ctx context.Context) error {
	schemaMode := "dev"
	if db.config.Mode == "prod" {
		schemaMode = "prod"
	}
	latestSchemaPath := fmt.Sprintf("%s/%s/%s", "migration", schemaMode, latestSchemaFileName)
	buf, err := migrationFS.ReadFile(latestSchemaPath)
	if err != nil {
		return fmt.Errorf("failed to read latest schema %q, error %w", latestSchemaPath, err)
	}
	stmt := string(buf)
	if err := db.execute(ctx, stmt); err != nil {
		return fmt.Errorf("migrate error: statement:%s err=%w", stmt, err)
	}
	return nil
}

// execute runs a single SQL statement within a transaction.
func (db *DB) execute(ctx context.Context, stmt string) error {
	tx, err := db.DBInstance.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if _, err := tx.ExecContext(ctx, stmt); err != nil {
		return fmt.Errorf("failed to execute statement, err: %w", err)
	}

	return tx.Commit()
}
