package router

import (
	"flying-admin/api"
	"github.com/gin-gonic/gin"
)

func InitAppRouter(Router *gin.RouterGroup) {
	AppRouter := Router.Group("app")
	{
		AppRouter.POST("createApp", api.CreateApp)             // 新建App
		AppRouter.DELETE("deleteApp", api.DeleteApp)           // 删除App
		AppRouter.DELETE("deleteAppByIds", api.DeleteAppByIds) // 批量删除App
		AppRouter.PUT("updateApp", api.UpdateApp)              // 更新App
		AppRouter.GET("findApp/:id", api.FindApp)              // 根据ID获取App
		AppRouter.GET("getAppList", api.GetAppList)            // 获取App列表

	}
}
