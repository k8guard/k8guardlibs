package kafka

import "github.com/k8guard/k8guardlibs/messaging/types"

const (
	ACTION_CLIENTID   types.ClientID = "k8guard-action-kafka-client"
	DISCOVER_CLIENTID types.ClientID = "k8guard-discover-kafka-client"

	// Event consumer feature will be implemented in the future.
	EVENT_CONSUMER_CLIENTID types.ClientID = "k8guard-event-consumer-kafka-client"
	EVENT_PARSER_CLIENTID   types.ClientID = "k8guard-event-parser-kafka-client"
)
