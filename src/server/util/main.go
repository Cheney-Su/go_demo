package server

import (
	"fmt"
	"github.com/kataras/iris"
	"go/src/server/Test"
	//"go/src/server/redis"
	//"log"
	//"go/src/server/common/nsq"
	//"github.com/samuel/go-zookeeper/zk"
)

func SetUpServer() {
	fmt.Println("Iris Start...")

	var user = user.UserController{}
	iris.Get("/getUserInfo", user.GetUserInfoById)
	iris.Post("/createUser", user.CreateUser)

	iris.Get("/test", test)

	iris.Listen(":8080")
}

func test(c *iris.Context) {
	//key := "shq:test"
	//user := make(map[string]string)
	//user["user"] = "shq"
	//user["password"] = "123456"
	//myMap := make(map[string]string)
	//result := redis.Hmset(key, myMap)
	//if result {
	//	log.Println("redis操作成功")
	//}else {
	//	log.Println("redis操作异常")
	//}
	//发布和订阅nsq
	//common.TestPublist()
	//common.TestConsume()
	//模拟http请求，转换json
	//respone, err := grequests.Get("http://cardcenter.maizuo.com/getBatchInfo.htm?batchNo=115782&format=json", nil)
	//if err != nil {
	//	log.Fatalln("Unable to make request: ", err)
	//}
	//
	//fmt.Println(respone.String())
	//json, _ := simplejson.NewJson(respone.Bytes())
	//fmt.Println(json)
	//获取配置项中的值
	//fmt.Println(common.GetString("ACTIVE_IN_DOMAIN"))

}
