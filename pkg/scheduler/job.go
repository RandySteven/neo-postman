package scheduler

import "context"

type Job interface {
	RunAllJob(ctx context.Context) error
	autoDeleteUnsavedTestData(ctx context.Context) error
	autoCreatedTestRecord(ctx context.Context) error
}
