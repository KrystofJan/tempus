package db

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/KrystofJan/tempus/internal/constants"
	"github.com/golang-migrate/migrate/v4"
)

func MigrateUp() error {
	env := constants.EnvironmentType(os.Getenv("ENVIRONMENT_TYPE"))
	if env == "" {
		env = constants.Development
	}

	connStr, err := GetConnString()
	if err != nil {
		return err
	}

	wd, _ := os.Getwd()

	schemaDir := filepath.Join(wd, "internal", "migrations", "schema")
	envDir := filepath.Join(wd, "internal", "migrations", string(env))

	tempDir, err := os.MkdirTemp("", "combined_migrations_*")
	if err != nil {
		return err
	}
	defer os.RemoveAll(tempDir)

	if err := copyFiles(schemaDir, tempDir); err != nil {
		return fmt.Errorf("failed copying schema migrations: %w", err)
	}

	if _, err := os.Stat(envDir); err == nil {
		if err := copyFiles(envDir, tempDir); err != nil {
			return fmt.Errorf("failed copying env migrations: %w", err)
		}
		fmt.Printf("Using %s migrations.\n", env)
	} else {
		fmt.Printf("No specific migrations for %s.\n", env)
	}

	migrationSource := "file://" + tempDir
	return runMigrations(connStr, migrationSource)
}

func copyFiles(srcDir, dstDir string) error {
	return filepath.Walk(srcDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		dstPath := filepath.Join(dstDir, info.Name())
		return copyFile(path, dstPath)
	})
}

func copyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	return err
}

func runMigrations(connStr, path string) error {
	m, err := migrate.New(path, connStr)
	if err != nil {
		return fmt.Errorf("failed to create migration: %w", err)
	}
	defer m.Close()

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("migration error: %w", err)
	}
	return nil
}
