package types

type ClientID string

type MessageType string

type MessageBrokerType string

const (
	// supported message brokers
	KAFKA_BROKER MessageBrokerType = "KAFKA"
	RMQ_BROKER   MessageBrokerType = "RMQ"

	// message types
	NAMESPACE_MESSAGE  MessageType = "namespace"
	POD_MESSAGE        MessageType = "pod"
	DEPLOYMENT_MESSAGE MessageType = "deployment"
	DAEMONSET_MESSAGE  MessageType = "daemonset"
	INGRESS_MESSAGE    MessageType = "ingress"
	JOB_MESSAGE        MessageType = "job"
	CRONJOB_MESSAGE    MessageType = "cronjob"
	TEST_MESSAGE       MessageType = "test"
)

type MessageProducer interface {
	SendData(kind MessageType, message interface{}) error
	Close()
}

type MessageConsumer interface {
	ConsumeMessages(messages chan []byte)
	Close()
}
