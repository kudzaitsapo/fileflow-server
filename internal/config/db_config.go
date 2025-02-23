package config

type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
	Driver   string
	SkipMigrations bool
	SkipSeeding bool
}