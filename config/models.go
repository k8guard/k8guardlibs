package config

import "time"

type Config struct {
	Namespace               string   `env:"K8GUARD_NAMESPACE" envDefault:""`    // default is all namespaces
	ClusterName             string   `env:"K8GUARD_CLUSTER_NAME" envDefault:""` // the cluster name to be used in k8guard-action
	IgnoredNamespaces       []string `env:"K8GUARD_IGNORE_NAMESPACES" envSeparator:","`
	IgnoredPodsPrefix       []string `env:"K8GUARD_IGNORE_PODS_PREFIX" envSeparator:","`
	IgnoredDeployments      []string `env:"K8GUARD_IGNORE_DEPLOYMENTS" envSeparator:","`
	IgnoredDaemonSets       []string `env:"K8GUARD_IGNORE_DAEMONSETS" envSeparator:","`
	IgnoredJobs             []string `env:"K8GUARD_IGNORE_JOBS" envSeparator:","`
	IgnoredCronJobs         []string `env:"K8GUARD_IGNORE_CRONJOBS" envSeparator:","`
	ApprovedImageRepos      []string `env:"K8GUARD_APPROVED_IMAGE_REPOS" envSeparator:","`
	IngressMustContain      []string `env:"K8GUARD_INGRESS_MUST_CONTAIN" envSeparator:","`
	IngressMustNOTContain   []string `env:"K8GUARD_INGRESS_MUST_NOT_CONTAIN" envSeparator:","`
	ApprovedIngressSuffixes []string `env:"K8GUARD_APPROVED_INGRESS_SUFFIXES" envSeparator:","`
	ApprovedImageSize       int64    `env:"K8GUARD_APPROVED_IMAGE_SIZE"`
	OutputPodsToFile        bool     `env:"K8GUARD_OUTPUT_PODS_TO_FILE"`
	IgnoredViolations       []string `env:"K8GUARD_IGNORED_VIOLATIONS" envSeparator:"," `
	IncludeAlpha            bool     `env:"K8GUARD_INCLUDE_ALPHA" envDefault:"false"`

	RequiredEntities    []string `env:"K8GUARD_REQUIRED_ENTITIES" envSeparator:"," `
	RequiredAnnotations []string `env:"K8GUARD_REQUIRED_ANNOTATIONS" envSeparator:"," `
	RequiredLabels      []string `env:"K8GUARD_REQUIRED_LABELS" envSeparator:"," `

	CacheType string `env:"K8GUARD_CACHE_TYPE" envDefault:"MEMCACHED"`

	// CacheExpirationSeconds int32 `env:"K8GUARD_CACHE_EXPIRATION_SECONDS"`
	MemCachedHostname string `env:"K8GUARD_MEMCACHED_HOSTNAME"`
	LogLevel          string `env:"K8GUARD_LOG_LEVEL"`

	MessageBroker string `env:"K8GUARD_MESSAGE_BROKER" envDefault:"KAFKA"`

	KafkaBrokers      []string `env:"K8GUARD_KAFKA_BROKERS" envSeparator:","`
	KafkaCertFilePath string   `env:"K8GUARD_KAFKA_CERT_FILE_PATH"`
	KafkaKeyFilePath  string   `env:"K8GUARD_KAFKA_KEY_FILE_PATH"`
	KafkaKeyPassword  string   `env:"K8GUARD_KAFKA_KEY_PASSWORD"`
	KafkaActionTopic  string   `env:"K8GUARD_KAFKA_ACTION_TOPIC"`
	KafkaEventTopic   string   `env:"K8GUARD_KAFKA_EVENT_TOPIC"`

	RmqBroker      string `env:"K8GUARD_RMQ_BROKER" envDefault:"redis:6379"`
	RmqActionTopic string `env:"K8GUARD_RMQ_ACTION_TOPIC"`
	RmqEventTopic  string `env:"K8GUARD_RMQ_EVENT_TOPIC"`

	// Action Specific Configs
	CassandraHosts             []string `env:"K8GUARD_ACTION_CASSANDRA_HOSTS" envSeparator:","`
	CassandraKeyspace          string   `env:"K8GUARD_ACTION_CASSANDRA_KEYSPACE"`
	CassandraUsername          string   `env:"K8GUARD_ACTION_CASSANDRA_USERNAME"`
	CassandraPassword          string   `env:"K8GUARD_ACTION_CASSANDRA_PASSWORD"`
	CassandraCaPath            string   `env:"K8GUARD_ACTION_CASSANDRA_CAPATH"`
	CassandraSslHostValidation bool     `env:"K8GUARD_ACTION_CASSANDRA_SSL_HOST_VALIDATION" envDefault:"true"`
	CassandraCreateKeyspace    bool     `env:"K8GUARD_CASSANDRA_CREATE_KEYSPACE" envDefault:"true"`
	CassandraCreateTables      bool     `env:"K8GUARD_CASSANDRA_CREATE_TABLES" envDefault:"true"`

	HipchatToken   string `env:"K8GUARD_ACTION_HIPCHAT_TOKEN"`
	HipchatRoomID  string `env:"K8GUARD_ACTION_HIPCHAT_ROOM_ID"`
	HipchatBaseURL string `env:"K8GUARD_ACTION_HIPCHAT_BASE_URL"`
	//Tag the Namespace owner in hipchat
	HipchatTagNamespaceOwner bool `env:"K8GUARD_ACTION_HIPCHAT_TAG_NAMESPACE_OWNER"`

	// The format for the field in which to expect chat ids for namespace owners.
	AnnotationFormatForChatIds string `env:"K8GUARD_ACTION_ANNOTATION_FORMAT_FOR_CHAT_IDS" envSeparator:"," envDefault:"team/hipchat-ids"`
	AnnotationFormatForEmails  string `env:"K8GUARD_ACTION_ANNOTATION_FORMAT_FOR_EMAILS" envSeparator:"," envDefault:"team/email-ids"`

	DurationBetweenChatNotifications time.Duration `env:"K8GUARD_ACTION_DURATION_BETWEEN_CHAT_NOTIFICATIONS"`
	DurationBetweenNotifyingAgain    time.Duration `env:"K8GUARD_ACTION_DURATION_BETWEEN_NOTIFYING_AGAIN"`
	// This means after this duration without any new violation, we expect the violation to either be fixed or hard action to be done on it.
	DurationViolationExpires time.Duration `env:"K8GUARD_ACTION_DURATION_VIOLATION_EXPIRES"`
	// Parse messages from kafka and dont do anything
	ActionDryRun bool `env:"K8GUARD_ACTION_DRY_RUN"`

	// Parse messages, Notify but don't do any hard action such as scaling down or delete.
	ActionSafeMode           bool   `env:"K8GUARD_ACTION_SAFE_MODE"`
	WarningCountBeforeAction int    `env:"K8GUARD_ACTION_WARNING_COUNT_BEFORE_ACTION"`
	SmtpServer               string `env:"K8GUARD_ACTION_SMTP_SERVER"`
	SmtpPort                 int    `env:"K8GUARD_ACTION_SMTP_PORT"`
	SmtpUsername             string `env:"K8GUARD_ACTION_SMTP_USERNAME"`
	SmtpPassword             string `env:"K8GUARD_ACTION_SMTP_PASSWORD"`
	SmtpSendFrom             string `env:"K8GUARD_ACTION_SMTP_SEND_FROM"`
	// Email to send to if namespaces doesn't have an emails annotation
	SmtpFallbackSendTo string `env:"K8GUARD_ACTION_SMTP_FALLBACK_SEND_TO"`
	// Email the namespace owner, if false email fallback
	SmtpSendToNamespaceOwner bool `env:"K8GUARD_ACTION_SMTP_SEND_TO_NAMESAPCE_OWNER"`
	// Optional footer for emails, e.g. to include links to more information
	ViolationEmailFooter string `env:"K8GUARD_ACTION_VIOLATION_EMAIL_FOOTER"`
}
