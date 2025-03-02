CREATE TABLE IF NOT EXISTS stored_files (
    id SERIAL PRIMARY KEY,
    file_name VARCHAR(255) NOT NULL,
    file_size BIGINT NOT NULL,
    mime_type VARCHAR(255) NOT NULL,
    folder VARCHAR(255),
    saved_as VARCHAR(255) NOT NULL,
    original_extension VARCHAR(255) NOT NULL,
    original_file_name VARCHAR(255) NOT NULL,
    uploaded_at TIMESTAMP DEFAULT NOW(),
    project_id INT NOT NULL REFERENCES projects(id),
    icon VARCHAR(255)
)