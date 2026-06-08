package db

import (
	"database/sql"
	_ "embed"

	_ "modernc.org/sqlite"
)

//go:embed schema.sql
var schema string

//go:embed seed.sql
var seed string

// Open creates (if needed) and seeds a SQLite database at path.
// Use ":memory:" for an ephemeral, isolated database (e.g. in tests).
func Open(path string) (*sql.DB, error) {
	conn, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, err
	}

	if _, err := conn.Exec(schema); err != nil {
		conn.Close()
		return nil, err
	}

	if _, err := conn.Exec(seed); err != nil {
		conn.Close()
		return nil, err
	}

	return conn, nil
}
