package worker

import (
	"github.com/hibiken/asynq"
)

func (processor *RedisTaskProcessor) registerTasks(mux *asynq.ServeMux) {
	mux.HandleFunc(SendNotification, processor.SendNotification)
}
