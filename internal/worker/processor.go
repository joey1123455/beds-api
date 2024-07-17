package worker

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/hibiken/asynq"
	"github.com/joey1123455/beds-api/internal/config"
	db "github.com/joey1123455/beds-api/internal/db/sqlc"
	"github.com/joey1123455/beds-api/internal/logger"
	"github.com/joey1123455/beds-api/internal/mailer"
	"github.com/joey1123455/beds-api/internal/notifier"
)

var (
	ErrResourceNotAvailable = errors.New("no resource is available")
)

const (
	QueueCritical       = "critical"
	QueueDefault        = "default"
	QueueNormal         = "normal"
	QueuePriorityHigher = 10
	QueuePriorityLower  = 5
	QueuePriorityNormal = 1
)

type TaskProcessor interface {
	Start() error
	Shutdown()
}

type RedisTaskProcessor struct {
	server          *asynq.Server
	store           db.Store
	logger          logger.Logger
	config          *config.Config
	taskDistributor TaskDistributor
	timezone        *time.Location
	notifier        *notifier.NotificationSender
}

func NewRedisTaskProcessor(
	redisOpt asynq.RedisClientOpt,
	store db.Store,
	logger logger.Logger,
	cfg *config.Config,
	taskDistributor TaskDistributor,
	timezone *time.Location,
) TaskProcessor {
	mail, err := mailer.NewMailer(*cfg)
	if err != nil {
		return nil
	}
	notificationSender := notifier.NewNotificationSender(mail)

	rd := &RedisTaskProcessor{
		server: asynq.NewServer(redisOpt, asynq.Config{
			Queues: map[string]int{
				QueueCritical: QueuePriorityHigher,
				QueueDefault:  QueuePriorityLower,
				QueueNormal:   QueuePriorityNormal,
			},
			StrictPriority: true,
			Logger:         newTaskLogger(logger),
			IsFailure: func(err error) bool {
				return err != ErrResourceNotAvailable
			},
		}),
		store:           store,
		logger:          logger,
		notifier:        notificationSender,
		config:          cfg,
		taskDistributor: taskDistributor,
		timezone:        timezone,
	}

	return rd
}

func (processor *RedisTaskProcessor) Start() error {
	mux := asynq.NewServeMux()
	processor.registerTasks(mux)
	return processor.server.Start(mux)
}

func (processor *RedisTaskProcessor) Shutdown() {
	processor.logger.Info("shutting down task processor", nil)
	processor.server.Shutdown()
	processor.logger.Info("task processor shutdown", nil)
}

func (processor *RedisTaskProcessor) sendNotification(
	ctx context.Context,
	to notifier.Recipient,
	title, templateName string,
	data map[string]any,
	channels []string,
) {
	notification := notifier.NewNotification(to, title, templateName, channels)
	for k, v := range data {
		notification.AddData(k, v)
	}

	if err := processor.taskDistributor.SendNotification(ctx, notification); err != nil {
		processor.logger.Error(fmt.Errorf("error sending notification %s", err), map[string]interface{}{
			"notification": notification,
		})
	}

	processor.logger.Info(fmt.Sprintf("notification sent to %s", to), map[string]interface{}{
		"channel": channels,
	})
}
