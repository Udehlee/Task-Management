package db

import (
	"log"

	"github.com/Udehlee/Task-Management/pkg/store"
)

func InitDB(config Config) (store.PgConn, error) {

	var pg store.PgConn

	db, err := ConnectDB(config)
	if err != nil {
		log.Fatal(err)
	}

	pg.Conn = db

	return pg, nil

}
