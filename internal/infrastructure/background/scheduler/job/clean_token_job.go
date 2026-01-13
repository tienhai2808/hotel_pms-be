package job

import (
	"context"
	"time"

	"github.com/InstaySystem/is_v2-be/internal/domain/repository"
	"go.uber.org/zap"
)

type cleanTokenJob struct {
	log       *zap.Logger
	tokenRepo repository.TokenRepository
}

func NewCleanTokenJob(
	log *zap.Logger,
	tokenRepo repository.TokenRepository,
) Job {
	return &cleanTokenJob{
		log,
		tokenRepo,
	}
}

func (j *cleanTokenJob) Name() string {
	return "cleanup_expired_tokens"
}

func (j *cleanTokenJob) Run() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	startTime := time.Now()

	rowsDeleted, err := j.tokenRepo.DeleteAllExpired(ctx)
	if err != nil {
		j.log.Error("delete all expired tokens failed", zap.Error(err))
		return
	}

	duration := time.Since(startTime)

	j.log.Info(
		"Cleanup expired tokens completed",
		zap.Int64("deleted_count", rowsDeleted),
		zap.Duration("duration", duration),
	)
}
