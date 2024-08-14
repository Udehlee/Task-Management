package store

import (
	"database/sql"

	"github.com/Udehlee/Task-Management/pkg/models"
)


type PgConn struct {
	Conn *sql.DB
}

func NewMysql(db *sql.DB) PgConn {
	return PgConn{
		Conn: db,
	}
}
