package testutils

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joey1123455/beds-api/internal/common"
	"github.com/joey1123455/beds-api/internal/config"
	"github.com/joey1123455/beds-api/internal/logger"
)

const (
	databasePrefix      = "monieverse_"
	databasePlaceHolder = "[databaseName]"
)

func dataFromEnv(key, defaultValue string) string {
	v := os.Getenv(key)
	if v == "" && defaultValue != "" {
		return defaultValue
	}
	return v
}

var (
	databaseDefaultURL = dataFromEnv("TEST_BASE_DB_SOURCE", "postgresql://root:secret@localhost:5432/flux_core?sslmode=disable")
)

type TestSetup struct {
	Config    config.Config
	TestDB    *sql.DB
	Migration *common.Migration
}

func SetupTest(packageName string, configPath string, envFileName string) TestSetup {
	var err error

	cfg, err := config.LoadTest(configPath)
	if err != nil {
		log.Fatal(err)
	}

	testDatabaseName, err := CreateTestDatabase(packageName)
	if err != nil {
		if strings.Contains(err.Error(), "already exists") {
			err = DeleteTestDatabase(packageName)
			if err != nil {
				log.Fatal(err)
			}
			testDatabaseName, err = CreateTestDatabase(packageName)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			log.Fatal(err)
		}
	}

	cfg.DbUrl = dataFromEnv("TEST_DB_SOURCE", cfg.DbUrl)

	cfg.DbUrl = ReplaceDatabaseURLPlaceholder(cfg.DbUrl, testDatabaseName)
	testDB, err := sql.Open(cfg.DbType, cfg.DbUrl)
	if err != nil {
		log.Fatal(err)
	}

	err = testDB.Ping()
	if err != nil {
		log.Fatal(err)
	}

	zeroLogger := logger.NewZeroLogger(nil, logger.LevelOff)
	migration := common.NewMigration(zeroLogger, common.GetMigrationsPathAsURL(), cfg.DbUrl)
	err = migration.Up()
	if err != nil {
		println("migration error")
		log.Fatal(err)
	}

	return TestSetup{
		Config:    *cfg,
		TestDB:    testDB,
		Migration: migration,
	}
}

func TeardownTest(setup TestSetup, packageName string) {
	err := setup.Migration.Drop()
	if err != nil {
		log.Fatal(err)
	}

	err = setup.TestDB.Close()
	if err != nil {
		log.Fatal(err)
	}

	err = DeleteTestDatabase(packageName)
	if err != nil {
		log.Fatal(err)
	}
}

// ReplaceDatabaseURLPlaceholder replaces the placeholder in the database URL with the test database name.
func ReplaceDatabaseURLPlaceholder(databaseURL string, packageName string) string {
	return strings.Replace(databaseURL, databasePlaceHolder, packageName, 1)
}

func CreateTestDatabase(packageName string) (string, error) {
	dbName := databasePrefix + packageName
	createDBSQL := fmt.Sprintf("CREATE DATABASE  %s;", dbName)

	// Connect to the default "postgres" database to create a new database
	db, err := sql.Open("postgres", databaseDefaultURL)
	if err != nil {
		return "", err
	}
	defer db.Close()

	_, err = db.Exec(createDBSQL)
	if err != nil {
		return "", err
	}

	return dbName, nil
}

// DeleteTestDatabase deletes the test database.
func DeleteTestDatabase(packageName string) error {
	dbName := databasePrefix + packageName
	terminateConnectionsSQL := fmt.Sprintf("SELECT pg_terminate_backend(pg_stat_activity.pid) FROM pg_stat_activity WHERE pg_stat_activity.datname = '%s';", dbName)
	dropDBSQL := fmt.Sprintf("DROP DATABASE IF EXISTS %s;", dbName)

	// Connect to the default "postgres" database to drop the test database
	db, err := sql.Open("postgres", databaseDefaultURL)
	if err != nil {
		return err
	}
	defer db.Close()

	// Terminate active connections to the test database
	_, err = db.Exec(terminateConnectionsSQL)
	if err != nil {
		return err
	}

	_, err = db.Exec(dropDBSQL)
	if err != nil {
		return err
	}

	return nil
}
