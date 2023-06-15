package domain

import (
	"encoding/json"

	"github.com/tae2089/fis-chaos/types"
)

type PodChaosDto struct {
	Fis          `json:"fis" binding:"required"`
	PodChaosSpec `json:"chaosSpec"`
}

type PodChaosSpec struct {
	Action         string              `json:"action" binding:"required"`
	ContainerNames []string            `json:"containerNames" binding:"required"`
	GracePeriod    int64               `json:"gracePeriod,omitempty"`
	Mode           types.ChaosMeshMode `json:"mode" binding:"required"`
	Selector       ChaosSelectors      `json:"selector" binding:"required"`
	Value          string              `json:"value,omitempty"`
}

func (d PodChaosDto) GetSec() int {
	return d.Sec
}

func (d PodChaosDto) GetExperimentName() string {
	return d.ExperimentName
}

func (d PodChaosDto) GetSpec() string {
	jsonbinary, err := json.Marshal(d.PodChaosSpec)
	if err != nil {
		return ""
	}
	return string(jsonbinary)
}
