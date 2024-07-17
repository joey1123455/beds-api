package common

import (
	"os"
	"path/filepath"
)

// GetProjectRoot return the current project root directory
func GetProjectRoot() string {
	curDir, err := os.Getwd()
	if err != nil {
		return ""
	}

	for {
		path := filepath.Join(curDir, "go.mod")

		_, err := os.Stat(path)
		if err == nil {
			return curDir
		}

		if !os.IsNotExist(err) {
			return ""
		}

		parentDir := filepath.Dir(curDir)
		if parentDir == curDir {
			break
		}
		curDir = parentDir
	}

	return curDir
}

// GetDBPath returns the path to the db directory
func GetDBPath() string {
	return filepath.Join(GetProjectRoot(), "db")
}

// GetMigrationsPath returns the path to the migrations directory
func GetMigrationsPath() string {
	return filepath.Join(GetProjectRoot(), "internal", "db", "migrations")
}

// GetMigrationsPathAsURL returns the path to the migrations directory as a URL
func GetMigrationsPathAsURL() string {
	return "file://" + GetMigrationsPath()
}
