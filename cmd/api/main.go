package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/hibiken/asynq"
	"github.com/joey1123455/beds-api/core/router"
	"github.com/joey1123455/beds-api/core/server"
	"github.com/joey1123455/beds-api/internal/config"
	db "github.com/joey1123455/beds-api/internal/db/sqlc"
	"github.com/joey1123455/beds-api/internal/logger"
	_ "github.com/lib/pq"
)

func main() {
	// Load config
	cfg, err := config.Load("../..")

	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// Connect to the database
	conn, err := sql.Open("postgres", cfg.DbUrl)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
		return
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatalf("Error closing database: %v", err)
		}
	}(conn)

	// Test the connection
	err = conn.Ping()
	if err != nil {
		log.Fatalf("Error pinging database: %v", err)
		return
	}

	fmt.Println("Connected to the database!")

	store := db.NewStore(conn)

	redisOpt := asynq.RedisClientOpt{
		Addr:     cfg.RedisAddress,
		Username: cfg.RedisUsername,
		Password: cfg.RedisPassword,
		DB:       cfg.RedisDB,
	}

	srv, err := server.NewServer(cfg, store, redisOpt)
	if err != nil {
		logger.ErrorLogger(err)
	}

	router.SetupRouter(srv)
	server.RunServer(srv)
}
