package domain

import "github.com/tae2089/fis-chaos/types"

type TimeChaosSpec struct {
	// ClockIds defines all affected clock id All available options are
	ClockIDs []types.TimeChaosClockID `json:"clockIds, omitempty"`

	// ContainerNames indicates list of the name of affected container. If not set, the first container will be injected
	ContainerNames []string `json:"containerNames"`

	// Duration represents the duration of the chaos action
	Duration string `json:"duration,omitempty"`

	// Mode defines the mode to run chaos action. Supported mode: one / all / fixed / fixed-percent / random-max-percent
	Mode types.ChaosMeshMode `json:"mode" binding:"required"`

	// RemoteCluster represents the remote cluster where the chaos will be deployed
	RemoteCluster string `json:"remoteCluster,omitempty"`

	// Selector is used to select pods that are used to inject chaos action.
	Selector ChaosSelectors `json:"selector" binding:"required"`

	//     TimeOffset defines the delta time of injected program. It's a possibly
	// signed sequence of decimal numbers, such as "300ms", "-1.5h" or "2h45m".
	// Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".
	TimeOffset string `json:"timeOffset,omitempty"`
	Value      string `json:"value,omitempty"`
}
