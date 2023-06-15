package domain

type PodStressSpec struct {
	Action         string         `json:"action" binding:"required"`
	ContainerNames []string       `json:"containerNames" binding:"required"`
	GracePeriod    int64          `json:"gracePeriod,omitempty"`
	Mode           string         `json:"mode" binding:"required"`
	Selector       ChaosSelectors `json:"selector" binding:"required"`
}
