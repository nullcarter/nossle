package main

import (
	"context"
	"database/sql"
	"log"
	"time"

	// SQLite3 driver for database/sql
	_ "github.com/mattn/go-sqlite3"

	// Application-specific packages
	"github.com/nullcarter/nossle/cmd/api"
	"github.com/nullcarter/nossle/internal/store"

	// Migration packages
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	// Application configuration setup
	cfg := api.Config{
		Addr:               ":8080",                // Address where the server listens
		DbPath:             "./sqlite.db",          // Path to the SQLite database file
		MigrationPath:      "file://db/migrations", // Path to migration files
		DatabaseUrl:        "sqlite://./sqlite.db", // SQLite URL for migrations
		DatabaseDriverName: "sqlite3",              // SQLite driver name
	}

	// Open database connection
	db, err := sql.Open(cfg.DatabaseDriverName, cfg.DbPath)
	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}
	defer db.Close()

	// Create a context with timeout for database ping
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// Check that the database is reachable
	if err := db.PingContext(ctx); err != nil {
		log.Fatalf("failed to ping database: %v", err)
	}

	// Create a migration instance with the configured paths
	m, err := migrate.New(
		cfg.MigrationPath,
		cfg.DatabaseUrl,
	)
	if err != nil {
		log.Fatalf("failed to create migrate instance: %v", err)
	}
	defer m.Close()

	// Apply all up migrations
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("migration up failed: %v", err)
	}
	log.Println("Migration complete.")

	// Initialize application store (data layer)
	appStore := store.New(db)

	nossle := &api.Nossle{
		Config: cfg,
		Store:  appStore,
	}

	// Mount routes and middleware
	mux := nossle.Mount()

	// Start the server and block on fatal errors
	log.Fatal(nossle.Run(mux))
}
