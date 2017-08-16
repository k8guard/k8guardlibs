package kafka

type MessageType string

const (
	NAMESPACE_MESSAGE  MessageType = "namespace"
	POD_MESSAGE        MessageType = "pod"
	DEPLOYMENT_MESSAGE MessageType = "deployment"
	DAEMONSET_MESSAGE  MessageType = "daemonset"
	INGRESS_MESSAGE    MessageType = "ingress"
	JOB_MESSAGE        MessageType = "job"
	CRONJOB_MESSAGE    MessageType = "cronjob"
	// To be used to see if kafka topic is there.
	TEST_MESSAGE MessageType = "test"
)

type ClientID string

const (
	ACTION_CLIENTID   ClientID = "k8guard-action-kafka-client"
	DISCOVER_CLIENTID ClientID = "k8guard-discover-kafka-client"

	// Event consumer feature will be implemented in the future.
	EVENT_CONSUMER_CLIENTID ClientID = "k8guard-event-consumer-kafka-client"
	EVENT_PARSER_CLIENTID   ClientID = "k8guard-event-parser-kafka-client"
)
