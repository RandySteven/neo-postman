package scheduler

import (
	"context"
	"time"
)

type Job interface {
	RunAllJob(ctx context.Context) error
	autoDeleteUnsavedTestData(ctx context.Context) error
	autoCreatedTestRecord(ctx context.Context) error
	refreshRedis(ctx context.Context, duration time.Duration) error
}
