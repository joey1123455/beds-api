package worker

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hibiken/asynq"
	"github.com/joey1123455/beds-api/internal/notifier"
)

func (processor *RedisTaskProcessor) SendNotification(_ context.Context, task *asynq.Task) error {

	msg := &notifier.Notification{}

	if err := json.Unmarshal(task.Payload(), &msg); err != nil {
		processor.logger.Error(err, nil)
		return fmt.Errorf("failed to unmarshal task payload: %w", asynq.SkipRetry)
	}
	err := processor.notifier.SendNotification(msg)
	if err != nil {
		processor.logger.Error(fmt.Errorf("error sending notification %s", err), map[string]interface{}{
			"msg": msg,
		})
		return err
	}

	return nil
}
