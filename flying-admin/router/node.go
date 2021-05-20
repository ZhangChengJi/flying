package router

import (
	"flying-admin/api"
	"github.com/gin-gonic/gin"
)

func InitNodeRouter(Router *gin.RouterGroup) {
	NodeRouter := Router.Group("node")
	{
		NodeRouter.POST("createNode", api.CreateNode)             // 新建Node
		NodeRouter.DELETE("deleteNode", api.DeleteNode)           // 删除Node
		NodeRouter.DELETE("deleteNodeByIds", api.DeleteNodeByIds) // 批量删除Node
		NodeRouter.PUT("updateNode", api.UpdateNode)              // 更新Node
		NodeRouter.GET("findNode/:id", api.FindNode)              // 根据ID获取Node
		NodeRouter.GET("getNodeList", api.GetNodeList)            // 获取Node列表
		NodeRouter.GET("getNodeAll", api.GetNodeAll)              // 获取所有的环境
		NodeRouter.GET("ping/:id", api.Ping)                      // 测试是否连接
		NodeRouter.GET("getAppNode/:appId", api.GetAppNode)       // 获取当前app的环境
		NodeRouter.GET("getSyncNode", api.GetSyncNode)            // 获取要进行同步的namespace
	}
}
