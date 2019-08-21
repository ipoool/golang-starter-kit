package interfaces

import "database/sql"

// ISQL - Interface SQL Connection
type ISQL interface {
	GetConnect() (db *sql.DB, err error)
}
