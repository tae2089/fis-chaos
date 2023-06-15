package domain

import "github.com/tae2089/fis-chaos/types"

type PodStressSpec struct {
	Action         string              `json:"action" binding:"required"`
	ContainerNames []string            `json:"containerNames" binding:"required"`
	GracePeriod    int64               `json:"gracePeriod,omitempty"`
	Mode           types.ChaosMeshMode `json:"mode" binding:"required"`
	Selector       ChaosSelectors      `json:"selector" binding:"required"`
	Value          string              `json:"value,omitempty"`
}
