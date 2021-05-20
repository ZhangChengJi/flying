package main

import (
	"flying-admin/core"
	"flying-admin/global"
	"flying-admin/initialize"
	"flying-admin/model"
	"flying-admin/utils"
)

func main() {
	global.GVA_VP = core.Viper()      // 初始化Viper
	global.LOG = core.Zap()           // 初始化zap日志库
	global.DB = initialize.Gorm()     // gorm连接数据库
	initialize.MysqlTables(global.DB) // 初始化表
	// 程序结束前关闭数据库链接
	db, _ := global.DB.DB()
	defer db.Close()
	model.NewGrpcService()
	initialize.NewEvent()
	utils.ConnectTestTask()
	core.RunWindowsServer()

}
