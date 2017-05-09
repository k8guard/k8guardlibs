package kafka

import (
	"github.com/k8guard/k8guardlibs/config"
	"github.com/Shopify/sarama"
)

func NewConsumer(clientId ClientID, Cfg config.Config) (sarama.Consumer, error) {

	brokers := Cfg.KafkaBrokers

	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.ClientID = string(clientId)

	master, err := sarama.NewConsumer(brokers, config)
	return master, err
}