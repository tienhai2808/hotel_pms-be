package worker

import (
	"encoding/json"

	"github.com/InstayPMS/backend/internal/application/dto"
	"github.com/InstayPMS/backend/internal/application/port"
	"github.com/InstayPMS/backend/pkg/constants"
	"go.uber.org/zap"
)

type MessageQueueWorker struct {
	mq   port.MessageQueueProvider
	smtp port.SMTPProvider
	log  *zap.Logger
}

func NewMessageQueueWorker(
	mq   port.MessageQueueProvider,
	smtp port.SMTPProvider,
	log  *zap.Logger,
) *MessageQueueWorker {
	return &MessageQueueWorker{
		mq,
		smtp,
		log,
	}
}

func (w *MessageQueueWorker) Start() {
	go w.startSendAuthEmail()
}

func (w *MessageQueueWorker) startSendAuthEmail() {
	if err := w.mq.ConsumeMessage(constants.QueueNameAuthEmail, constants.ExchangeEmail, constants.RoutingKeyAuthEmail, func(body []byte) error {
		var emailMsg dto.AuthEmailMessage
		if err := json.Unmarshal(body, &emailMsg); err != nil {
			return err
		}

		if err := w.smtp.AuthEmail(emailMsg.To, emailMsg.Subject, emailMsg.Otp); err != nil {
			return err
		}

		return nil
	}); err != nil {
		w.log.Error("start consumer send auth email failed", zap.Error(err))
	}
}
