package schedulers

import (
	"context"
	repositories_interfaces "github.com/RandySteven/neo-postman/interfaces/repositories"
	schedulers_interfaces "github.com/RandySteven/neo-postman/interfaces/schedulers"
)

type testDataScheduler struct {
	testDataRepository repositories_interfaces.TestDataRepository
}

func (t *testDataScheduler) AutoDeleteUnsavedTestData(ctx context.Context) error {
	err := t.testDataRepository.DeletedUnsavedTestData(ctx)
	if err != nil {
		return err
	}
	return nil
}

var _ schedulers_interfaces.TestDataScheduler = &testDataScheduler{}

func NewTestDataScheduler(testDataRepository repositories_interfaces.TestDataRepository) *testDataScheduler {
	return &testDataScheduler{
		testDataRepository: testDataRepository,
	}
}
