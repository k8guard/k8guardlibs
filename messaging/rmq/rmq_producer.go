package rmq

import (
	"encoding/json"

	"github.com/adjust/rmq"
	libs "github.com/k8guard/k8guardlibs"
	"github.com/k8guard/k8guardlibs/config"
	"github.com/k8guard/k8guardlibs/messaging/types"
)

//implements MessageProducer interface
type rmqProducer struct {
	producer rmq.Queue
}

func NewProducer(clientID types.ClientID, Cfg config.Config) (types.MessageProducer, error) {
	topic := libs.Cfg.RmqActionTopic
	broker := libs.Cfg.RmqBroker
	connection := rmq.OpenConnection("redis", "tcp", broker, 1)
	queue := connection.OpenQueue(topic)

	return &rmqProducer{producer: queue}, nil
}

func (producer *rmqProducer) SendData(kind types.MessageType, message interface{}) error {
	message_data := map[string]interface{}{
		"kind": kind,
		"data": message,
	}
	libs.Log.Debug("Sending %v", message_data)
	bytes, err := json.Marshal(message_data)
	if err != nil {
		libs.Log.WithError(err).Error("Error Marshaling Rmq Data Message")
		return err
	}
	producer.producer.PublishBytes(bytes)
	return nil
}

func (producer *rmqProducer) Close() {
	libs.Log.Info("Closing rmq producer")
	producer.producer.Close()
}
