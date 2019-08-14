package interfaces

import "database/sql"

// ISQL - Interface SQL Connection
type ISQL interface {
	Connect() (*sql.DB, error)
}
