package usecase

import (
	"context"

	"github.com/InstayPMS/backend/internal/application/dto"
)

type FileUseCase interface {
	CreateUploadURLs(ctx context.Context, req dto.UploadPresignedURLsRequest) ([]*dto.UploadPresignedURLResponse, error)
}
