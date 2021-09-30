//go:build !go1.15
// +build !go1.15

package dbx

import (
	"database/sql"
	"time"
)

func dbSetConnMaxIdleTime(db *sql.DB, duration time.Duration) {
	// Not supported yet
}
