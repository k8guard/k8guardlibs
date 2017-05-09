package k8guardlibs

import (
	"github.com/k8guard/k8guardlibs/violations"
)

// Any entity from the kubernetes world
type Entity struct {
	Name string
}

// Entity that can potentially have violations
type ViolatableEntity struct {
	Entity
	Namespace  string
	Cluster    string
	Violations []violations.Violation
}

type Ingress struct {
	ViolatableEntity
}

type Deployment struct {
	ViolatableEntity
}

type Job struct {
	ViolatableEntity
}

type CronJob struct {
	ViolatableEntity
}

type Pod struct {
	ViolatableEntity
}

var (
	Version string
	Build   string
)
