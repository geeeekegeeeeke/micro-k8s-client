package main

import (
	"encoding/json"
	"gin-dubbogo-consumer/init/db"
	"gin-dubbogo-consumer/init/migration"
	"gin-dubbogo-consumer/init/routers"
	"github.com/gin-gonic/gin"
	"gitlab.stagingvip.net/publicGroup/public/common"
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
