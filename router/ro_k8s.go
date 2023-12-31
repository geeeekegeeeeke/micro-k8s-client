package router

import (
	v1 "gin-dubbogo-consumer/api/v1"
	"github.com/gin-gonic/gin"
)

type K8sRouter struct {
}

func (a *K8sRouter) InitK8sRouter(Router *gin.RouterGroup) {
	appRouter := Router.Group("k8s")
	//appRouter.Use(middleware.JwtAuth()).Use(middleware.SessionAuth()).Use(middleware.PasswordExpired())

	baseApi := v1.ApiGroupApp.BaseApi
	{
		appRouter.POST("/create", baseApi.Create)
		appRouter.GET("/update", baseApi.Update)
		appRouter.POST("/list", baseApi.List)
	}
}
