package initialization

import (
	"fmt"

	"github.com/InstaySystem/is_v2-be/internal/infrastructure/config"
	"github.com/rabbitmq/amqp091-go"
)

type MessageQueue struct {
	conn *amqp091.Connection
}

func InitMessageQueue(cfg config.RabbitMQ) (*MessageQueue, error) {
	protocol := "amqp"
	if cfg.UseSSL {
		protocol += "s"
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

	return &MessageQueue{
		conn,
	}, nil
}

func (m *MessageQueue) Close() {
	_ = m.conn.Close()
}

func (m *MessageQueue) Connection() *amqp091.Connection {
	return m.conn
}
