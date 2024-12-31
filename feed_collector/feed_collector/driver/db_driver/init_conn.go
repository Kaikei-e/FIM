package db_driver

import (
	"context"
	"database/sql"
	"feed_collector/slogger"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func InitConn(ctx context.Context) (*sql.DB, error) {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	postgresHost := os.Getenv("POSTGRES_HOST")
	if postgresHost == "" {
		panic("POSTGRES_HOST is not set in .env file")
	}

	postgresPort := os.Getenv("POSTGRES_PORT")
	if postgresPort == "" {
		panic("POSTGRES_PORT is not set in .env file")
	}

	postgresUser := os.Getenv("POSTGRES_USER")
	if postgresUser == "" {
		panic("POSTGRES_USER is not set in .env file")
	}

	postgresPassword := os.Getenv("POSTGRES_PASSWORD")
	if postgresPassword == "" {
		panic("POSTGRES_PASSWORD is not set in .env file")
	}

	postgresDB := os.Getenv("POSTGRES_DB")
	if postgresDB == "" {
		panic("POSTGRES_DB is not set in .env file")
	}

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", postgresUser, postgresPassword, postgresHost, postgresPort, postgresDB)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		slogger.Logger.Error("Failed to open a connection to the database.", "Caused by", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		slogger.Logger.Error("Failed to ping the database.", "Caused by", err)
		return nil, err
	}

	return nil, nil
}
