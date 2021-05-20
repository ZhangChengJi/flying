package api

import (
	"flying-admin/constant"
	"flying-admin/global"
	"flying-admin/logic"
	"flying-admin/middleware"
	"flying-admin/model"
	"flying-admin/model/request"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis"
	"time"

	"flying-admin/model/response"
	"flying-admin/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Login(c *gin.Context) {
	var L request.Login
	_ = c.ShouldBind(&L)
	if err := utils.Verify(L, utils.LoginVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	U := &model.User{Username: L.Username, Password: L.Password}
	if err, user := logic.Login(U); err != nil {
		global.LOG.Error("登陆失败! 用户名不存在或者密码错误", zap.Any("err", err))
		response.FailWithAuth(constant.INVALID_USERNAME_OR_PASSWORD, c)
	} else {
		tokenNext(c, *user)
	}

}

// 登录以后签发jwt
func tokenNext(c *gin.Context, user model.User) {
	j := &middleware.JWT{SigningKey: []byte(global.GVA_CONFIG.JWT.SigningKey)} // 唯一签名
	claims := request.CustomClaims{
		ID:         user.Id,
		NickName:   user.NickName,
		Username:   user.Username,
		BufferTime: global.GVA_CONFIG.JWT.BufferTime, // 缓冲时间1天 缓冲时间内会获得新的token刷新令牌 此时一个用户会存在两个有效令牌 但是前端只留一个 另一个会丢失
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,                              // 签名生效时间
			ExpiresAt: time.Now().Unix() + global.GVA_CONFIG.JWT.ExpiresTime, // 过期时间 半个小时  配置文件
			Issuer:    "qmPlus",                                              // 签名的发行者
		},
	}
	token, err := j.CreateToken(claims)
	if err != nil {
		global.LOG.Error("获取token失败", zap.Any("err", err))
		response.FailWithAuth(constant.GET_TOKEN_FAILED, c)
		return
	}
	if !global.GVA_CONFIG.System.UseMultipoint {
		response.OkWithDetailed(model.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功", c)
		return
	}
	if err, jwtStr := logic.GetRedisJWT(user.Username); err == redis.Nil {
		if err := logic.SetRedisJWT(token, user.Username); err != nil {
			global.LOG.Error("设置登录状态失败", zap.Any("err", err))
			response.FailWithAuth(constant.LOGIN_STATUS_FAILED, c)
			return
		}
		response.OkWithDetailed(model.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功", c)
	} else if err != nil {
		global.LOG.Error("设置登录状态失败", zap.Any("err", err))
		response.FailWithAuth(constant.LOGIN_STATUS_FAILED, c)
	} else {
		var blackJWT model.JwtBlacklist
		blackJWT.Jwt = jwtStr
		if err := logic.JsonInBlacklist(blackJWT); err != nil {
			response.FailWithAuth(constant.JWT_INVALIDATION_FAILED, c)
			return
		}
		if err := logic.SetRedisJWT(token, user.Username); err != nil {
			response.FailWithAuth(constant.LOGIN_STATUS_FAILED, c)
			return
		}
		response.OkWithDetailed(model.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功", c)
	}
}

func UpPass(c *gin.Context) {
	var u request.UpPass
	_ = c.ShouldBind(&u)
	if err := utils.Verify(u, utils.LoginVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	user := &model.User{Username: u.Username, Password: u.Password}
	if err, user := logic.Login(user); err != nil {
		global.LOG.Error("当前密码错误", zap.Any("err", err))
		response.BuildResult(constant.CURRENT_PASSWORD_FAILED, c)

	} else if u.Password != u.Pass {
		user.Password = logic.MD5V([]byte(u.Pass))
		err := logic.UpPass(user)
		if err != nil {
			global.LOG.Error("更新密码错误", zap.Any("err", err))
			response.BuildResult(constant.UPDATE_PASSWORD_FAILED, c)
			return
		} else {
			response.BuildResult(constant.UPDATE_SUCCESS, c)
		}
	} else {
		response.BuildResult(constant.NEW_CURRENT_PASSWORD_NOT_SAME, c)
	}
}
