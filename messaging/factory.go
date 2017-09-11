package messaging

import (
	"errors"

	"github.com/k8guard/k8guardlibs/config"
	"github.com/k8guard/k8guardlibs/messaging/kafka"
	"github.com/k8guard/k8guardlibs/messaging/rmq"
	"github.com/k8guard/k8guardlibs/messaging/types"
)

func CreateMessageProducer(p types.MessageBrokerType, clientId types.ClientID, cfg config.Config) (types.MessageProducer, error) {
	switch p {
	case types.KAFKA_BROKER:
		return kafka.NewProducer(clientId, cfg)
	case types.RMQ_BROKER:
		return rmq.NewProducer(clientId, cfg)
	default:
		return nil, errors.New("Invalid MessageProducer Type")
	}
}

func CreateMessageConsumer(p types.MessageBrokerType, clientId types.ClientID, cfg config.Config) (types.MessageConsumer, error) {
	switch p {
	case types.KAFKA_BROKER:
		return kafka.NewConsumer(clientId, cfg)
	case types.RMQ_BROKER:
		return rmq.NewConsumer(clientId, cfg)
	default:
		return nil, errors.New("Invalid MessageConsumer Type")
	}
}
