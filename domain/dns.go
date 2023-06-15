package domain

import "github.com/tae2089/fis-chaos/types"

type DnsChaosSpec struct {
	Action         string              `json:"action" binding:"required"`
	ContainerNames []string            `json:"containerNames"`
	Duration       string              `json:"duration,omitempty"`
	Selector       ChaosSelectors      `json:"selector" binding:"required"`
	RemoteCluster  string              `json:"remoteCluster,omitempty"`
	Mode           types.ChaosMeshMode `json:"mode" binding:"required"`
	Patterns       []string            `json:"patterns,omitempty"`
	Value          string              `json:"value,omitempty"`
}
