package domain

import "github.com/tae2089/fis-chaos/types"

type KernalChaosSpec struct {
	ContainerNames  []string       `json:"containerNames" binding:"required"`
	Duration        string         `json:"duration,omitempty"`
	Selector        ChaosSelectors `json:"selector" binding:"required"`
	FailKernRequest `json:"failKernRequest" binding:"required"`
	RemoteCluster   string              `json:"remoteCluster,omitempty"`
	Mode            types.ChaosMeshMode `json:"mode" binding:"required"`
	Value           string              `json:"value,omitempty"`
}

type FailKernRequest struct {
	CallChain   *[]KernalCallChain `json:"callchain,omitempty"`
	FailType    int                `json:"failtype" binding:"required"`
	Headers     []string           `json:"headers,omitempty"`
	Probability int                `json:"probability,omitempty"`
	Times       int                `json:"times,omitempty"`
}

type KernalCallChain struct {
	Funcname   string `json:"funcname,omitempty"`
	Parameters string `json:"parameters,omitempty"`
	Predicate  string `json:"predicate,omitempty"`
}
