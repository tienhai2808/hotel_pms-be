package initialization

import (
	"fmt"

	"github.com/InstayPMS/backend/internal/infrastructure/config"
	"github.com/rabbitmq/amqp091-go"
)

type MQ struct {
	Conn *amqp091.Connection
	Chan *amqp091.Channel
}

func InitRabbitMQ(cfg config.RabbitMQ) (*MQ, error) {
	protocol := "amqp"
	if cfg.UseSSL {
		protocol = protocol + "s"
	}

	dsn := fmt.Sprintf("%s://%s:%s@%s:%d/%s",
		protocol,
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Vhost,
	)

	conn, err := amqp091.Dial(dsn)
	if err != nil {
		return nil, err
	}

	chann, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return &MQ{
		conn,
		chann,
	}, nil
}

func (mq *MQ) Close() {
	_ = mq.Chan.Close()
	_ = mq.Conn.Close()
}
