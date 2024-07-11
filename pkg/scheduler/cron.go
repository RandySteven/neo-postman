package scheduler

import (
	"context"
)

type scheduler struct {
}

func (s *scheduler) RunAllJob(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (s *scheduler) deleteAllTestRecord(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

var _ Job = &scheduler{}

func NewScheduler() *scheduler {
	return &scheduler{}
}
