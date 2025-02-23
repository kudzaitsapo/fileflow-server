package cache

import "database/sql"

type Storage struct {
	db *sql.DB
}

func InitialiseStorage(db *sql.DB) *Storage {
	return &Storage{db}
}