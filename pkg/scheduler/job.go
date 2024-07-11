package scheduler

import "context"

type Job interface {
	RunAllJob(ctx context.Context) error
	deleteAllTestRecord(ctx context.Context) error
}
