package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/kudzaitsapo/fileflow-server/internal/config"
)

func Initialise(cfg *config.DBConfig) (*sql.DB, error) {
	connString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Database)
	db, err := sql.Open(cfg.Driver, connString)

	if err != nil {
        return nil, err
    }

    err = db.Ping()
    if err != nil {
        return nil, err
    }

	return db, err
}

