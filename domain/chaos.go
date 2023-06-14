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

type ChaosSelectors struct {
	Annotations map[string]string `json:"annotationSelectors,omitempty"`
	Labels      map[string]string `json:"labelSelectors,omitempty"`
	Namespace   []string          `json:"namespaces,omitempty"`
	Nodes       map[string]string `json:"nodeSelectors,omitempty"`
}
