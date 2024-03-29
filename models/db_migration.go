package models

import (
	"database/sql"
	"fmt"
	"io/fs"

	"github.com/pressly/goose/v3"
)

func Migrate(db *sql.DB, dir string) error {
	err := goose.SetDialect("postgres")
	if err != nil {
		return fmt.Errorf("set dialect: %w", err)
	}
	err = goose.Up(db, dir)
	if err != nil {
		return fmt.Errorf("up migrations: %w", err)
	}
	return nil
}

func MigrateFromFS(db *sql.DB, migrationsFS fs.FS, dir string) error {
	if dir == "" {
		dir = "."
	}
	goose.SetBaseFS(migrationsFS)
	defer func() {
		goose.SetBaseFS(nil) // roll back migrations from FS
	}()
	return Migrate(db, dir)
}
