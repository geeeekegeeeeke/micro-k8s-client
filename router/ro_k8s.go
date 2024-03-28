package router

import (
	"github.com/gin-gonic/gin"
	v1 "micro-k8s-client/api/v1"
)

type K8sRouter struct {
}

func (a *K8sRouter) InitK8sRouter(Router *gin.RouterGroup) {
	appRouter := Router.Group("k8scluster")
	//appRouter.Use(middleware.JwtAuth()).Use(middleware.SessionAuth()).Use(middleware.PasswordExpired())

	baseApi := v1.ApiGroupApp.BaseApi
	{
		appRouter.POST("/create", baseApi.Create)
		appRouter.POST("/update", baseApi.Update)
		appRouter.DELETE("/delete/:id", baseApi.Delete)
		appRouter.GET("/info/:id", baseApi.GetInfo)
		appRouter.GET("/list", baseApi.List)
	}
}
