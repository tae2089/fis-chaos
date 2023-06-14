# fis-chaos

## Description

This repository is a server that creates and executes templates using AWS FIS and Chaos Mesh.

## Usage

Usage in the local environment:

Copy the .env.development file and create a new file called .env.
We determine whether to use the .env file or not based on the environment variable APP_ENV. If you want to use the .env file, simply set the value of APP_ENV to either blank or "local".

Usage in the development, release(production) environment:

If you set environment values such as development, release, or production, they will be configured based on the environment variable values maintained in the operating system.

Run the code.

```go
package main

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tae2089/fis-chaos/api/router"
	"github.com/tae2089/fis-chaos/config"
)

func main() {
	config.LoadEnv()
	timeout := time.Duration(3) * time.Second
	gin := gin.Default()
	router.SetUp(timeout, gin)
	gin.Run(":8080")
}


```

## Error state on AWS FIS experiment
**Unable to get custom resource: Unexpected status.**
> AWS FIS can't find the third party chaos tool.
- You need to install chaos-mesh in the cluster.

**Action exceeded the specified maximum duration.**
> AWS FIS cant't find the corresponding resources you specify in the experiment.
- Check the target arns, tags, etc.
> The execution time is more then the `maxDuration` you specify in the experiment.
- Reduce the `duration` in the field `kubernetesSpec` if the action type is `aws:eks:inject-kubernetes-custom-resource`.
- It need a few second for the third party chaos tool to deploy the service and inject the chaos. It will be failed if the `duration` is equal or more than the `maxDuration` you specify in the experiment.

**Unable to inject custom resource: Kubernetes API returned status code 422. Please check EKS logs for more details.**
> The JSON data in the field `kubernetesSpec` is insufficient if the action type is `aws:eks:inject-kubernetes-custom-resource`.
- Check the JSON format and the the required fields.

**Unable to inject custom resource: Not authorized to perform the required action.**
> Action type `aws:eks:inject-kubernetes-custom-resource` need rbac permission in your cluster.
- You need to setup the IAM Role the configmap `aws-auth` in the namespace `kube-system`
- You can also use the following command to modify the setting manually.
```bash
# Use the command in your EKS Cluster
kubectl edit -n kube-system configmap/aws-auth
```

## Author

👤 tae2089 tae2089002@gmail.com

Github: @tae2089

## 🤝 Contributing

Contributions, issues and feature requests are welcome!
Feel free to check issues page.


## Reference
[fis-chaos-mesh](https://github.com/BalinLin)