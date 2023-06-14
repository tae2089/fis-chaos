package domain

// ChaosMeshParameters defines the parameters for CHaosMesh
type ChaosMeshParameters struct {
	APIVersion string
	Kind       string
	Namespace  string
	Spec       string
	RoleARN    string
}

type ChaosReaderable interface {
	GetSec() int
	GetExperimentName() string
	GetSpec() string
}
