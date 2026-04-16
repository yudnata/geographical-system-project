package database

import (
	"context"
	"log"
	"path/filepath"
	"runtime"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func Connect(url string) *pgxpool.Pool {
	pool, err := pgxpool.New(context.Background(), url)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}
	return pool
}

func Migrate(databaseURL string) {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)

	cleanPath := filepath.ToSlash(filepath.Join(basepath, "..", "..", "migrations"))
	migrationsPath := "file://" + cleanPath

	m, err := migrate.New(
		migrationsPath,
		databaseURL,
	)
	if err != nil {
		log.Fatalf("Migration initialization failed: %v", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Migration execution failed: %v", err)
	}

	log.Println("Database Migrated Successfully via golang-migrate!")
}
