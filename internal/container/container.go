package container

import (
	authUC "github.com/InstayPMS/backend/internal/application/usecase/auth"
	fileUC "github.com/InstayPMS/backend/internal/application/usecase/file"
	"github.com/InstayPMS/backend/internal/infrastructure/api/http/handler"
	"github.com/InstayPMS/backend/internal/infrastructure/api/http/middleware"
	"github.com/InstayPMS/backend/internal/infrastructure/config"
	"github.com/InstayPMS/backend/internal/infrastructure/initialization"
	"github.com/minio/minio-go/v7"
	"github.com/redis/go-redis/v9"
	"github.com/sony/sonyflake/v2"
	"go.uber.org/zap"
)

type Container struct {
	Cfg     *config.Config
	Log     *zap.Logger
	DB      *initialization.Database
	Cache   *redis.Client
	Stor    *minio.Client
	IDGen   *sonyflake.Sonyflake
	FileUC  fileUC.FileUseCase
	AuthUC  authUC.AuthUseCase
	FileHdl *handler.FileHandler
	AuthHdl *handler.AuthHandler
	CtxMid  *middleware.ContextMiddleware
}

func NewContainer(cfg *config.Config) (*Container, error) {
	c := &Container{
		Cfg: cfg,
	}

	if err := c.initInfrastructure(cfg); err != nil {
		return nil, err
	}

	c.initUseCases()

	c.initHandlers()

	return c, nil
}

func (c *Container) Cleanup() {
	if c.DB != nil {
		c.DB.Close()
	}
}
