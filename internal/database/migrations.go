package database

import (
	"database/sql"
	"embed"
	"fmt"
	"log"
	"sort"
	"strings"
)

type Migration struct {
	Name    string
	Content string
	Version string
}

// Migration handling code
const createMigrationsTableSQL = `
CREATE TABLE IF NOT EXISTS schema_migrations (
    version TEXT PRIMARY KEY,
    applied_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
)`

// Get already applied migrations
func getAppliedMigrations(db *sql.DB) (map[string]bool, error) {
	rows, err := db.Query("SELECT version FROM schema_migrations")
	if err != nil {
		return nil, fmt.Errorf("failed to query migrations: %w", err)
	}
	defer rows.Close()

	applied := make(map[string]bool)
	for rows.Next() {
		var version string
		if err := rows.Scan(&version); err != nil {
			return nil, fmt.Errorf("failed to scan version: %w", err)
		}
		applied[version] = true
	}
	return applied, nil
}

// Read migration files provided
func readMigrationFiles(fs embed.FS, applied map[string]bool) ([]Migration, error) {
	var migrations []Migration

	files, err := fs.ReadDir("migrations")
	if err != nil {
		return nil, fmt.Errorf("failed to read migrations dir: %w", err)
	}

	for _, f := range files {
		if f.IsDir() || !strings.HasSuffix(f.Name(), ".sql") {
			continue
		}

		version, _, found := strings.Cut(f.Name(), "_")
		if !found || version == "" {
			return nil, fmt.Errorf("invalid migration name: %s", f.Name())
		}

		if applied[version] {
			continue
		}

		content, err := fs.ReadFile("migrations/" + f.Name())
		if err != nil {
			return nil, fmt.Errorf("failed to read migration file: %w", err)
		}

		migrations = append(migrations, Migration{
			Name:    f.Name(),
			Content: string(content),
			Version: version,
		})
	}

	// Sort migrations by version
	sort.Slice(migrations, func(i, j int) bool {
		return migrations[i].Version < migrations[j].Version
	})

	return migrations, nil
}

func applyMigration(db *sql.DB, m Migration) error  {
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	// Execute migration SQL
	if _, err := tx.Exec(m.Content); err != nil {
		return fmt.Errorf("failed to execute migration: %w", err)
	}

	// Record migration
	if _, err := tx.Exec(
		"INSERT INTO schema_migrations (version) VALUES ($1)",
		m.Version,
	); err != nil {
		return fmt.Errorf("failed to record migration: %w", err)
	}

	return tx.Commit()
}

func RunMigrations(db *sql.DB, fs embed.FS) error {
	// Create migrations table if not exists
	if _, err := db.Exec(createMigrationsTableSQL); err != nil {
		return fmt.Errorf("failed to create migrations table: %w", err)
	}

	// Get already applied migrations
	applied, err := getAppliedMigrations(db)
	if err != nil {
		return err
	}

	// Read migration files from embedded FS
	migrations, err := readMigrationFiles(fs, applied)
	if err != nil {
		return err
	}

	// Apply migrations in order
	for _, m := range migrations {
		if err := applyMigration(db, m); err != nil {
			return fmt.Errorf("migration failed: %s - %w", m.Name, err)
		}
		log.Printf("Applied migration: %s", m.Name)
	}

	return nil
}
