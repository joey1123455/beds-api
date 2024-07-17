package server

import (
	"context"
	"errors"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/hibiken/asynq"
	"github.com/joey1123455/beds-api/internal/config"
	db "github.com/joey1123455/beds-api/internal/db/sqlc"
	logg "github.com/joey1123455/beds-api/internal/logger"
	"github.com/joey1123455/beds-api/internal/mailer"
	"github.com/joey1123455/beds-api/internal/notifier"
	"github.com/joey1123455/beds-api/internal/worker"

	"os/signal"
	"syscall"
	"time"
)

type Server struct {
	Router          *gin.Engine
	Store           db.Store
	Config          *config.Config
	Logger          logg.Logger
	Mail            notifier.Notifier
	TaskProcessor   worker.TaskProcessor
	TaskDistributor worker.TaskDistributor
	Scheduler       worker.Scheduler
}

func NewServer(cfg *config.Config, store db.Store, redisOpt asynq.RedisClientOpt) (*Server, error) {

	loggerInstance := logg.NewZeroLogger(os.Stdout, logg.LevelInfo)

	taskDistributor := worker.NewRedisTaskTaskDistributor(redisOpt, loggerInstance)
	scheduler := worker.NewRedisScheduler(redisOpt, time.UTC)

	taskProcessor := worker.NewRedisTaskProcessor(
		redisOpt,
		store,
		loggerInstance,
		cfg,
		taskDistributor,
		time.UTC,
	)

	// mailer package initialization
	mail, err := mailer.NewMailer(*cfg)
	if err != nil {
		return nil, err
	}

	return &Server{
		Config:          cfg,
		Store:           store,
		Logger:          loggerInstance,
		Mail:            mail,
		TaskProcessor:   taskProcessor,
		TaskDistributor: taskDistributor,
		Scheduler:       scheduler,
	}, nil

}

func RunServer(srv *Server) {

	if srv == nil {
		srv.Logger.Fatal(errors.New("server instance cannot be nil"), nil)
	}

	go func() {
		err := srv.TaskProcessor.Start()
		if err != nil {
			srv.Logger.Fatal(err, nil)
		}
	}()

	go func() {
		err := srv.Scheduler.Start()
		if err != nil {
			srv.Logger.Fatal(err, nil)
		}
	}()

	// Setup server
	httpServer := &http.Server{
		Addr:         srv.Config.HttpAddress,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      srv.Router, // Pass our instance of gin in.
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			srv.Logger.Fatal(err, nil)
		}
	}()
	srv.Logger.Info("Server is up and running on", map[string]interface{}{
		"addr": srv.Config.HttpAddress,
	})

	quit := make(chan os.Signal, 1)
	// Accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Block until we receive our signal.
	<-quit

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	// Shutdown the server.
	if err := httpServer.Shutdown(ctx); err != nil {
		srv.Logger.Fatal(err, nil)
	}

}
