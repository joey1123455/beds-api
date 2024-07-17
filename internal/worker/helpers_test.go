package worker

import (
	"testing"

	"github.com/hibiken/asynq"
	"github.com/stretchr/testify/assert"
)

func TestDispatchProcessInQueue(t *testing.T) {

	got := DispatchProcessInQueue("test")
	want := []asynq.Option{
		asynq.MaxRetry(MaxRetry),
		asynq.Queue("test"),
	}

	assert.Equal(t, got, want)
}

func TestRemoveOptions(t *testing.T) {

	got := RemoveOptions([]asynq.Option{asynq.MaxRetry(1), asynq.Queue("test")}, asynq.MaxRetry(1))
	want := []asynq.Option{
		asynq.Queue("test"),
	}

	assert.Equal(t, got, want)
}

func TestDispatchWithOptions(t *testing.T) {

	got := DispatchWithOptions(asynq.MaxRetry(1))
	want := []asynq.Option{
		asynq.Queue(QueueNormal),
		asynq.MaxRetry(1),
	}

	assert.Equal(t, got, want)
}

func TestDispatchWithDefaultOptions(t *testing.T) {

	got := DispatchWithDefaultOptions()
	want := []asynq.Option{
		asynq.Queue(QueueNormal),
		asynq.MaxRetry(MaxRetry),
	}

	assert.Equal(t, got, want)
}

func TestDispatchProcessInTime(t *testing.T) {

	got := DispatchProcessInTime(1)
	want := []asynq.Option{
		asynq.Queue(QueueNormal),
		asynq.MaxRetry(MaxRetry),
		asynq.ProcessIn(1),
	}

	assert.Equal(t, got, want)
}

func TestDispatchProcessInTimeWithQueue(t *testing.T) {

	got := DispatchProcessInTimeWithQueue(1, "test")
	want := []asynq.Option{
		asynq.MaxRetry(MaxRetry),
		asynq.ProcessIn(1),
		asynq.Queue("test"),
	}

	assert.Equal(t, got, want)
}

func TestDispatchInCriticalQueue(t *testing.T) {

	got := DispatchInCriticalQueue()
	want := []asynq.Option{
		asynq.Queue(QueueCritical),
	}

	assert.Equal(t, got, want)
}
