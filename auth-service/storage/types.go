package storage

import "database/sql"

type Storage struct {
	DB *sql.DB
}

type ConnOptions struct {
	Attempts      int
	DelayInSecond int
}
