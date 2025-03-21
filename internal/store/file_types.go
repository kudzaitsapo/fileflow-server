package store

import (
	"context"
	"database/sql"
	"time"
)

type FileType struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	MimeType    string `json:"mime_type"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
	CreatedAt   string `json:"created_at"`
}

type ProjectAllowedFileType struct {
	ID         int64    `json:"id"`
	ProjectID  int64    `json:"project_id"`
	FileTypeID int64    `json:"file_type_id"`
	FileType   FileType `json:"file_type"`
	CreatedAt  string   `json:"created_at"`
}

type FileTypeStore struct {
	db *sql.DB
}

type ProjectAllowedFileTypeStore struct {
	db *sql.DB
}

func (s *FileTypeStore) Create(ctx context.Context, fileType *FileType) error {
	query := `INSERT INTO file_types (name, mimetype, description, icon, created_at) VALUES ($1, $2, $3, $4, $5) RETURNING id`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	err := s.db.QueryRowContext(ctx, query, fileType.Name, fileType.MimeType, fileType.Description, fileType.Icon, time.Now()).Scan(&fileType.ID)
	if err != nil {
		return err
	}

	return nil
}

func (s *FileTypeStore) Count(ctx context.Context) (int64, error) {
	query := `SELECT COUNT(id) FROM file_types`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	var count int64
	err := s.db.QueryRowContext(ctx, query).Scan(&count)
	return count, err
}

func (s *FileTypeStore) GetById(ctx context.Context, id int64) (*FileType, error) {
	query := `SELECT id, name, mime_type, description, icon, created_at FROM file_types WHERE id = $1`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	fileType := &FileType{}
	err := s.db.QueryRowContext(ctx, query, id).Scan(&fileType.ID, &fileType.Name, &fileType.MimeType, &fileType.Description, &fileType.Icon, &fileType.CreatedAt)
	return fileType, err
}

func (s *FileTypeStore) GetByMimeType(ctx context.Context, mimeType string) (*FileType, error) {
	query := `SELECT id, name, mime_type, description, icon, created_at FROM file_types WHERE mime_type = $1`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	fileType := &FileType{}
	err := s.db.QueryRowContext(ctx, query, mimeType).Scan(&fileType.ID, &fileType.Name, &fileType.MimeType, &fileType.Description, &fileType.Icon, &fileType.CreatedAt)
	return fileType, err
}

func (s *FileTypeStore) GetAll(ctx context.Context, limit int64, offset int64) ([]*FileType, error) {
	query := `SELECT id, name, mime_type, description, icon, created_at FROM file_types LIMIT $1 OFFSET $2`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	rows, err := s.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	fileTypes := make([]*FileType, 0)
	for rows.Next() {
		fileType := &FileType{}
		err := rows.Scan(&fileType.ID, &fileType.Name, &fileType.MimeType, &fileType.Description, &fileType.Icon, &fileType.CreatedAt)
		if err != nil {
			return nil, err
		}
		fileTypes = append(fileTypes, fileType)
	}

	return fileTypes, nil
}

func (s *ProjectAllowedFileTypeStore) Create(ctx context.Context, projectAllowedFileType *ProjectAllowedFileType) error {
	query := `INSERT INTO project_allowed_file_types (project_id, file_type_id, created_at) VALUES ($1, $2, $3) RETURNING id`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	err := s.db.QueryRowContext(ctx, query, projectAllowedFileType.ProjectID, projectAllowedFileType.FileTypeID, time.Now()).Scan(&projectAllowedFileType.ID)
	if err != nil {
		return err
	}

	return nil
}

func (s *ProjectAllowedFileTypeStore) GetByProjectId(ctx context.Context, projectId int64) ([]*ProjectAllowedFileType, error) {
	query := `SELECT paft.id, paft.project_id, paft.file_type_id, paft.created_at, ft.name, ft.mimetype, ft.description, ft.icon FROM project_allowed_file_types paft INNER JOIN file_types ft ON paft.file_type_id = ft.id WHERE paft.project_id = $1`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	rows, err := s.db.QueryContext(ctx, query, projectId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	projectAllowedFileTypes := make([]*ProjectAllowedFileType, 0)
	for rows.Next() {
		projectAllowedFileType := &ProjectAllowedFileType{}
		err := rows.Scan(
			&projectAllowedFileType.ID,
			&projectAllowedFileType.ProjectID,
			&projectAllowedFileType.FileTypeID,
			&projectAllowedFileType.CreatedAt,
			&projectAllowedFileType.FileType.Name,
			&projectAllowedFileType.FileType.MimeType,
			&projectAllowedFileType.FileType.Description,
			&projectAllowedFileType.FileType.Icon,
		)
		if err != nil {
			return nil, err
		}
		projectAllowedFileTypes = append(projectAllowedFileTypes, projectAllowedFileType)
	}

	return projectAllowedFileTypes, nil
}

func (s *ProjectAllowedFileTypeStore) FileTypeIsAllowed(ctx context.Context, projectId int64, mimetype string) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM project_allowed_file_types paft INNER JOIN file_types ft ON paft.file_type_id = ft.id WHERE project_id = $1 AND ft.mimetype = $2)`
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	var exists bool
	err := s.db.QueryRowContext(ctx, query, projectId, mimetype).Scan(&exists)
	return exists, err
}
