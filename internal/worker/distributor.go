package worker

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/hibiken/asynq"
	// "github.com/joey1123455/beds-api/external/paystack"
	"github.com/joey1123455/beds-api/internal/logger"
	"github.com/joey1123455/beds-api/internal/notifier"
)

const (
	delayTime = 2 * time.Second
)

type TaskDistributor interface {
	// PaystackVirtualAccountNumberAssignment(
	// 	ctx context.Context,
	// 	payload *paystack.DedicatedAccountNumberAssignmentPayload,
	// 	opts ...asynq.Option,
	// ) error

	// Fire is a helper method that enqueues a task with given name and payload.
	// It is for generic jobs that are not specific to a particular domain.
	Fire(ctx context.Context, name string, payload any, opts ...asynq.Option) error

	SendNotification(ctx context.Context, msg *notifier.Notification, opts ...asynq.Option) error
}

type RedisTaskDistributor struct {
	client *asynq.Client
	logger logger.Logger
}

func (distributor *RedisTaskDistributor) Fire(ctx context.Context, name string, payload any, opts ...asynq.Option) error {
	return distributor.sendJob(ctx, name, payload, DispatchProcessInQueue(QueueNormal, opts...)...)
}

func (distributor *RedisTaskDistributor) SendNotification(ctx context.Context, msg *notifier.Notification, opts ...asynq.Option) error {
	return distributor.sendJob(ctx, SendNotification, msg, opts...)
}

func NewRedisTaskTaskDistributor(redisOpt asynq.RedisClientOpt, logger logger.Logger) TaskDistributor {
	client := asynq.NewClient(redisOpt)
	return &RedisTaskDistributor{
		client: client,
		logger: logger,
	}

}

func (distributor *RedisTaskDistributor) sendJob(ctx context.Context, name string, payload any, opts ...asynq.Option) error {

	taskPayload, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal task payload: %w", err)
	}

	task := asynq.NewTask(name, taskPayload, DispatchProcessInTimeWithQueue(delayTime, QueueNormal, opts...)...)
	info, err := distributor.client.EnqueueContext(ctx, task)
	if err != nil {
		return fmt.Errorf("unable to enqeueu task: %w", err)
	}

	distributor.logger.Info("Enqueued task", map[string]interface{}{
		"task_type": task.Type(),
		"queue":     info.Queue,
		"max_retry": info.MaxRetry,
	})

	return nil
}
