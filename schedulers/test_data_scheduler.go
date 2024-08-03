package schedulers

import (
	"context"
	caches_interfaces "github.com/RandySteven/neo-postman/interfaces/caches"
	repositories_interfaces "github.com/RandySteven/neo-postman/interfaces/repositories"
	schedulers_interfaces "github.com/RandySteven/neo-postman/interfaces/schedulers"
	"sync"
)

type testDataScheduler struct {
	testDataRepository repositories_interfaces.TestDataRepository
	testDataCache      caches_interfaces.TestDataCache
}

func (t *testDataScheduler) AutoDeleteUnsavedTestData(ctx context.Context) error {
	var (
		wg    sync.WaitGroup
		errCh = make(chan error)
	)

	wg.Add(2)

	go func() {
		defer wg.Done()
		if err := t.testDataCache.Del(ctx, "all.test_datas"); err != nil {
			errCh <- err
			return
		}
	}()

	go func() {
		defer wg.Done()
		err := t.testDataRepository.DeletedUnsavedTestData(ctx)
		if err != nil {
			errCh <- err
			return
		}
	}()

	wg.Wait()

	select {
	case err := <-errCh:
		return err
	default:
		return nil
	}
}

var _ schedulers_interfaces.TestDataScheduler = &testDataScheduler{}

func NewTestDataScheduler(
	testDataRepository repositories_interfaces.TestDataRepository,
	testDataCache caches_interfaces.TestDataCache,
) *testDataScheduler {
	return &testDataScheduler{
		testDataRepository: testDataRepository,
		testDataCache:      testDataCache,
	}
}
