package schedulers_interfaces

import (
	"context"
)

type TestRecordScheduler interface {
	AutosaveTestRecords(ctx context.Context) error
}
