package v1

import "micro-k8s-client/service"

type ApiGroup struct {
	BaseApi
}

var ApiGroupApp = new(ApiGroup)

var (
	containerService = service.NewIDockerService()
	k8sService       = service.NewIK8sClusterService()
)
