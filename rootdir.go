package server

import "embed"

//go:embed migrations/*.sql
var MigrationsDir embed.FS