package common

import (
	"github.com/samuel/go-zookeeper/zk"
	"github.com/spf13/viper"
	"time"
	//"fmt"
)

func SetUpZookeeper() *zk.Conn {
	zooKeeperHosts := viper.GetString("zookeeper.zooKeeperHosts")
	if ZookeeperService == nil {
		zkService, _, err := zk.Connect([]string{zooKeeperHosts}, time.Second)
		if err != nil {
			panic(err)
		}
		ZookeeperService = zkService
	}
	getAllNodes(ZookeeperService)
	return ZookeeperService
}

func ConfigName() string {
	if rootPath == "" {
		rootPath = "/config/" + SYSTEMNAME
	}
	return rootPath
}

func getAllNodes(zk *zk.Conn) {
	rootPath := ConfigName()
	go func() {
		children, _, _, err := zk.ChildrenW(rootPath)
		if err != nil {
			panic(err)
		}
		ChildrenW <- children
		for _, key := range children {
			val, _, err := ZookeeperService.Get(rootPath + "/" + key)
			if err != nil {
				panic(err)
			}
			nodeList[key] = string(val)
		}
	}()
}

func GetString(name string) string {
	if name == "" {
		return ""
	}
	value, ok := nodeList[name]
	if ok {
		return value
	}
	return ""
}

var (
	ZookeeperService *zk.Conn
	rootPath string
	ChildrenW = make(chan []string)
	nodeList = make(map[string]string)
)

const (
	SYSTEMNAME = "manager"
)
