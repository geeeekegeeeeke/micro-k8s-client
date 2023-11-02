package router

import (
	"gin-dubbogo-consumer/api/v1"
	"github.com/gin-gonic/gin"
)

//var Router *gin.Engine

type DockerRouter struct {
}

func (a *K8sRouter) InitDockerRouter(Router *gin.RouterGroup) {
	/*	Router = gin.New()

		//Dobubo-Demo-API接口
		var baseApi v1.UserController
		var baseApi v1.ClientController
		var baseApi v1.NodeController
		var svcctl v1.ServiceController
		var applyctl v1.ApplyController
		var baseApi v1.IngressController
		var baseApi v1.StorageController
		var baseApi v1.DcokerController
		var baseApi v1.YmlDeployController
		//var baseApi controllers.DcokerDeployController
		var dockerDeployctl v1.DcokerDeployController
		baseApi := Router.Group("/v1/resource")*/
	baseRouter := Router.Group("")
	baseApi := v1.ApiGroupApp.BaseApi
	{
		/*baseRouter.GET("/user/*id", baseApi.User)
		baseRouter.GET("/users", baseApi.Users)
		baseRouter.POST("/user", baseApi.Store)
		baseRouter.PUT("/user", baseApi.Update)
		baseRouter.DELETE("/user", baseApi.Destroy)*/

		//baseRouter.GET("/pod", baseApi.OperateDeploy)
		baseRouter.GET("/pod/deploy", baseApi.ListDeploy)
		baseRouter.GET("/pod", baseApi.GetNodeInfo)
		baseRouter.POST("/pod/deploy", baseApi.CreateDeploy)
		baseRouter.PUT("/pod/deploy", baseApi.UpdateDeploy)
		baseRouter.GET("/pod/image", baseApi.GetImage)
		baseRouter.GET("/pod/container", baseApi.GetContainer)
		baseRouter.GET("/pod/component", baseApi.GetComponent)
		//baseRouter.GET("/pod/deploy", baseApi.ListDeploy)
		baseRouter.GET("/node", baseApi.ListNode)
		baseRouter.GET("/node/info", baseApi.GetNodeInfo)
		//baseRouter.GET("/svc/account", baseApi.lise)
		//baseRouter.GET("/svc", baseApi.ListService)
		baseRouter.GET("/apply/config", baseApi.ConfigMap)
		baseRouter.GET("/apply/namespace", baseApi.Namespace)
		baseRouter.PUT("/apply/secret", baseApi.SetSecret)
		baseRouter.GET("/apply/secret", baseApi.Secret)
		baseRouter.GET("/apply/tomcat", baseApi.DeployTomcatApp)
		//baseRouter.GET("/apply", applyctl.GetApply)
		baseRouter.GET("/ingress/info", baseApi.ListIngress)
		baseRouter.GET("/ingress", baseApi.ListIngressInfo)
		baseRouter.GET("/storage", baseApi.ListPersistent)
		baseRouter.GET("/storage/vol", baseApi.ListPersistentVol)
		//baseRouter.GET("/dockerdeploy", baseApi.DeployAppBydirectCompose)
		baseRouter.POST("/k8s/ymldeploy", baseApi.YmlDeploy)
		//
		//baseRouter.GET("/docker/list", baseApi.ListContainer)
		//baseRouter.GET("/docker/search", baseApi.SearchContainer)
		//baseRouter.GET("/docker/stats", baseApi.ContainerStats)
		//baseRouter.GET("/docker/log", baseApi.ContainerLogs)
		//baseRouter.GET("/docker/rename", baseApi.ContainerOperation)
		//baseRouter.GET("/docker/exec", baseApi.ContainerInfo)
		//baseRouter.GET("/docker/", baseApi.ContainerCreate)
		//baseRouter.GET("/storage", baseApi.ListIngressInfo)
		//baseRouter.GET("/exec", baseApi.ContainerWsSsh)  todo
		baseRouter.GET("/docker/stats/:id", baseApi.ContainerStats)

		baseRouter.POST("/docker", baseApi.ContainerCreate)
		baseRouter.POST("/docker/update", baseApi.ContainerUpdate)
		baseRouter.POST("/docker/upgrade", baseApi.ContainerUpgrade)
		baseRouter.POST("/docker/info", baseApi.ContainerInfo)
		baseRouter.POST("/docker/search", baseApi.SearchContainer)
		baseRouter.POST("/docker/list", baseApi.ListContainer)
		baseRouter.GET("/docker/list/stats", baseApi.ContainerStats)
		baseRouter.GET("/docker/search/log", baseApi.ContainerLogs)
		baseRouter.GET("/docker/limit", baseApi.LoadResouceLimit)
		baseRouter.POST("/docker/clean/log", baseApi.CleanContainerLog)
		baseRouter.POST("/docker/load/log", baseApi.LoadContainerLog)
		baseRouter.POST("/docker/inspect", baseApi.Inspect)
		baseRouter.POST("/docker/operate", baseApi.ContainerOperation)
		baseRouter.POST("/docker/prune", baseApi.ContainerPrune)

		//

		//baseRouter.GET("/image/pull", baseApi.ImagePull)
		//baseRouter.GET("/image/push", baseApi.ImagePush)
		baseRouter.GET("/image", baseApi.ListImage)
		baseRouter.POST("/image/search", baseApi.SearchImage)
		baseRouter.POST("/image/pull", baseApi.ImagePull)
		baseRouter.POST("/image/push", baseApi.ImagePush)
		baseRouter.POST("/image/save", baseApi.ImageSave)
		baseRouter.POST("/image/load", baseApi.ImageLoad)
		baseRouter.POST("/image/remove", baseApi.ImageRemove)
		baseRouter.POST("/image/tag", baseApi.ImageTag)
		baseRouter.POST("/image/build", baseApi.ImageBuild)

		baseRouter.GET("/repo", baseApi.ListRepo)
		baseRouter.POST("/repo/status", baseApi.CheckRepoStatus)
		baseRouter.POST("/repo/search", baseApi.SearchRepo)
		baseRouter.POST("/repo/update", baseApi.UpdateRepo)
		baseRouter.POST("/repo", baseApi.CreateRepo)
		//baseRouter.POST("/repo/del", baseApi.DELETERepo)

		//
		baseRouter.GET("/volume", baseApi.ListVolume)
	}
	/*baseRouter.GET("/exec", baseApi.ContainerWsSsh)
	baseRouter.GET("/stats/:id", baseApi.ContainerStats)

	baseRouter.POST("", baseApi.ContainerCreate)
	baseRouter.POST("/update", baseApi.ContainerUpdate)
	baseRouter.POST("/upgrade", baseApi.ContainerUpgrade)
	baseRouter.POST("/info", baseApi.ContainerInfo)
	baseRouter.POST("/search", baseApi.SearchContainer)
	baseRouter.POST("/list", baseApi.ListContainer)
	baseRouter.GET("/list/stats", baseApi.ContainerListStats)
	baseRouter.GET("/search/log", baseApi.ContainerLogs)
	baseRouter.GET("/limit", baseApi.LoadResourceLimit)
	baseRouter.POST("/clean/log", baseApi.CleanContainerLog)
	baseRouter.POST("/load/log", baseApi.LoadContainerLog)
	baseRouter.POST("/inspect", baseApi.Inspect)
	baseRouter.POST("/operate", baseApi.ContainerOperation)
	baseRouter.POST("/prune", baseApi.ContainerPrune)

	baseRouter.GET("/repo", baseApi.ListRepo)
	baseRouter.POST("/repo/status", baseApi.CheckRepoStatus)
	baseRouter.POST("/repo/search", baseApi.SearchRepo)
	baseRouter.POST("/repo/update", baseApi.UpdateRepo)
	baseRouter.POST("/repo", baseApi.CreateRepo)
	baseRouter.POST("/repo/del", baseRouter.DELETERepo)

	baseRouter.POST("/compose/search", baseApi.SearchCompose)
	baseRouter.POST("/compose", baseApi.CreateCompose)
	baseRouter.POST("/compose/test", baseApi.TestCompose)
	baseRouter.POST("/compose/operate", baseApi.OperatorCompose)
	baseRouter.POST("/compose/update", baseApi.ComposeUpdate)
	baseRouter.GET("/compose/search/log", baseApi.ComposeLogs)

	baseRouter.GET("/template", baseApi.ListComposeTemplate)
	baseRouter.POST("/template/search", baseApi.SearchComposeTemplate)
	baseRouter.POST("/template/update", baseApi.UpdateComposeTemplate)
	baseRouter.POST("/template", baseApi.CreateComposeTemplate)
	baseRouter.POST("/template/del", baseRouter.DELETEComposeTemplate)

	baseRouter.GET("/image", baseApi.ListImage)
	baseRouter.POST("/image/search", baseApi.SearchImage)
	baseRouter.POST("/image/pull", baseApi.ImagePull)
	baseRouter.POST("/image/push", baseApi.ImagePush)
	baseRouter.POST("/image/save", baseApi.ImageSave)
	baseRouter.POST("/image/load", baseApi.ImageLoad)
	baseRouter.POST("/image/remove", baseApi.ImageRemove)
	baseRouter.POST("/image/tag", baseApi.ImageTag)
	baseRouter.POST("/image/build", baseApi.ImageBuild)

	baseRouter.GET("/network", baseApi.ListNetwork)
	baseRouter.POST("/network/del", baseRouter.DELETENetwork)
	baseRouter.POST("/network/search", baseApi.SearchNetwork)
	baseRouter.POST("/network", baseApi.CreateNetwork)
	baseRouter.GET("/volume", baseApi.ListVolume)
	baseRouter.POST("/volume/del", baseRouter.DELETEVolume)
	baseRouter.POST("/volume/search", baseApi.SearchVolume)
	baseRouter.POST("/volume", baseApi.CreateVolume)

	baseRouter.GET("/daemonjson", baseApi.LoadDaemonJson)
	baseRouter.GET("/daemonjson/file", baseApi.LoadDaemonJsonFile)
	baseRouter.GET("/docker/status", baseApi.LoadDockerStatus)
	baseRouter.POST("/docker/operate", baseApi.OperateDocker)
	baseRouter.POST("/daemonjson/update", baseApi.UpdateDaemonJson)
	baseRouter.POST("/logoption/update", baseApi.UpdateLogOption)
	baseRouter.POST("/daemonjson/update/byfile", baseApi.UpdateDaemonJsonByFile)*/
}
