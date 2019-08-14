package drivers

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// SQL struct
type SQL struct {
	Host          string
	Port          int
	Username      string
	Password      string
	DBName        string
	Charset       string
	MaxConnection int
	MaxLifeTime   int
}

// Connect - Create connection to SQL
func (s *SQL) Connect() (*sql.DB, error) {
	var dataSource = fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		s.Username,
		s.Password,
		s.Host,
		s.Port,
		s.DBName,
		s.Charset,
	)
	db, err := sql.Open("mysql", dataSource)
	if err != nil {
		db.Close()
		return nil, err
	}
	db.SetMaxOpenConns(s.MaxConnection)
	db.SetConnMaxLifetime(time.Duration(time.Duration(s.MaxLifeTime) * time.Millisecond))
	return db, nil
}
