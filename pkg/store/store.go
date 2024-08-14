package store

import (
	"database/sql"
)


type PgConn struct {
	Conn *sql.DB
}

func NewMysql(db *sql.DB) PgConn {
	return PgConn{
		Conn: db,
	}
}
