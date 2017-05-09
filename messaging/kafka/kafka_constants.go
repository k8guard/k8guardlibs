package kafka

type MessageType string

const (
	POD_MESSAGE MessageType = "pod"
	DEPLOYMENT_MESSAGE MessageType = "deployment"
	INGRESS_MESSAGE MessageType = "ingress"
	JOB_MESSAGE MessageType = "job"
	CRONJOB_MESSAGE MessageType = "cronjob"
)

type ClientID string

const (
	ACTION_CLIENTID ClientID = "k8guard-action-kafka-client"
	DISCOVER_CLIENTID ClientID = "k8guard-discover-kafka-client"
	EVENT_CONSUMER_CLIENTID ClientID = "k8guard-event-consumer-kafka-client"
	EVENT_PARSER_CLIENTID ClientID = "k8guard-event-parser-kafka-client"
)
