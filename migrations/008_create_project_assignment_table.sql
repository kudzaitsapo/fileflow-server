CREATE TABLE IF NOT EXISTS user_assigned_projects (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id) NOT NULL,
    project_id INT REFERENCES projects(id) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);