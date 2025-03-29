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

type Counter interface {
	Count(ctx context.Context) (int64, error)
}

type Storage struct {
	Roles interface {
		Counter
		GetAll(ctx context.Context, limit int64, offset int64) ([]*Role, error)
		GetByName(ctx context.Context, name string) (*Role, error)
		Create(ctx context.Context, tx *sql.Tx, role *Role) error
	}

	Users interface {
		Counter
		Create(ctx context.Context, tx *sql.Tx, user *User) error
		GetByEmail(ctx context.Context, email string) (*User, error)
		GetById(ctx context.Context, id int64) (*User, error)
		GetAll(ctx context.Context, limit int64, offset int64) ([]*User, error)
	}

	Projects interface {
		Counter
		Create(ctx context.Context, project *Project) error
		GetById(ctx context.Context, id int64) (*Project, error)
		GetByKey(ctx context.Context, key string) (*Project, error)
		GetAll(ctx context.Context, limit int64, offset int64) ([]*Project, error)
		Update(ctx context.Context, project *Project) error
		Delete(ctx context.Context, id int64) error
	}

	StoredFiles interface {
		Create(ctx context.Context, storedFile *StoredFile) error
		GetById(ctx context.Context, id uuid.UUID) (*StoredFile, error)
		GetAllByProjectId(ctx context.Context, projectId int64, limit int64, offset int64) ([]*StoredFile, error)
		CountProjectFiles(ctx context.Context, projectId int64) (int64, error)
		GetByIdAndProjectKey(ctx context.Context, id uuid.UUID, projectKey string) (*StoredFile, error)
		GetAllByProjectKey(ctx context.Context, projectKey string, limit int64, offset int64) ([]*StoredFile, error)
	}

	FileTypes interface {
		Counter
		Create(ctx context.Context, fileType *FileType) error
		GetById(ctx context.Context, id int64) (*FileType, error)
		GetByMimeType(ctx context.Context, mimeType string) (*FileType, error)
		GetAll(ctx context.Context, limit int64, offset int64) ([]*FileType, error)
	}

	ProjectAllowedFileTypes interface {
		Create(ctx context.Context, projectAllowedFileType *ProjectAllowedFileType) error
		GetByProjectId(ctx context.Context, projectId int64) ([]*ProjectAllowedFileType, error)
		FileTypeIsAllowed(ctx context.Context, projectId int64, mimetype string) (bool, error)
	}

	UserAssignedProjects interface {
		Create(ctx context.Context, tx *sql.Tx, userAssignedProject *UserAssignedProject) error
		CreateWithoutTx(ctx context.Context, userAssignedProject *UserAssignedProject) error
		GetByProjectId(ctx context.Context, projectId int64, limit int64, offset int64) ([]*UserAssignedProject, error)
		GetByUserId(ctx context.Context, userId int64) ([]*UserAssignedProject, error)
		CountByUserId(ctx context.Context, userId int64) (int64, error)
		CountUsersByProjectId(ctx context.Context, projectId int64) (int64, error)
		ProjectIsAssignedToUser(ctx context.Context, projectId int64, userId int64) (bool, error)
	}
}

func InitialiseStorage(db *sql.DB) *Storage {
	return &Storage{
		Roles:                   &RoleStore{db},
		Users:                   &UserStore{db},
		Projects:                &ProjectStore{db},
		StoredFiles:             &StoredFileStore{db},
		FileTypes:               &FileTypeStore{db},
		ProjectAllowedFileTypes: &ProjectAllowedFileTypeStore{db},
		UserAssignedProjects:    &UserProjectStore{db},
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
