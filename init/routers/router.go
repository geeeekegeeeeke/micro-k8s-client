package routers

import (
	ginI18n "github.com/gin-contrib/i18n"
	"github.com/gin-gonic/gin"
	"html/template"
	rou "micro-k8s-client/router"
)

func Routers() *gin.Engine {
	Router := gin.Default()
	//Router.Use(i18n.GinI18nLocalize())
	Router.SetFuncMap(template.FuncMap{
		"Localize": ginI18n.GetMessage,
	})
	systemRouter := rou.RouterGroupApp
	PrivateGroup := Router.Group("/v1/resource")
	{
		systemRouter.InitK8sRouter(PrivateGroup)
		systemRouter.InitDockerRouter(PrivateGroup)
	}

	return Router
}
