package kafka

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/Shopify/sarama"
	lib "github.com/k8guard/k8guardlibs"
	"github.com/k8guard/k8guardlibs/config"
	"github.com/k8guard/k8guardlibs/messaging/types"
	log "github.com/sirupsen/logrus"
)

//implements KafkaProducer interface
type kafkaProducer struct {
	producer sarama.SyncProducer
}

func NewProducer(clientId types.ClientID, cfg config.Config) (types.MessageProducer, error) {
	brokers := cfg.KafkaBrokers
	config := sarama.NewConfig()
	config.Producer.Retry.Max = 5
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	config.ClientID = string(clientId)
	p, err := sarama.NewSyncProducer(brokers, config)

	return &kafkaProducer{producer: p}, err
}

func (producer *kafkaProducer) sendMessage(topic string, bytes []byte) error {
	strTime := strconv.Itoa(int(time.Now().Unix()))
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder(strTime),
		Value: sarama.ByteEncoder(bytes),
	}
	_, _, err := producer.producer.SendMessage(msg)
	return err
}

func (producer *kafkaProducer) SendData(kind types.MessageType, message interface{}) error {
	message_data := map[string]interface{}{
		"kind": kind,
		"data": message,
	}
	bytes, err := json.Marshal(message_data)
	if err != nil {
		log.WithError(err).Error("Error Marshaling Kafka Data Message")
		return err
	}
	return producer.sendMessage(lib.Cfg.KafkaActionTopic, bytes)
}

func (producer *kafkaProducer) Close() {
	log.Info("Closing kafka producer")
	if err := producer.producer.Close(); err != nil {
		log.Fatal("Failed to close producer: ", err)
	}
}
