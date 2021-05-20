package initialize

import (
	"flying-admin/api"
	"flying-admin/global"
	"flying-admin/middleware"
	"flying-admin/router"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

// 初始化总路由
type Login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

var (
	key                  = []byte("secret key")
	defaultAuthenticator = func(c *gin.Context) (interface{}, error) {
		var loginVals Login
		userID := loginVals.Username
		password := loginVals.Password

		if userID == "admin" && password == "admin" {
			return userID, nil
		}

		return userID, jwt.ErrFailedAuthentication
	}
)

func Routers() *gin.Engine {

	gin.ForceConsoleColor()
	var Router = gin.Default()

	Router.Static("/css", "dist/css")
	Router.Static("/img", "dist/img")
	Router.Static("/js", "dist/js")
	Router.Static("/favicon.ico", "dist/favicon.ico")
	Router.StaticFile("/favicon.ico", "dist/favicon.ico")
	Router.LoadHTMLGlob("dist/index.html")
	Router.GET("/", api.IndexHandler)

	// Router.Use(middleware.LoadTls())  // 打开就能玩https了
	global.LOG.Info("use middleware logger")
	// 跨域
	Router.Use(middleware.Cors())
	global.LOG.Info("use middleware cors")
	// 方便统一添加路由组前缀 多服务器上线使用
	PublicGroup := Router.Group("base")
	{
		PublicGroup.POST("login", api.Login)
	}
	PrivateGroup := Router.Group("")
	PrivateGroup.Use(middleware.Authenticator())
	{
		PrivateGroup.POST("/base/upPass", api.UpPass)
		router.InitNodeRouter(PrivateGroup)      // 节点路由
		router.InitAppRouter(PrivateGroup)       // 应用信息路由
		router.InitNamespaceRouter(PrivateGroup) // 命名空间路由
		router.InitReleaseRouter(PrivateGroup)   // 发布路由
	}
	global.LOG.Info("router register success")
	return Router
}
