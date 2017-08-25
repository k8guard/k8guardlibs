package rmq

import (
	"encoding/json"

	"github.com/adjust/rmq"
	"github.com/k8guard/k8guardlibs/config"
	"github.com/k8guard/k8guardlibs/messaging/types"
	log "github.com/sirupsen/logrus"
)

//implements MessageProducer interface
type rmqProducer struct {
	producer rmq.Queue
}

func NewProducer(clientId types.ClientID, Cfg config.Config) (types.MessageProducer, error) {
	topic := libs.Cfg.RmqActionTopic
	connection := rmq.OpenConnection("redis", "tcp", "redis:6379", 1)
	queue := connection.OpenQueue(topic)

	return &rmqProducer{producer: queue}, nil
}

func (producer *rmqProducer) SendData(kind types.MessageType, message interface{}) error {
	message_data := map[string]interface{}{
		"kind": kind,
		"data": message,
	}
	bytes, err := json.Marshal(message_data)
	if err != nil {
		log.WithError(err).Error("Error Marshaling Rmq Data Message")
		return err
	}
	producer.producer.PublishBytes(bytes)
	return nil
}

func (producer *rmqProducer) Close() {
	log.Info("Closing rmq producer")
	producer.producer.Close()
}
