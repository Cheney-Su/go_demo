package common

import (
	"github.com/spf13/viper"
	"database/sql"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

func SetUPDB() {
	setDB()
}

func setDB() {
	dialect := viper.GetString("db.dialect")
	host := viper.GetString("db.host")
	port := viper.GetString("db.port")
	user := viper.GetString("db.user")
	password := viper.GetString("db.password")
	database := viper.GetString("db.database")
	maxIdle := viper.GetInt("db.maxIdle")
	maxOpen := viper.GetInt("db.maxOpen")
	var url string
	switch dialect {
	case "mysql":
		url = user + ":" + password + "@tcp(" + host + ":" + port + ")/" + database + "?charset=utf8&parseTime=True&loc=Local"
	}
	log.Print("数据库连接中...")
	db, err := sql.Open(dialect, url)
	if err!=nil {
		log.Fatalln("连接数据库异常，err=", err)
	}else {
		log.Print("连接数据库成功")
	}

	db.SetMaxIdleConns(maxIdle)
	db.SetMaxOpenConns(maxOpen)

	DB = db
}

var (
	DB *sql.DB
)