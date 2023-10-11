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
	apiV1Group := Router.Group("/v1")
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
		//apiV1Group.GET("/pod/deploy", podctl.ListDeploy)
		apiV1Group.GET("/node", nodectl.ListNode)
		apiV1Group.GET("/node/info", nodectl.GetNodeInfo)
		apiV1Group.GET("/svc/account", svcctl.ListServiceAccount)
		apiV1Group.GET("/svc", svcctl.ListService)
		apiV1Group.GET("/apply/config", applyctl.ConfigMap)
		apiV1Group.GET("/apply/namespace", applyctl.Namespace)
		apiV1Group.PUT("/apply/secret", applyctl.SetSecret)
		apiV1Group.GET("/apply/secret", applyctl.Secret)
		apiV1Group.GET("/apply", applyctl.GetApply)
		apiV1Group.GET("/ingress/info", ingressctl.ListIngress)
		apiV1Group.GET("/ingress", ingressctl.ListIngressInfo)
		apiV1Group.GET("/storage", storagectl.ListPersistent)
		apiV1Group.GET("/storage/vol", storagectl.ListPersistentVol)
		//apiV1Group.GET("/storage", ingressctl.ListIngressInfo)

	}
}
