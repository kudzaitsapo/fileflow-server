package main

import (
	"log"

	server "github.com/kudzaitsapo/fileflow-server"
	"github.com/kudzaitsapo/fileflow-server/internal/cache"
	"github.com/kudzaitsapo/fileflow-server/internal/config"
	"github.com/kudzaitsapo/fileflow-server/internal/database"
	"github.com/kudzaitsapo/fileflow-server/internal/middleware"
	"github.com/kudzaitsapo/fileflow-server/internal/routes"
	"github.com/kudzaitsapo/fileflow-server/internal/seeds"
	"github.com/kudzaitsapo/fileflow-server/internal/store"
)

func main() {
	cfg, err := config.LoadConfig()

	if err != nil {
		log.Fatalf("error loading config: %v", err)
	}

	app := CreateApplication(*cfg)

	// Initialise the database
	db, err := database.Initialise(&cfg.DbConfig)

	if err != nil {
		log.Fatalf("error initialising database: %v", err)
	}
	defer db.Close()
	log.Printf("Database connection established")

	// Initialise redis cache
	redis := cache.Initialise(cfg.RedisConfig)
	defer redis.Close()
	log.Printf("Redis connection established")

	// Register middleware
	middlewares := middleware.GetMiddlewares()
	for _, middleware := range middlewares {
		app.Use(middleware)
	}

	// Handle route registration
	routes := routes.CreateRoutes()
	for _, route := range routes {
		app.Handle(route.Pattern, route.Handler)
	}

	// Run database migrations
	if !cfg.DbConfig.SkipMigrations {
		if err := database.RunMigrations(db, server.MigrationsDir); err != nil {
			log.Fatalf("error running migrations: %v", err)
		}
	}

	// Set the store and cache
	store := store.InitialiseStorage(db)
	app.SetStore(store)

	cache := cache.InitialiseStorage(db)
	app.SetCache(cache)


	// Seed the database
	if !cfg.DbConfig.SkipSeeding {
		if err := seeds.Seed(store, db); err != nil {
			log.Fatalf("error seeding database: %v", err)
		}
	}

	log.Printf("Server started on port %d", cfg.Config.Port)

	if err := app.ListenAndServe(); err != nil {
		log.Fatalf("error starting server: %v", err)
	}

}