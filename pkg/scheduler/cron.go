package scheduler

import (
	"context"
	"github.com/RandySteven/neo-postman/apps"
	usecases_interfaces "github.com/RandySteven/neo-postman/interfaces/usecases"
	"github.com/robfig/cron"
	"log"
	"time"
)

type (
	scheduler struct {
		scheduler          *cron.Cron
		testCaseDependency TestCaseDependency
	}

	TestCaseDependency struct {
		testCaseUsecase usecases_interfaces.TestDataUsecase
	}
)

func (s *scheduler) RunAllJob(ctx context.Context) (err error) {
	err = s.deleteAllTestRecord(ctx)
	if err != nil {
		return err
	}
	go s.scheduler.Start()
	return nil
}

func (s *scheduler) deleteAllTestRecord(ctx context.Context) error {
	err := s.scheduler.AddFunc("@daily", func() {
		err := s.testCaseDependency.testCaseUsecase.AutoDeleteUnsavedRecord(ctx)
		if err != nil {
			log.Printf("failed to delete unsaved record: %v", err)
			return
		}
	})
	if err != nil {
		return err
	}
	return nil
}

var _ Job = &scheduler{}

func NewScheduler(usecases apps.Usecases) *scheduler {
	jakartaTime, _ := time.LoadLocation("Asia/Jakarta")
	return &scheduler{
		scheduler:          cron.NewWithLocation(jakartaTime),
		testCaseDependency: TestCaseDependency{testCaseUsecase: usecases.TestDataUsecase},
	}
}
