package usecase

import (
	"context"

	"github.com/InstayPMS/backend/internal/application/dto"
	"github.com/minio/minio-go/v7"
	"go.uber.org/zap"
)

type fileUseCaseImpl struct {
	client *minio.Client
	log *zap.Logger
}

func NewFileUseCase(client *minio.Client, log *zap.Logger) FileUseCase {
	return &fileUseCaseImpl{
		client,
		log,
	}
}

func (s *fileUseCaseImpl) CreateUploadURLs(ctx context.Context, req dto.UploadPresignedURLsRequest) ([]*dto.UploadPresignedURLResponse, error) {
	return nil, nil
}