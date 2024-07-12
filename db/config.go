package db

import (
	"os"
	"strconv"
)

type Config struct {
	Host     string
	Port     int
	Username string
	Password string
	DbName   string
}

func LoadConfig() (Config, error) {

	var cfg Config

	port, err := strconv.Atoi(os.Getenv("POSTGRES_PORT"))
	if err != nil {
		return cfg, err
	}

	cfg = Config{

		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     port,
		Username: os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		DbName:   os.Getenv("POSTGRES_DB"),
	}
	return cfg, nil

}
