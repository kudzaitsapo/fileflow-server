CREATE TABLE IF NOT EXISTS file_types (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    mimetype VARCHAR(255) NOT NULL,
    description TEXT,
    icon TEXT,
    created_at TIMESTAMP DEFAULT NOW()
);


CREATE TABLE IF NOT EXISTS project_allowed_file_types (
    id SERIAL PRIMARY KEY,
    project_id INT NOT NULL REFERENCES projects(id),
    file_type_id INT NOT NULL REFERENCES file_types(id),
    created_at TIMESTAMP DEFAULT NOW()
)