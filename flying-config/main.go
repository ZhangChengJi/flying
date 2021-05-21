package main

import (
	"flying-config/core"
	"flying-config/global"
	"flying-config/initialize"
	"flying-config/listener"
)

func main() {
	listener.Init(10*1024 * 1024)
	global.GVA_VP = core.Viper()      // 初始化Viper
	global.LOG = core.Zap()           // 初始化zap日志库
	global.DB = initialize.Gorm()     // gorm连接数据库
	initialize.MysqlTables(global.DB) // 初始化表
	// 程序结束前关闭数据库链接
	db, _ := global.DB.DB()
	defer db.Close()
	core.RunServer()

}
