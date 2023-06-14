package domain

import "encoding/json"

type StressChaosDto struct {
	Fis             `json:"fis" binding:"required"`
	StressChaosSpec `json:"chaosSpec"`
}

type StressChaosSpec struct {
	ContainerNames       []string              `json:"containerNames"`
	Duration             string                `json:"duration,omitempty"`
	Selector             map[string]string     `json:"selector" binding:"required"`
	StressChaosStressors *StressChaosStressors `json:"stressors,omitempty"`
	RemoteCluster        string                `json:"remoteCluster,omitempty"`
	Mode                 string                `json:"mode" binding:"required"`
}

type StressChaosStressors struct {
	*StressChaosCpu    `json:"cpu,omitempty"`
	*StressChaosMemory `json:"memory,omitempty"`
}

type StressChaosCpu struct {
	Load    int `json:"load"`
	Workers int `json:"workers"`
}

type StressChaosMemory struct {
	OomScoreAdj int    `json:"oom_score_adj"`
	Size        string `json:"size"`
	Workers     int    `json:"workers"`
}

func (d StressChaosDto) GetSec() int {
	return d.Sec
}

func (d StressChaosDto) GetExperimentName() string {
	return d.ExperimentName
}

func (d StressChaosDto) GetSpec() string {
	jsonbinary, err := json.Marshal(d.StressChaosSpec)
	if err != nil {
		return ""
	}
	return string(jsonbinary)
}
