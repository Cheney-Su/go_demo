package main

import (
	"go/src/server/common"
	"go/src/server/service"
)

func main() {
	common.SetUpConfig()
	common.Static()
	common.Template()
	common.SetUPDB()
	common.SetUpRedis()
	common.TestConsume()
	common.SetUpZookeeper()
	//server.SetUpServer()
	service.Websocket()
}
