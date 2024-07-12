package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// ConnectDB connects to database
func ConnectDB(config Config) (*sql.DB, error) {

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.Username,
		config.Password, config.DbName)

	dbConn, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("error creating database connection")
	}

	dbConn.SetMaxOpenConns(10)
	dbConn.SetMaxIdleConns(10)

	err = dbConn.Ping()
	if err != nil {
		return nil, fmt.Errorf("connection not alive")

	}

	fmt.Println("connected successfully")
	return dbConn, nil

}
