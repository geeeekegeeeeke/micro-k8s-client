package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"gitlab.stagingvip.net/publicGroup/public/common"
	"micro-k8s-client/init/db"
	"micro-k8s-client/init/migration"
	"micro-k8s-client/init/routers"
)

func main() {
	Router := gin.Default()
	db.Init()
	migration.Init()
	confByte, err := common.ReadFile("./conf/conf.json")
	if err != nil {
		panic(err)
	}
	Router = routers.Routers()
	var jsonConf map[string]string
	err = json.Unmarshal(confByte, &jsonConf)
	//Router = router.Router
	if err != nil {
		panic(err)
	}
	_ = Router.Run(jsonConf["port"])
}
