package worker

import (
	"time"

	"github.com/hibiken/asynq"
)

// Scheduler this is used to control the cron jobs
type Scheduler interface {
	Start() error
	Shutdown()
	RegisterJob(cronspec string, name string, payload []byte) (string, error)
	RemoveJob(name string) error
}

type RedisScheduler struct {
	server *asynq.Scheduler
}

func NewRedisScheduler(redisOpt asynq.RedisClientOpt, loc *time.Location) Scheduler {
	return &RedisScheduler{
		server: asynq.NewScheduler(redisOpt, &asynq.SchedulerOpts{
			Location: loc,
		}),
	}
}

func (r *RedisScheduler) Start() error {
	return r.server.Run()
}

func (r *RedisScheduler) Shutdown() {
	r.server.Shutdown()
}

func (r *RedisScheduler) RegisterJob(cronspec string, name string, payload []byte) (string, error) {
	return r.server.Register(cronspec, asynq.NewTask(name, payload))
}

func (r *RedisScheduler) RemoveJob(name string) error {
	return r.server.Unregister(name)
}
