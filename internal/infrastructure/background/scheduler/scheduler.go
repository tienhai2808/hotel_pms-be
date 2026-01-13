package scheduler

import (
	"github.com/InstaySystem/is_v2-be/internal/infrastructure/background/scheduler/job"
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
)

type Scheduler struct {
	log  *zap.Logger
	cron *cron.Cron
}

func NewScheduler(log *zap.Logger) *Scheduler {
	cron := cron.New()

	return &Scheduler{
		log,
		cron,
	}
}

func (s *Scheduler) AddJob(schedule string, job job.Job) error {
	if _, err := s.cron.AddFunc(schedule, func() {
		s.log.Info("Running scheduled", zap.String("job", job.Name()))
		job.Run()
	}); err != nil {
		return err
	}

	s.log.Info("Job scheduled", zap.String("job", job.Name()))
	return nil
}

func (s *Scheduler) Start() {
	s.cron.Start()
}

func (s *Scheduler) Stop() {
	s.cron.Stop()
}
