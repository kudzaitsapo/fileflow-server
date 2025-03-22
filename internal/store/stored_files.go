package store

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type StoredFile struct {
	ID                uuid.UUID `json:"id"`
	FileName          string    `json:"name"`
	FileSize          int64     `json:"size"`
	MimeType          string    `json:"mime_type"`
	Folder            string    `json:"folder"`
	SavedAs           string    `json:"saved_as"`
	OriginalExtension string    `json:"original_extension"`
	UploadedAt        string    `json:"uploaded_at"`
	Project           Project   `json:"project"`
	ProjectID         int64     `json:"project_id"`
	Icon              string    `json:"icon"`
	FileType          FileType  `json:"file_type"`
}

type StoredFileStore struct {
	db *sql.DB
}

func (s *StoredFileStore) Create(ctx context.Context, storedFile *StoredFile) error {

	query := `INSERT INTO stored_files (file_name,
							file_size,
							mime_type,
							folder,
							saved_as,
							original_extension,
							uploaded_at,
							project_id,
							icon) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id, file_name`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	err := s.db.QueryRowContext(ctx,
		query,
		storedFile.FileName,
		storedFile.FileSize,
		storedFile.MimeType,
		storedFile.Folder,
		storedFile.SavedAs,
		storedFile.OriginalExtension,
		time.Now(),
		storedFile.ProjectID,
		storedFile.Icon,
	).Scan(
		&storedFile.ID,
		&storedFile.FileName,
	)

	return err
}

func (s *StoredFileStore) GetById(ctx context.Context, id uuid.UUID) (*StoredFile, error) {
	query := `SELECT sf.id, sf.file_name, sf.file_size, sf.mime_type, sf.folder, sf.saved_as, sf.original_extension, sf.uploaded_at, sf.project_id, sf.icon, p.name, p.description, p.created_at, COALESCE(p.created_by_id, 0)
	FROM stored_files sf
	INNER JOIN projects p ON sf.project_id = p.id
	WHERE sf.id = $1`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	storedFile := &StoredFile{}
	err := s.db.QueryRowContext(
		ctx,
		query,
		id,
	).Scan(
		&storedFile.ID,
		&storedFile.FileName,
		&storedFile.FileSize,
		&storedFile.MimeType,
		&storedFile.Folder,
		&storedFile.SavedAs,
		&storedFile.OriginalExtension,
		&storedFile.UploadedAt,
		&storedFile.ProjectID,
		&storedFile.Icon,
		&storedFile.Project.Name,
		&storedFile.Project.Description,
		&storedFile.Project.CreatedAt,
		&storedFile.Project.CreatedById,
	)

	return storedFile, err
}

func (s *StoredFileStore) GetByIdAndProjectKey(ctx context.Context, id uuid.UUID, projectKey string) (*StoredFile, error) {
	query := `SELECT sf.id, sf.file_name, sf.file_size, sf.mime_type, sf.folder, sf.saved_as, sf.original_extension, sf.uploaded_at, sf.project_id, sf.icon
	FROM stored_files sf
	JOIN projects p ON sf.project_id = p.id
	WHERE sf.id = $1 AND p.project_key = $2`
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	storedFile := &StoredFile{}
	err := s.db.QueryRowContext(
		ctx,
		query,
		id,
		projectKey,
	).Scan(
		&storedFile.ID,
		&storedFile.FileName,
		&storedFile.FileSize,
		&storedFile.MimeType,
		&storedFile.Folder,
		&storedFile.SavedAs,
		&storedFile.OriginalExtension,
		&storedFile.UploadedAt,
		&storedFile.ProjectID,
		&storedFile.Icon,
	)

	return storedFile, err
}

func (s *StoredFileStore) GetAllByProjectKey(ctx context.Context,
	projectKey string, limit int64, offset int64) ([]*StoredFile, error) {
	query := `SELECT sf.id, sf.file_name, sf.file_size, sf.mime_type, sf.folder, sf.saved_as, sf.original_extension, sf.uploaded_at, sf.project_id, sf.icon
	FROM stored_files sf
	JOIN projects p ON sf.project_id = p.id
	WHERE p.project_key = $1
	ORDER BY sf.uploaded_at DESC
	LIMIT $2 OFFSET $3`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	rows, err := s.db.QueryContext(ctx, query, projectKey, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	storedFiles := make([]*StoredFile, 0)
	for rows.Next() {
		storedFile := &StoredFile{}
		err := rows.Scan(
			&storedFile.ID,
			&storedFile.FileName,
			&storedFile.FileSize,
			&storedFile.MimeType,
			&storedFile.Folder,
			&storedFile.SavedAs,
			&storedFile.OriginalExtension,
			&storedFile.UploadedAt,
			&storedFile.ProjectID,
			&storedFile.Icon,
		)
		if err != nil {
			return nil, err
		}
		storedFiles = append(storedFiles, storedFile)
	}

	return storedFiles, nil
}

func (s *StoredFileStore) GetAllByProjectId(ctx context.Context, projectId int64, limit int64, offset int64) ([]*StoredFile, error) {
	query := `SELECT sf.id, sf.file_name, sf.file_size, sf.mime_type, sf.folder, sf.saved_as, sf.original_extension, sf.uploaded_at, sf.project_id, 
	sf.icon, ft.name, ft.id, ft.mimetype FROM stored_files sf INNER JOIN file_types ft on sf.mime_type = ft.mimetype WHERE sf.project_id = $1 ORDER BY uploaded_at DESC LIMIT $2 OFFSET $3`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	rows, err := s.db.QueryContext(ctx, query, projectId, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	storedFiles := make([]*StoredFile, 0)
	for rows.Next() {
		storedFile := &StoredFile{}
		err := rows.Scan(
			&storedFile.ID,
			&storedFile.FileName,
			&storedFile.FileSize,
			&storedFile.MimeType,
			&storedFile.Folder,
			&storedFile.SavedAs,
			&storedFile.OriginalExtension,
			&storedFile.UploadedAt,
			&storedFile.ProjectID,
			&storedFile.Icon,
			&storedFile.FileType.Name,
			&storedFile.FileType.ID,
			&storedFile.FileType.MimeType,
		)
		if err != nil {
			return nil, err
		}
		storedFiles = append(storedFiles, storedFile)
	}

	return storedFiles, nil
}

func (s *StoredFileStore) CountProjectFiles(ctx context.Context, projectId int64) (int64, error) {
	query := `SELECT COUNT(*) FROM stored_files WHERE project_id = $1`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	var count int64
	err := s.db.QueryRowContext(ctx, query, projectId).Scan(&count)
	return count, err
}
