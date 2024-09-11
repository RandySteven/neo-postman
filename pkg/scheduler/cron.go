package scheduler

import (
	"context"
	"fmt"
	schedulers_interfaces "github.com/RandySteven/neo-postman/interfaces/schedulers"
	"github.com/RandySteven/neo-postman/pkg/postgres"
	"github.com/RandySteven/neo-postman/pkg/redis"
	"github.com/RandySteven/neo-postman/schedulers"
	"github.com/robfig/cron"
	"time"
)

type (
	scheduler struct {
		scheduler           *cron.Cron
		schedulerDependency SchedulersDependency
	}

	SchedulersDependency struct {
		testDataScheduler   schedulers_interfaces.TestDataScheduler
		testReportScheduler schedulers_interfaces.TestRecordScheduler
		redisClient         *redis.RedisClient
	}
)

func (s *scheduler) refreshRedis(ctx context.Context, duration time.Duration) error {
	err := s.schedulerDependency.redisClient.FlushAll(ctx)
	if err != nil {
		return fmt.Errorf("failed to clear cache: %w", err)
	}
	return nil
}

func (s *scheduler) RunAllJob(ctx context.Context) (err error) {
	err = s.autoDeleteUnsavedTestData(ctx)
	if err != nil {
		return err
	}
	err = s.autoCreatedTestRecord(ctx)
	if err != nil {
		return err
	}
	go s.scheduler.Start()
	return nil
}

func (s *scheduler) autoCreatedTestRecord(ctx context.Context) error {
	return s.runScheduler(ctx, "@daily", s.schedulerDependency.testReportScheduler.AutosaveTestRecords)
}

func (s *scheduler) autoDeleteUnsavedTestData(ctx context.Context) error {
	return s.runScheduler(ctx, "@daily", s.schedulerDependency.testDataScheduler.AutoDeleteUnsavedTestData)
}

func (s *scheduler) runScheduler(ctx context.Context, spec string, schedulerFunc func(ctx context.Context) error) error {
	err := s.scheduler.AddFunc(spec, func() {
		err := schedulerFunc(ctx)
		if err != nil {
			return
		}
	})
	if err != nil {
		return err
	}
	return nil
}

var _ Job = &scheduler{}

func NewScheduler(repo postgres.Repositories, cache *redis.RedisClient) *scheduler {
	jakartaTime, _ := time.LoadLocation("Asia/Jakarta")
	return &scheduler{
		scheduler: cron.NewWithLocation(jakartaTime),
		schedulerDependency: SchedulersDependency{
			testReportScheduler: schedulers.NewTestRecordScheduler(repo.TestRecordRepo),
			testDataScheduler:   schedulers.NewTestDataScheduler(repo.TestDataRepo, cache.TestDataCache),
			redisClient:         cache,
		},
	}
}
