package schedulers_interfaces

import "context"

type TestDataScheduler interface {
	AutoDeleteUnsavedTestData(ctx context.Context) error
}
