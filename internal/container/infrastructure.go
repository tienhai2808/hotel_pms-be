package container

import (
	"github.com/InstayPMS/backend/internal/infrastructure/api/http/middleware"
	"github.com/InstayPMS/backend/internal/infrastructure/config"
	"github.com/InstayPMS/backend/internal/infrastructure/initialization"
)

func (c *Container) initInfrastructure(cfg *config.Config) error {
	log, err := initialization.InitZap(cfg)
	if err != nil {
		return err
	}
	c.Log = log

	db, err := initialization.InitDatabase(cfg)
	if err != nil {
		return err
	}
	c.DB = db

	rdb, err := initialization.InitRedis(cfg)
	if err != nil {
		return err
	}
	c.Cache = rdb

	stor, err := initialization.InitMinIO(cfg)
	if err != nil {
		return err
	}
	c.Stor = stor

	idGen, err := initialization.InitSnowFlake()
	if err != nil {
		return err
	}
	c.IDGen = idGen

	ctxMid := middleware.NewContextMiddleware(log)
	c.CtxMid = ctxMid

	return nil
}
