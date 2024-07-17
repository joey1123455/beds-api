package worker

import (
	"errors"

	"github.com/hibiken/asynq"
	"github.com/joey1123455/beds-api/internal/logger"
)

type taskLogger struct {
	logger logger.Logger
}

func (t taskLogger) Debug(args ...interface{}) {
	t.logger.Info(args[0].(string), map[string]interface{}{})
}

func (t taskLogger) Info(args ...interface{}) {
	t.logger.Info(args[0].(string), nil)
}

func (t taskLogger) Warn(args ...interface{}) {
	t.logger.Info(args[0].(string), nil)
}

func (t taskLogger) Error(args ...interface{}) {
	t.logger.Error(errors.New(args[0].(string)), nil)
}

func (t taskLogger) Fatal(args ...interface{}) {
	t.logger.Fatal(errors.New(args[0].(string)), nil)
}

// newTaskLogger returns a new task logger
func newTaskLogger(logger logger.Logger) asynq.Logger {
	return &taskLogger{logger: logger}
}
