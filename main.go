package main

import (
	"main/dao"
	"main/mqttclient"
	"main/router"
)

func main(){
	dao.ConnectMysql()//连接数据库
	mqttclient.Subcribe()//开启订阅

	r := router.SetupRouter()
	r.Run("127.0.0.1:8083")
}
