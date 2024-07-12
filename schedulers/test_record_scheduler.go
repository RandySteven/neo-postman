package schedulers

import (
	"context"
	repositories_interfaces "github.com/RandySteven/neo-postman/interfaces/repositories"
	schedulers_interfaces "github.com/RandySteven/neo-postman/interfaces/schedulers"
)

type testRecordScheduler struct {
	testRecordRepository repositories_interfaces.TestRecordRepository
}

func (t *testRecordScheduler) AutosaveTestRecords(ctx context.Context) error {
	err := t.testRecordRepository.SaveSavedTestData(ctx)
	if err != nil {
		return err
	}
	return nil
}

var _ schedulers_interfaces.TestRecordScheduler = &testRecordScheduler{}

func NewTestRecordScheduler(testRecordRepository repositories_interfaces.TestRecordRepository) *testRecordScheduler {
	return &testRecordScheduler{
		testRecordRepository: testRecordRepository,
	}
}
