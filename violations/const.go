package violations

type ViolationType string

const (
	// Deployment
	SINGLE_REPLICA_TYPE ViolationType = "SINGLE_REPLICA"
	// IMAGES
	IMAGE_REPO_TYPE ViolationType = "IMAGE_REPO"
	IMAGE_SIZE_TYPE ViolationType = "IMAGE_SIZE"
	// VOLUMES
	HOST_VOLUMES_TYPE ViolationType = "HOST_VOLUMES"
	// INGRESS
	INGRESS_HOST_INVALID_TYPE ViolationType = "INGRESS_HOST_INVALID"
	// CAPS
	PRIVILEGED_TYPE ViolationType = "PRIVILEGED"
	CAPABILITIES_TYPE ViolationType = "CAPABILITIES"
)