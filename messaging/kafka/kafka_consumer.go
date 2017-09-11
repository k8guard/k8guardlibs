package kafka

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	libs "github.com/k8guard/k8guardlibs"

	"github.com/Shopify/sarama"
	"github.com/k8guard/k8guardlibs/config"
	"github.com/k8guard/k8guardlibs/messaging/types"
	log "github.com/sirupsen/logrus"
)

type kafkaConsumer struct {
	consumer sarama.Consumer
}

func NewConsumer(clientId types.ClientID, Cfg config.Config) (types.MessageConsumer, error) {

	brokers := Cfg.KafkaBrokers

	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.ClientID = string(clientId)

	master, err := sarama.NewConsumer(brokers, config)

	return &kafkaConsumer{consumer: master}, err
}

func (kc *kafkaConsumer) ConsumeMessages(messages chan []byte) {
	signalChannel := make(chan os.Signal, 2)
	signal.Notify(signalChannel, syscall.SIGTERM)

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	topic := libs.Cfg.KafkaActionTopic

	partitions, _ := kc.consumer.Partitions(topic)
	for _, partition := range partitions {

		libs.Log.Info("Creating Consumer ", topic, " on partition ", partition)
		consumer, err := kc.consumer.ConsumePartition(topic, partition, sarama.OffsetNewest)
		if err != nil {
			panic(err)
		}

		go func(consumer sarama.PartitionConsumer) {
			for {
				select {
				case err := <-consumer.Errors():
					fmt.Println(err)
				case msg := <-consumer.Messages():
					messages <- msg.Value
				case <-signals:
					libs.Log.Debug("Interrupt is detected")
				}
			}
		}(consumer)
	}
}

func (consumer *kafkaConsumer) Close() {
	log.Info("Closing kafka consumer")
	if err := consumer.consumer.Close(); err != nil {
		log.Fatal("Failed to close consumer: ", err)
	}
}
