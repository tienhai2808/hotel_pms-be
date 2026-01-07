package di

import (
	fileUC "github.com/InstayPMS/backend/internal/application/usecase/file"
	"github.com/InstayPMS/backend/internal/infrastructure/api/http/handler"
	"github.com/minio/minio-go/v7"
	"go.uber.org/zap"
)

type Container struct {
	Log         *zap.Logger
	Storage     *minio.Client
	FileUseCase fileUC.FileUseCase
	FileHandler *handler.FileHandler
}

func NewContainer(stor *minio.Client) *Container {
	c := &Container{}

	c.initUseCases()

	c.initHandlers()

	return c
}
