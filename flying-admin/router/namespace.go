package router

import (
	"flying-admin/api"
	"github.com/gin-gonic/gin"
)

func InitNamespaceRouter(Router *gin.RouterGroup) {
	NamespaceRouter := Router.Group("namespace")
	{
		//NamespaceRouter.POST("release", v1.Release)                             // 发布Namespace
		NamespaceRouter.POST("createNamespace", api.CreateNamespace)             // 新建Namespace
		NamespaceRouter.DELETE("deleteNamespace", api.DeleteNamespace)           // 删除Namespace
		NamespaceRouter.DELETE("deleteNamespaceByIds", api.DeleteNamespaceByIds) // 批量删除Namespace
		NamespaceRouter.PUT("updateNamespace", api.UpdateNamespace)              // 更新Namespace
		NamespaceRouter.GET("findNamespace", api.FindNamespace)                  // 根据ID获取Namespace
		NamespaceRouter.GET("getNamespaceList", api.GetNamespaceList)            // 获取Namespace列表
		NamespaceRouter.GET("getAppNamespace", api.GetAppNamespace)              // 根据环境获取Namespace
		NamespaceRouter.POST("syncNamespace", api.SyncNamespace)                 //同步配置
		NamespaceRouter.POST("findRemoteNamespace", api.FindRemoteNamespace)     //远程查询配置是否存在
	}
}
