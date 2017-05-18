package kafka

import (
	"encoding/json"
	"github.com/Shopify/sarama"
	log "github.com/sirupsen/logrus"
	"time"
	"github.com/k8guard/k8guardlibs/config"

	"strconv"
)

//implements KafkaProducer interface
type kafkaProducer struct {
	producer sarama.SyncProducer
}

type KafkaProducer interface {
	SendMessage(topic string, bytes []byte) error
	SendData(topic string, kind MessageType, message interface{}) error
	Close()
}

func NewProducer(clientId ClientID, Cfg config.Config) (KafkaProducer, error) {
	brokers := Cfg.KafkaBrokers
	config := sarama.NewConfig()
	config.Producer.Retry.Max = 5
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	config.ClientID = string(clientId)
	p, err := sarama.NewSyncProducer(brokers, config)

	return &kafkaProducer{producer: p}, err
}

func (producer *kafkaProducer) SendMessage(topic string, bytes []byte) error {
	strTime := strconv.Itoa(int(time.Now().Unix()))
	msg := &sarama.ProducerMessage{
		Topic:     topic,
		Key:	   sarama.StringEncoder(strTime),
		Value:     sarama.ByteEncoder(bytes),
	}
	_, _, err := producer.producer.SendMessage(msg)
	return err
}

func (producer *kafkaProducer) SendData(topic string, kind MessageType, message interface{}) error {
	message_data := map[string]interface{}{
		"kind": kind,
		"data": message,
	}
	bytes, err := json.Marshal(message_data)
	if err != nil {
		log.WithError(err).Error("Error Marshaling Kafka Data Message")
		return err
	}
	return  producer.SendMessage(topic, bytes)
}

func (producer *kafkaProducer) Close() {
	log.Info("Closing kafka producer")
	if err := producer.producer.Close(); err != nil {
		log.Fatal("Failed to close producer: ", err)
	}
}