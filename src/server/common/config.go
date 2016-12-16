package common

import (
	"flag"
	"github.com/spf13/viper"
	"log"
	"github.com/kataras/iris"
	"github.com/kataras/go-template/html"
)

func SetUpConfig() {
	var (
		config = flag.String("config", "", "config file path")
	)
	flag.Parse()
	if *config == "" {
		*config = "./src/resources/default"
	}
	viper.SetConfigType("yaml")
	viper.SetConfigName(*config)
	viper.AddConfigPath("./")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Read config fail:", err.Error())
	}
}

func Static()  {
	iris.Static("/js", "./src/static/js", 1)
	iris.Static("/css", "./src/static/css", 1)
}

// 渲染模板配置
func Template() {
	iris.UseTemplate(html.New()).Directory("./src/views", ".html")
}
