package router

import (
	"gin-dubbogo-consumer/controllers"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func init() {
	Router = gin.New()

	//Dobubo-Demo-API接口
	var uctl controllers.UserController
	var podctl controllers.ClientController
	var nodectl controllers.NodeController
	var svcctl controllers.ServiceController
	var applyctl controllers.ApplyController
	var ingressctl controllers.IngressController
	var storagectl controllers.StorageController
	var dockerctl controllers.DcokerController
	var dockerDeployctl controllers.DcokerDeployController
	apiV1Group := Router.Group("/v1/resource")
	{
		apiV1Group.GET("/user/*id", uctl.User)
		apiV1Group.GET("/users", uctl.Users)
		apiV1Group.POST("/user", uctl.Store)
		apiV1Group.PUT("/user", uctl.Update)
		apiV1Group.DELETE("/user", uctl.Destroy)

		//apiV1Group.GET("/pod", podctl.OperateDeploy)
		apiV1Group.GET("/pod/deploy", podctl.ListDeploy)
		apiV1Group.GET("/pod", podctl.GetPodInfo)
		apiV1Group.POST("/pod/deploy", podctl.CreateDeploy)
		apiV1Group.PUT("/pod/deploy", podctl.UpdateDeploy)
		apiV1Group.GET("/pod/image", podctl.GetImage)
		apiV1Group.GET("/pod/container", podctl.GetContainer)
		apiV1Group.GET("/pod/component", podctl.GetComponent)
		//apiV1Group.GET("/pod/deploy", podctl.ListDeploy)
		apiV1Group.GET("/node", nodectl.ListNode)
		apiV1Group.GET("/node/info", nodectl.GetNodeInfo)
		apiV1Group.GET("/svc/account", svcctl.ListServiceAccount)
		apiV1Group.GET("/svc", svcctl.ListService)
		apiV1Group.GET("/apply/config", applyctl.ConfigMap)
		apiV1Group.GET("/apply/namespace", applyctl.Namespace)
		apiV1Group.PUT("/apply/secret", applyctl.SetSecret)
		apiV1Group.GET("/apply/secret", applyctl.Secret)
		apiV1Group.GET("/apply/tomcat", applyctl.DeployTomcatApp)
		//apiV1Group.GET("/apply", applyctl.GetApply)
		apiV1Group.GET("/ingress/info", ingressctl.ListIngress)
		apiV1Group.GET("/ingress", ingressctl.ListIngressInfo)
		apiV1Group.GET("/storage", storagectl.ListPersistent)
		apiV1Group.GET("/storage/vol", storagectl.ListPersistentVol)
		apiV1Group.GET("/dockerdeploy", dockerDeployctl.DeployAppBydirectCompose)
		apiV1Group.GET("/composedeploy", dockerDeployctl.DeployAppComposeParam)
		//
		//apiV1Group.GET("/docker/list", dockerctl.ListContainer)
		//apiV1Group.GET("/docker/search", dockerctl.SearchContainer)
		//apiV1Group.GET("/docker/stats", dockerctl.ContainerStats)
		//apiV1Group.GET("/docker/log", dockerctl.ContainerLogs)
		//apiV1Group.GET("/docker/rename", dockerctl.ContainerOperation)
		//apiV1Group.GET("/docker/exec", dockerctl.ContainerInfo)
		//apiV1Group.GET("/docker/", dockerctl.ContainerCreate)
		//apiV1Group.GET("/storage", ingressctl.ListIngressInfo)
		//apiV1Group.GET("/exec", dockerctl.ContainerWsSsh)  todo
		apiV1Group.GET("/docker/stats/:id", dockerctl.ContainerStats)

		apiV1Group.POST("/docker", dockerctl.ContainerCreate)
		apiV1Group.POST("/docker/update", dockerctl.ContainerUpdate)
		apiV1Group.POST("/docker/upgrade", dockerctl.ContainerUpgrade)
		apiV1Group.POST("/docker/info", dockerctl.ContainerInfo)
		apiV1Group.POST("/docker/search", dockerctl.SearchContainer)
		apiV1Group.POST("/docker/list", dockerctl.ListContainer)
		apiV1Group.GET("/docker/list/stats", dockerctl.ContainerStats)
		apiV1Group.GET("/docker/search/log", dockerctl.ContainerLogs)
		apiV1Group.GET("/docker/limit", dockerctl.LoadResouceLimit)
		apiV1Group.POST("/docker/clean/log", dockerctl.CleanContainerLog)
		apiV1Group.POST("/docker/load/log", dockerctl.LoadContainerLog)
		apiV1Group.POST("/docker/inspect", dockerctl.Inspect)
		apiV1Group.POST("/docker/operate", dockerctl.ContainerOperation)
		apiV1Group.POST("/docker/prune", dockerctl.ContainerPrune)

		//

		apiV1Group.GET("/image/pull", dockerctl.ImagePull)
		apiV1Group.GET("/image/push", dockerctl.ImagePush)
		apiV1Group.GET("/image", dockerctl.ListImage)
		apiV1Group.POST("/image/search", dockerctl.SearchImage)
		apiV1Group.POST("/image/pull", dockerctl.ImagePull)
		apiV1Group.POST("/image/push", dockerctl.ImagePush)
		apiV1Group.POST("/image/save", dockerctl.ImageSave)
		apiV1Group.POST("/image/load", dockerctl.ImageLoad)
		apiV1Group.POST("/image/remove", dockerctl.ImageRemove)
		apiV1Group.POST("/image/tag", dockerctl.ImageTag)
		apiV1Group.POST("/image/build", dockerctl.ImageBuild)

		apiV1Group.GET("/repo", dockerctl.ListRepo)
		apiV1Group.POST("/repo/status", dockerctl.CheckRepoStatus)
		apiV1Group.POST("/repo/search", dockerctl.SearchRepo)
		apiV1Group.POST("/repo/update", dockerctl.UpdateRepo)
		apiV1Group.POST("/repo", dockerctl.CreateRepo)
		apiV1Group.POST("/repo/del", dockerctl.DeleteRepo)

	}
	/*apiV1Group.GET("/exec", dockerctl.ContainerWsSsh)
	apiV1Group.GET("/stats/:id", dockerctl.ContainerStats)

	apiV1Group.POST("", dockerctl.ContainerCreate)
	apiV1Group.POST("/update", dockerctl.ContainerUpdate)
	apiV1Group.POST("/upgrade", dockerctl.ContainerUpgrade)
	apiV1Group.POST("/info", dockerctl.ContainerInfo)
	apiV1Group.POST("/search", dockerctl.SearchContainer)
	apiV1Group.POST("/list", dockerctl.ListContainer)
	apiV1Group.GET("/list/stats", dockerctl.ContainerListStats)
	apiV1Group.GET("/search/log", dockerctl.ContainerLogs)
	apiV1Group.GET("/limit", dockerctl.LoadResourceLimit)
	apiV1Group.POST("/clean/log", dockerctl.CleanContainerLog)
	apiV1Group.POST("/load/log", dockerctl.LoadContainerLog)
	apiV1Group.POST("/inspect", dockerctl.Inspect)
	apiV1Group.POST("/operate", dockerctl.ContainerOperation)
	apiV1Group.POST("/prune", dockerctl.ContainerPrune)

	apiV1Group.GET("/repo", dockerctl.ListRepo)
	apiV1Group.POST("/repo/status", dockerctl.CheckRepoStatus)
	apiV1Group.POST("/repo/search", dockerctl.SearchRepo)
	apiV1Group.POST("/repo/update", dockerctl.UpdateRepo)
	apiV1Group.POST("/repo", dockerctl.CreateRepo)
	apiV1Group.POST("/repo/del", dockerctl.DeleteRepo)

	apiV1Group.POST("/compose/search", dockerctl.SearchCompose)
	apiV1Group.POST("/compose", dockerctl.CreateCompose)
	apiV1Group.POST("/compose/test", dockerctl.TestCompose)
	apiV1Group.POST("/compose/operate", dockerctl.OperatorCompose)
	apiV1Group.POST("/compose/update", dockerctl.ComposeUpdate)
	apiV1Group.GET("/compose/search/log", dockerctl.ComposeLogs)

	apiV1Group.GET("/template", dockerctl.ListComposeTemplate)
	apiV1Group.POST("/template/search", dockerctl.SearchComposeTemplate)
	apiV1Group.POST("/template/update", dockerctl.UpdateComposeTemplate)
	apiV1Group.POST("/template", dockerctl.CreateComposeTemplate)
	apiV1Group.POST("/template/del", dockerctl.DeleteComposeTemplate)

	apiV1Group.GET("/image", dockerctl.ListImage)
	apiV1Group.POST("/image/search", dockerctl.SearchImage)
	apiV1Group.POST("/image/pull", dockerctl.ImagePull)
	apiV1Group.POST("/image/push", dockerctl.ImagePush)
	apiV1Group.POST("/image/save", dockerctl.ImageSave)
	apiV1Group.POST("/image/load", dockerctl.ImageLoad)
	apiV1Group.POST("/image/remove", dockerctl.ImageRemove)
	apiV1Group.POST("/image/tag", dockerctl.ImageTag)
	apiV1Group.POST("/image/build", dockerctl.ImageBuild)

	apiV1Group.GET("/network", dockerctl.ListNetwork)
	apiV1Group.POST("/network/del", dockerctl.DeleteNetwork)
	apiV1Group.POST("/network/search", dockerctl.SearchNetwork)
	apiV1Group.POST("/network", dockerctl.CreateNetwork)
	apiV1Group.GET("/volume", dockerctl.ListVolume)
	apiV1Group.POST("/volume/del", dockerctl.DeleteVolume)
	apiV1Group.POST("/volume/search", dockerctl.SearchVolume)
	apiV1Group.POST("/volume", dockerctl.CreateVolume)

	apiV1Group.GET("/daemonjson", dockerctl.LoadDaemonJson)
	apiV1Group.GET("/daemonjson/file", dockerctl.LoadDaemonJsonFile)
	apiV1Group.GET("/docker/status", dockerctl.LoadDockerStatus)
	apiV1Group.POST("/docker/operate", dockerctl.OperateDocker)
	apiV1Group.POST("/daemonjson/update", dockerctl.UpdateDaemonJson)
	apiV1Group.POST("/logoption/update", dockerctl.UpdateLogOption)
	apiV1Group.POST("/daemonjson/update/byfile", dockerctl.UpdateDaemonJsonByFile)*/
}
