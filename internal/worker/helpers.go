package worker

import (
	"time"

	"github.com/hibiken/asynq"
)

const (
	MaxRetry = 1
)

func DispatchWithOptions(opt ...asynq.Option) []asynq.Option {
	options := []asynq.Option{
		asynq.Queue(QueueNormal),
	}
	return append(options, opt...)
}

func DispatchWithDefaultOptions() []asynq.Option {
	return DispatchWithOptions(
		asynq.MaxRetry(MaxRetry),
	)
}

func DispatchProcessInTime(t time.Duration) []asynq.Option {
	return DispatchWithOptions(
		asynq.MaxRetry(MaxRetry),
		asynq.ProcessIn(t),
	)
}

func DispatchProcessInQueue(queue string, opt ...asynq.Option) []asynq.Option {
	options := []asynq.Option{
		asynq.MaxRetry(MaxRetry),
		asynq.Queue(queue),
	}
	return append(RemoveOptions(options), opt...)

}

// DispatchProcessInTimeWithQueue dispatches a task to be processed in a given time and queue
func DispatchProcessInTimeWithQueue(t time.Duration, queue string, opt ...asynq.Option) []asynq.Option {
	options := []asynq.Option{
		asynq.MaxRetry(MaxRetry),
		asynq.ProcessIn(t),
		asynq.Queue(queue),
	}
	return append(options, opt...)
}

// DispatchInCriticalQueue dispatches a task to be processed in the critical queue
func DispatchInCriticalQueue() []asynq.Option {

	options := RemoveOptions(DispatchWithOptions(), asynq.Queue(QueueNormal))
	opts := []asynq.Option{
		asynq.Queue(QueueCritical),
	}
	return append(options, opts...)
}

// RemoveOptions options from a slice of options
func RemoveOptions(opts []asynq.Option, opt ...asynq.Option) []asynq.Option {
	var newOpts []asynq.Option
	for _, o := range opts {
		if !contains(opt, o) {
			newOpts = append(newOpts, o)
		}
	}
	return newOpts
}

func contains(opt []asynq.Option, o asynq.Option) bool {
	for _, v := range opt {
		if v == o {
			return true
		}
	}
	return false
}
