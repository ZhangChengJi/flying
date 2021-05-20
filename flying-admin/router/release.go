package router

import (
	"flying-admin/api"
	"github.com/gin-gonic/gin"
)

func InitReleaseRouter(Router *gin.RouterGroup) {
	ReleaseRouter := Router.Group("release")
	{
		ReleaseRouter.POST("createRelease", api.CreateRelease)
	}
}
