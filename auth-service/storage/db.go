package storage

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"
)

func NewStorage(dsn string, options *ConnOptions) (Storage, error) {
	db, err := connectDB(dsn, options)
	if err != nil {
		return Storage{}, err
	}

	return Storage{db}, nil
}

func connectDB(dsn string, options *ConnOptions) (*sql.DB, error) {
	attempts := 1
	delay := 2

	if options != nil {
		attempts = options.Attempts
		delay = options.DelayInSecond
	}

	for {
		attempts--

		db, err := connectPostgres(dsn)
		if err == nil {
			return db, err
		}

		log.Printf("Cannot connect to postgres.\n")

		if attempts == 0 {
			return nil, err
		}

		log.Printf("Retry in %d sec...", delay)
		time.Sleep(time.Duration(delay) * time.Second)
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
