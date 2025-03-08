package store

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
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

	Projects interface {
		Create(ctx context.Context, project *Project) error
		GetById(ctx context.Context, id int64) (*Project, error)
		GetByKey(ctx context.Context, key string) (*Project, error)
		GetAll(ctx context.Context, limit int64, offset int64) ([]*Project, error)
	}

	StoredFiles interface {
		Create(ctx context.Context, storedFile *StoredFile) error
		GetById(ctx context.Context, id uuid.UUID) (*StoredFile, error)
		GetAllByProjectId(ctx context.Context, projectId int64, limit int64, offset int64) ([]*StoredFile, error)
		GetByIdAndProjectKey(ctx context.Context, id uuid.UUID, projectKey string) (*StoredFile, error)
		GetAllByProjectKey(ctx context.Context, projectKey string, limit int64, offset int64) ([]*StoredFile, error)
	}
}

func InitialiseStorage(db *sql.DB) *Storage {
	return &Storage{
		Roles: &RoleStore{db},
		Users: &UserStore{db},
		Projects: &ProjectStore{db},
		StoredFiles: &StoredFileStore{db},
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