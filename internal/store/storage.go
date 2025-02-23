package store

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

var (
	ErrNotFound          = errors.New("resource not found")
	ErrConflict          = errors.New("resource already exists")
	QueryTimeoutDuration = time.Second * 5
)


type Storage struct {
	Roles interface {
		GetByName(ctx context.Context, name string) (*Role, error)
		Create(ctx context.Context, tx *sql.Tx, role *Role) error
	}

	Users interface {
		Create(ctx context.Context, tx *sql.Tx, user *User) error
		GetByEmail(ctx context.Context, email string) (*User, error)
		GetById(ctx context.Context, id int64) (*User, error)
		GetAll(ctx context.Context, limit int64, offset int64) ([]*User, error)
	}
}

func InitialiseStorage(db *sql.DB) *Storage {
	return &Storage{
		Roles: &RoleStore{db},
		Users: &UserStore{db},
	}
}

func withTx(db *sql.DB, ctx context.Context, fn func(*sql.Tx) error) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	if err := fn(tx); err != nil {
		_ = tx.Rollback()
		return err
	}

	return tx.Commit()
}