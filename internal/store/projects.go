package store

import (
	"context"
	"database/sql"
	"log"
)

type Project struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	CreatedById  int64 `json:"created_by_id"`
	CreatedBy   User `json:"created_by"`
	ProjectKey string `json:"project_key"`
}


type ProjectStore struct {
	db *sql.DB
}

func (s *ProjectStore) Create(ctx context.Context, project *Project) error {
	query := `INSERT INTO projects (name,
							description,
							created_at,
							created_by_id,
							project_key) VALUES ($1, $2, $3, $4, $5) RETURNING id, created_at`

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
			).Scan(&project.ID, &project.CreatedAt)

	return err
}

func (s *ProjectStore) GetById(ctx context.Context, id int64) (*Project, error) {
	query := `SELECT id, name, description, created_at, COALESCE(created_by_id, 0) FROM projects WHERE id = $1`

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
			)

	return project, err
}

func (s *ProjectStore) GetByKey(ctx context.Context, key string) (*Project, error) {
	query := `SELECT id, name, description, created_at, project_key FROM projects WHERE project_key = $1`

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
				&project.ProjectKey)

	if err != nil {
		log.Printf("Error getting project by key: %v", err)
	}

	return project, err
}

func (s *ProjectStore) GetAll(ctx context.Context, limit int64, offset int64) ([]*Project, error) {
	query := `SELECT id, name, description, created_at, COALESCE(created_by_id, 0) FROM projects LIMIT $1 OFFSET $2`

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
		)
		if err != nil {
			return nil, err
		}
		projects = append(projects, project)
	}

	return projects, nil
}