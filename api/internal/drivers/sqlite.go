package drivers

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3" // SQLite driver
)

// SQLiteOption struct
type SQLiteOption struct {
	URL     string
	MaxIdle int
	MaxOpen int
}

// NewSQLite return a client connection handle to a SQLite server.
func NewSQLite(option SQLiteOption) (*sql.DB, error) {
	// pwd, err := os.Getwd()
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
	// dbPath := pwd + option.URL

	db, err := sql.Open("sqlite3", option.URL)
	if err != nil {
		return nil, err
	}
	db.SetMaxIdleConns(option.MaxIdle)
	db.SetMaxOpenConns(option.MaxOpen)
	db.SetConnMaxLifetime(time.Hour)

	// Check if the database is reachable
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to connect to SQLite database: %w", err)
	}

	return db, nil
}
