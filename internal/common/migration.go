package common

import (
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	log "github.com/joey1123455/beds-api/internal/logger"
	_ "github.com/lib/pq"
)

// Migration configure database migration
type Migration struct {
	logger       log.Logger
	migrationURL string
	dbSource     string
}

// NewMigration creates a new migration instance
func NewMigration(logger log.Logger, migrationURL, dbSource string) *Migration {
	return &Migration{
		logger:       logger,
		migrationURL: migrationURL,
		dbSource:     dbSource,
	}
}

// Up run a forward migration
func (m *Migration) Up() error {
	migration, err := migrate.New(m.migrationURL, m.dbSource)

	if err != nil {
		newErr := fmt.Errorf("migration error %w", err)
		m.logger.Error(newErr, nil)
		return newErr
	}

	if err = migration.Up(); err != nil && err != migrate.ErrNoChange {
		newErr := fmt.Errorf("migration error %w", err)
		m.logger.Error(newErr, nil)
		return newErr
	}
	m.logger.Info("migration successful", nil)
	return nil
}

// Down run a reverse migration
func (m *Migration) Down() error {
	migration, err := migrate.New(m.migrationURL, m.dbSource)

	if err != nil {
		newErr := fmt.Errorf("migration error %w", err)
		m.logger.Error(newErr, nil)
		return newErr
	}

	if err = migration.Down(); err != nil && err != migrate.ErrNoChange {
		newErr := fmt.Errorf("migration error %w", err)
		m.logger.Error(newErr, nil)
		return newErr
	}

	m.logger.Info("migration successful", nil)
	return nil
}

// Drop  all tables
func (m *Migration) Drop() error {
	migration, err := migrate.New(m.migrationURL, m.dbSource)

	if err != nil {
		newErr := fmt.Errorf("migration error %w", err)
		m.logger.Error(newErr, nil)
		return newErr
	}

	if err = migration.Drop(); err != nil && err != migrate.ErrNoChange {
		newErr := fmt.Errorf("migration error %w", err)
		m.logger.Error(newErr, nil)
		return newErr
	}

	m.logger.Info("migration successful", nil)
	return nil
}
