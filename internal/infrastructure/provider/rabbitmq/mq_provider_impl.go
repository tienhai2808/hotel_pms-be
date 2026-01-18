package rabbitmq

import (
	"context"
	"fmt"
	"math"
	"time"

	"github.com/InstaySystem/is_v2-be/internal/application/port"
	"github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
)

type messageQueueProviderImpl struct {
	rmq *amqp091.Connection
	log *zap.Logger
}

func NewMessageQueueProvider(
	rmq *amqp091.Connection,
	log *zap.Logger,
) port.MessageQueueProvider {
	return &messageQueueProviderImpl{
		rmq,
		log,
	}
}

func (m *messageQueueProviderImpl) PublishMessage(exchange, routingKey string, body []byte) error {
	ch, err := m.rmq.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := ch.PublishWithContext(ctx, exchange, routingKey, false, false, amqp091.Publishing{
		ContentType:  "application/json",
		DeliveryMode: amqp091.Persistent,
		Body:         body,
	}); err != nil {
		return err
	}

	return nil
}

func (m *messageQueueProviderImpl) ConsumeMessage(queueName, exchange, routingKey string, handler func([]byte) error) error {
	ch, err := m.rmq.Channel()
	if err != nil {
		return err
	}

	if _, err := ch.QueueDeclare(queueName, true, false, false, false, nil); err != nil {
		return err
	}

	if err := ch.ExchangeDeclare(exchange, "direct", true, false, false, false, nil); err != nil {
		return err
	}

	if err := ch.QueueBind(queueName, routingKey, exchange, false, nil); err != nil {
		return err
	}

	if err := ch.Qos(5, 0, false); err != nil {
		return err
	}

	msgs, err := ch.Consume(queueName, "", true, false, false, false, nil)
	if err != nil {
		return err
	}

	for i := range 5 {
		go func(workerID int) {
			for msg := range msgs {
				m.processWithRetry(msg.Body, handler, workerID)
			}
		}(i)
	}

	return nil
}

func (m *messageQueueProviderImpl) processWithRetry(body []byte, handler func([]byte) error, workerID int) {
	maxAttempts := 5
	initialInterval := 1000 * time.Millisecond
	multiplier := 2.0
	maxInterval := 10000 * time.Millisecond

	for attempt := 1; attempt <= maxAttempts; attempt++ {
		err := handler(body)
		if err == nil {
			return
		}
		m.log.Error(fmt.Sprintf("work %d (%d/%d) failed", workerID, attempt, maxAttempts), zap.Error(err))

		if attempt < maxAttempts {
			delay := float64(initialInterval) * math.Pow(multiplier, float64(attempt-1))
			if delay > float64(maxInterval) {
				delay = float64(maxInterval)
			}
			time.Sleep(time.Duration(delay))
		}
	}

	m.log.Error(fmt.Sprintf("work %d", workerID), zap.Error(fmt.Errorf("message sending failed after %d attempts", maxAttempts)))
}
