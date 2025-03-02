package store

import (
	"context"
	"database/sql"
	"time"
)


type StoredFile struct {
	ID        int64  `json:"id"`
	FileName      string `json:"name"`
	FileSize      int64 `json:"size"`
	MimeType 	string `json:"mime_type"`
	Folder   string `json:"folder"`
	SavedAs  string `json:"saved_as"`
	OriginalExtension string `json:"original_extension"`
	OriginalFileName string `json:"original_file_name"`
	UploadedAt string `json:"uploaded_at"`
	Project  Project `json:"project"`
	ProjectID int64 `json:"project_id"`
	Icon string `json:"icon"`
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
							original_file_name,
							uploaded_at,
							project_id,
							icon) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id, file_name`

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
		storedFile.OriginalFileName,
		time.Now(),
		storedFile.ProjectID,
		storedFile.Icon,
	).Scan(
		&storedFile.ID,
		&storedFile.FileName,
	)

	return err
}

func (s *StoredFileStore) GetById(ctx context.Context, id int64) (*StoredFile, error) {
	query := `SELECT id, file_name, file_size, mime_type, folder, saved_as, original_extension, uploaded_at, project_id,
	icon FROM stored_files WHERE id = $1`

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

