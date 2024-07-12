package scheduler

import (
	"context"
	schedulers_interfaces "github.com/RandySteven/neo-postman/interfaces/schedulers"
	"github.com/RandySteven/neo-postman/pkg/postgres"
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
	}
)

func (s *scheduler) autoCreatedTestRecord(ctx context.Context) error {
	err := s.scheduler.AddFunc("@daily", func() {
		err := s.schedulerDependency.testReportScheduler.AutosaveTestRecords(ctx)
		if err != nil {
			return
		}
	})
	if err != nil {
		return err
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

func (s *scheduler) autoDeleteUnsavedTestData(ctx context.Context) error {
	err := s.scheduler.AddFunc("@daily", func() {
		err := s.schedulerDependency.testDataScheduler.AutoDeleteUnsavedTestData(ctx)
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

func NewScheduler(repo postgres.Repositories) *scheduler {
	jakartaTime, _ := time.LoadLocation("Asia/Jakarta")
	return &scheduler{
		scheduler: cron.NewWithLocation(jakartaTime),
		schedulerDependency: SchedulersDependency{
			testReportScheduler: schedulers.NewTestRecordScheduler(repo.TestRecordRepo),
			testDataScheduler:   schedulers.NewTestDataScheduler(repo.TestDataRepo),
		},
	}
}
