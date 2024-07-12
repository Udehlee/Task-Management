package store

import (
	"database/sql"

	"github.com/Udehlee/Task-Management/pkg/models"
)

type Store interface {
	SaveUser(user models.User) error
	UserByEmail(email string) (models.User, error)

	GetAllUser() ([]models.User, error)
	GetUserById(id int) (models.User, error)

	InsertTask(task models.Task) error
	UpdateTask(task models.Task) error
}

type PgConn struct {
	Conn *sql.DB
}

func NewMysql(db *sql.DB) PgConn {
	return PgConn{
		Conn: db,
	}
}
