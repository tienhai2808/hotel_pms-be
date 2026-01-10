package port

type MessageQueueProvider interface {
	PublishMessage(exchange, routingKey string, body []byte) error
	
	ConsumeMessage(queueName, exchange, routingKey string, handler func([]byte) error) error
}
