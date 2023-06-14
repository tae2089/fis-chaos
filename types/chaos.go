package types

// type ChaosMeshKind int
const (
	Awschaos             = "AWSChaos"
	Azurechaos           = "Azurechaos"
	Blockchaos           = "Blockchaos"
	Dnschaos             = "DNSChaos"
	Gcpchaos             = "GCPChaos"
	Httpchaos            = "HTTPChaos"
	Iochaos              = "IOChaos"
	Jvmchaos             = "JVMChaos"
	Kernelchaos          = "KernelChaos"
	Networkchaos         = "NetworkChaos"
	Physicalmachinechaos = "PhysicalMachineChaos"
	Physicalmachines     = "Physicalmachines"
	Podchaos             = "PodChaos"
	Podhttpchaos         = "PodHttpChaos"
	Podiochaos           = "PodIOChaos"
	Podnetworkchaos      = "PodNetworkChaos"
	Remoteclusters       = "Remoteclusters"
	Schedules            = "Schedule"
	Statuschecks         = "Statuschecks"
	Stresschaos          = "StressChaos"
	Timechaos            = "TimeChaos"
	ChaosEngine          = "ChaosEngine"
)

// AwsFisActionID is the type for the AWS FIS action ID
type AwsFisActionID int

// AWS FIS action IDs
const (
	InjectKubernetesCustomResource string = "aws:eks:inject-kubernetes-custom-resource"
)

// ChaosMeshApiVersion is the API version of the ChaosMesh API
const (
	ChaosMeshAPIVersionV1alpha1 string = "chaos-mesh.org/v1alpha1"
)
