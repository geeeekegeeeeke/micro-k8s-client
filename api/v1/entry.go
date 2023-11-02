package v1

import "gin-dubbogo-consumer/service"

type ApiGroup struct {
	BaseApi
}

var ApiGroupApp = new(ApiGroup)

var (
	containerService = service.NewIDockerService()
	k8sService       = service.NewIK8sClusterService()
)
