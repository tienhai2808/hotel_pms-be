package container

import (
	"github.com/InstaySystem/is_v2-be/internal/infrastructure/initialization"
	"github.com/InstaySystem/is_v2-be/internal/infrastructure/provider/jwt"
	"github.com/InstaySystem/is_v2-be/internal/infrastructure/provider/rabbitmq"
	"github.com/InstaySystem/is_v2-be/internal/infrastructure/provider/redis"
	"github.com/InstaySystem/is_v2-be/internal/infrastructure/provider/smtp"
)

func (c *Container) initCore() (err error) {
	c.Log, err = initialization.InitLog(c.cfg.Log)
	if err != nil {
		return err
	}

	c.DB, err = initialization.InitDatabase(c.cfg.PostgreSQL)
	if err != nil {
		return err
	}

	c.cache, err = initialization.InitCache(c.cfg.Redis)
	if err != nil {
		return err
	}

	c.stor, err = initialization.InitStorage(c.cfg.MinIO)
	if err != nil {
		return err
	}

	c.mq, err = initialization.InitMessageQueue(c.cfg.RabbitMQ)
	if err != nil {
		return err
	}

	c.IDGen, err = initialization.InitIDGen()
	if err != nil {
		return err
	}

	c.jwtPro = jwt.NewJWTProvider(c.cfg.JWT)

	c.cachePro = redis.NewCacheProvider(c.cache.Client())

	c.MQPro = rabbitmq.NewMessageQueueProvider(c.mq.Connection(), c.Log.Logger())

	c.SMTPPro = smtp.NewSMTPProvider(c.cfg.SMTPConfig)

	return nil
}
