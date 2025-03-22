package store

import (
	"context"
	"database/sql"
	"log"
)

type Project struct {
	ID            int64  `json:"id"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	CreatedAt     string `json:"created_at"`
	CreatedById   int64  `json:"created_by_id"`
	CreatedBy     User   `json:"created_by"`
	ProjectKey    string `json:"project_key"`
	MaxUploadSize int64  `json:"max_upload_size"`
}

type ProjectStore struct {
	db *sql.DB
}

func (s *ProjectStore) Count(ctx context.Context) (int64, error) {
	query := `SELECT COUNT(*) FROM projects`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	var count int64
	err := s.db.QueryRowContext(ctx, query).Scan(&count)

	return count, err
}

func (s *ProjectStore) Create(ctx context.Context, project *Project) error {
	query := `INSERT INTO projects (name,
							description,
							created_at,
							created_by_id,
							project_key, max_upload_size) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, created_at`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	// Handle NULL for created_by_id when it's 0
	var createdById *int64
	if project.CreatedById == 0 {
		createdById = nil // Set to NULL
	} else {
		createdById = &project.CreatedById // Use pointer to the value
	}

	err := s.db.QueryRowContext(ctx,
		query,
		project.Name,
		project.Description,
		project.CreatedAt,
		createdById,
		project.ProjectKey,
		project.MaxUploadSize,
	).Scan(&project.ID, &project.CreatedAt)

	return err
}

func (s *ProjectStore) GetById(ctx context.Context, id int64) (*Project, error) {
	query := `SELECT id, name, description, created_at, COALESCE(created_by_id, 0), max_upload_size, project_key FROM projects WHERE id = $1`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	project := &Project{}
	err := s.db.QueryRowContext(
		ctx,
		query,
		id,
	).Scan(
		&project.ID,
		&project.Name,
		&project.Description,
		&project.CreatedAt,
		&project.CreatedById,
		&project.MaxUploadSize,
		&project.ProjectKey,
	)

	return project, err
}

func (s *ProjectStore) GetByKey(ctx context.Context, key string) (*Project, error) {
	query := `SELECT id, name, description, created_at, project_key, max_upload_size FROM projects WHERE project_key = $1`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	project := &Project{}
	err := s.db.QueryRowContext(
		ctx,
		query,
		key,
	).Scan(
		&project.ID,
		&project.Name,
		&project.Description,
		&project.CreatedAt,
		&project.ProjectKey,
		&project.MaxUploadSize)

	if err != nil {
		log.Printf("Error getting project by key: %v", err)
	}

	return project, err
}

func (s *ProjectStore) GetAll(ctx context.Context, limit int64, offset int64) ([]*Project, error) {
	query := `SELECT id, name, description, created_at, COALESCE(created_by_id, 0), max_upload_size, project_key FROM projects ORDER BY created_at DESC LIMIT $1 OFFSET $2`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	rows, err := s.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	projects := make([]*Project, 0)
	for rows.Next() {
		project := &Project{}
		err := rows.Scan(
			&project.ID,
			&project.Name,
			&project.Description,
			&project.CreatedAt,
			&project.CreatedById,
			&project.MaxUploadSize,
			&project.ProjectKey,
		)
		if err != nil {
			return nil, err
		}
		projects = append(projects, project)
	}

	return projects, nil
}

func (s *ProjectStore) Update(ctx context.Context, project *Project) error {
	query := `UPDATE projects SET name = $1, description = $2, max_upload_size = $3, project_key = $4 WHERE id = $5`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	_, err := s.db.ExecContext(ctx, query, project.Name, project.Description, project.MaxUploadSize, project.ProjectKey, project.ID)
	return err
}

func (s *ProjectStore) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM projects WHERE id = $1`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	_, err := s.db.ExecContext(ctx, query, id)
	return err
}
