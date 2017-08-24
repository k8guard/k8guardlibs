package rmq

import (
	"time"

	"github.com/adjust/rmq"
	"github.com/k8guard/k8guardlibs/config"
	"github.com/k8guard/k8guardlibs/messaging/types"
	log "github.com/sirupsen/logrus"
)

const (
	unackedLimit = 1000
	numConsumers = 10
	batchSize    = 1000
)

type rmqConsumer struct {
	consumer rmq.Queue
	messages chan []byte
}

func NewConsumer(clientId types.ClientID, Cfg config.Config) (types.MessageConsumer, error) {
	connection := rmq.OpenConnection("my service", "tcp", "localhost:6379", 1)
	queue := connection.OpenQueue("my queue")
	queue.StartConsuming(unackedLimit, 1*time.Second)

	return &rmqConsumer{consumer: queue}, nil
}

func (rc *rmqConsumer) ConsumeMessages(messages chan []byte) {
	rc.messages = messages
	rc.consumer.AddConsumer("consumer", rc)
}

func (rc *rmqConsumer) Consume(delivery rmq.Delivery) {
	rc.messages <- []byte(delivery.Payload())
}

func (rc *rmqConsumer) Close() {
	log.Info("Closing rmq consumer")
	rc.consumer.Close()
}
