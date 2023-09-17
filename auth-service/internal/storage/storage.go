package storage

import (
	"auth-service/internal/config"
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"
)

type Storage struct {
	*sql.DB
}

func NewStorage() *Storage {
	return &Storage{}
}

func (s *Storage) InitDB(config *config.Config) {
	db, err := connectDB(&config.Storage)
	if err != nil {
		log.Fatalf("Failed to connect postgres: %v\n", err)
	}

	s.DB = db
}

func (s *Storage) CloseDB() {
	s.DB.Close()
}

func connectDB(dbconfig *config.StorageConfig) (*sql.DB, error) {
	attempts := dbconfig.Attempts

	for {
		attempts--

		db, err := connectPostgres(dbconfig.DSN)
		if err == nil {
			return db, err
		}

		log.Printf("Cannot connect to postgres.\n")

		if attempts == 0 {
			return nil, err
		}

		log.Printf("Retry in %d sec...", dbconfig.DelayInSecond)
		time.Sleep(time.Duration(dbconfig.DelayInSecond) * time.Second)
	}
}

func connectPostgres(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
