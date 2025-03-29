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

type UserAssignedProject struct {
	ID        int64   `json:"id"`
	ProjectID int64   `json:"project_id"`
	UserID    int64   `json:"user_id"`
	Project   Project `json:"project"`
	User      User    `json:"user"`
}

type ProjectStore struct {
	db *sql.DB
}

type UserProjectStore struct {
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

func (s *UserProjectStore) Create(ctx context.Context, tx *sql.Tx, userAssignedProject *UserAssignedProject) error {
	query := `INSERT INTO user_assigned_projects (project_id, user_id) VALUES ($1, $2) RETURNING id`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	err := tx.QueryRowContext(ctx, query, userAssignedProject.ProjectID, userAssignedProject.UserID).Scan(&userAssignedProject.ID)
	return err
}

func (s *UserProjectStore) CreateWithoutTx(ctx context.Context, userAssignedProject *UserAssignedProject) error {
	return withTx(s.db, ctx, func(tx *sql.Tx) error {
		query := `INSERT INTO user_assigned_projects (project_id, user_id) VALUES ($1, $2) RETURNING id`

		ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
		defer cancel()

		err := tx.QueryRowContext(ctx, query, userAssignedProject.ProjectID, userAssignedProject.UserID).Scan(&userAssignedProject.ID)
		if err != nil {
			return err
		}

		return nil
	},
	)
}

func (s *UserProjectStore) GetByProjectId(ctx context.Context, projectId int64, limit int64, offset int64) ([]*UserAssignedProject, error) {
	query := `SELECT uap.id, uap.project_id, uap.user_id, u.id, u.email, u.first_name, u.last_name, u.created_at, u.is_active, u.role_id
				FROM user_assigned_projects uap
				INNER JOIN projects p ON uap.project_id = p.id
				INNER JOIN users u ON uap.user_id = u.id
				WHERE uap.project_id = $1
				ORDER BY u.created_at DESC
				LIMIT $2 OFFSET $3`
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	rows, err := s.db.QueryContext(ctx, query, projectId, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	assignments := make([]*UserAssignedProject, 0)
	for rows.Next() {
		assignment := &UserAssignedProject{}
		err := rows.Scan(&assignment.ID,
			&assignment.ProjectID,
			&assignment.UserID,
			&assignment.User.ID,
			&assignment.User.Email,
			&assignment.User.FirstName,
			&assignment.User.LastName,
			&assignment.User.CreatedAt,
			&assignment.User.IsActive,
			&assignment.User.RoleID,
		)
		if err != nil {
			return nil, err
		}
		assignments = append(assignments, assignment)
	}
	return assignments, nil
}

// TODO: Return actual users and projects instead of ids
func (s *UserProjectStore) GetByUserId(ctx context.Context, userId int64) ([]*UserAssignedProject, error) {
	query := `SELECT 
					uap.id,
					uap.project_id, 
					uap.user_id,
					p.id,
					p.name,
					p.description,
					p.created_at,
					p.max_upload_size
			   FROM user_assigned_projects uap
			   INNER JOIN projects p ON uap.project_id = p.id
			   WHERE uap.user_id = $1`
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	rows, err := s.db.QueryContext(ctx, query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	assignments := make([]*UserAssignedProject, 0)
	for rows.Next() {
		assignment := &UserAssignedProject{}
		err := rows.Scan(&assignment.ID,
			&assignment.ProjectID,
			&assignment.UserID,
			&assignment.Project.ID,
			&assignment.Project.Name,
			&assignment.Project.Description,
			&assignment.Project.CreatedAt,
			&assignment.Project.MaxUploadSize,
		)
		if err != nil {
			return nil, err
		}
		assignments = append(assignments, assignment)
	}
	return assignments, nil
}

func (s *UserProjectStore) CountByUserId(ctx context.Context, userId int64) (int64, error) {
	query := `SELECT COUNT(*) FROM user_assigned_projects WHERE user_id = $1`
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	var count int64
	err := s.db.QueryRowContext(ctx, query, userId).Scan(&count)
	return count, err
}

func (s *UserProjectStore) CountUsersByProjectId(ctx context.Context, projectId int64) (int64, error) {
	query := `SELECT COUNT(u.*) 
				FROM user_assigned_projects uap
				INNER JOIN users u ON uap.user_id = u.id
				WHERE uap.project_id = $1`
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	var count int64
	err := s.db.QueryRowContext(ctx, query, projectId).Scan(&count)
	return count, err
}

func (s *UserProjectStore) ProjectIsAssignedToUser(ctx context.Context, projectId int64, userId int64) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM user_assigned_projects WHERE project_id = $1 AND user_id = $2)`
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	var exists bool
	err := s.db.QueryRowContext(ctx, query, projectId, userId).Scan(&exists)
	return exists, err
}
