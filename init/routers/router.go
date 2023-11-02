package routers

import (
	rou "gin-dubbogo-consumer/router"
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	systemRouter := rou.RouterGroupApp
	PrivateGroup := Router.Group("/v1/resource")
	{
		systemRouter.InitK8sRouter(PrivateGroup)
		systemRouter.InitDockerRouter(PrivateGroup)
	}

	return Router
}
