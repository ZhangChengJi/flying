package middleware

import (
	"crypto/rsa"
	"errors"
	"flying-admin/constant"
	"flying-admin/global"
	"flying-admin/logic"
	"flying-admin/model"
	request "flying-admin/model/request"
	"flying-admin/model/response"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io/ioutil"
	"log"
	"strconv"
	"time"
)

//回调函数，该函数应根据登录信息对用户执行身份验证。
//必须返回用户数据作为用户标识符，它将存储在声明数组中。必修的。
//检查错误（e）以确定适当的错误消息。

func Authenticator() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 我们这里jwt鉴权取头部信息 x-token 登录时回返回token信息 这里前端需要把token存储到cookie或者本地localStorage中 不过需要跟后端协商过期时间 可以约定刷新令牌或者重新登录
		token := c.Request.Header.Get("x-token")
		if token == "" {
			response.FailWithAuth(constant.INVALID_ACCESS, c)
			c.Abort()
			return
		}
		//if service.IsBlacklist(token) {
		//	response.FailWithDetailed(gin.H{"reload": true}, "您的帐户异地登陆或令牌失效", c)
		//	c.Abort()
		//	return
		//}
		j := NewJWT()
		// parseToken 解析token包含的信息
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == TokenExpired {
				response.FailWithAuth(constant.TOKEN_EXPIRED, c)
				c.Abort()
				return
			}

			//response.FailWithAuth(err.Error(), c)
			c.Abort()
			return
		}
		if err, _ = logic.FindUserById(claims.ID); err != nil {
			_ = logic.JsonInBlacklist(model.JwtBlacklist{Jwt: token})
			//response.FailWithAuth(err.Error(), c)
			c.Abort()
		}
		if claims.ExpiresAt-time.Now().Unix() < claims.BufferTime {
			claims.ExpiresAt = time.Now().Unix() + global.GVA_CONFIG.JWT.ExpiresTime
			newToken, _ := j.CreateToken(*claims)
			newClaims, _ := j.ParseToken(newToken)
			c.Header("new-token", newToken)
			c.Header("new-expires-at", strconv.FormatInt(newClaims.ExpiresAt, 10))
			if global.GVA_CONFIG.System.UseMultipoint {
				err, RedisJwtToken := logic.GetRedisJWT(newClaims.Username)
				if err != nil {
					global.LOG.Error("get redis jwt failed", zap.Any("err", err))
				} else { // 当之前的取成功时才进行拉黑操作
					_ = logic.JsonInBlacklist(model.JwtBlacklist{Jwt: RedisJwtToken})
				}
				// 无论如何都要记录当前的活跃状态
				_ = logic.SetRedisJWT(newToken, newClaims.Username)
			}
		}
		c.Set("claims", claims)
		c.Next()
	}
}

type JWT struct {
	SigningKey []byte
}

var (
	TokenExpired     = errors.New("Token is expired")
	TokenNotValidYet = errors.New("Token not active yet")
	TokenMalformed   = errors.New("That's not even a token")
	TokenInvalid     = errors.New("Couldn't handle this token:")
	publicKey        *rsa.PublicKey
	privateKey       *rsa.PrivateKey
)

func NewJWT() *JWT {
	return &JWT{
		[]byte(global.GVA_CONFIG.JWT.SigningKey),
	}
}
func init() {
	publicKeyByte, err := ioutil.ReadFile("public.key")
	if err != nil {
		log.Println(err.Error())
	}
	publicKey, err = jwt.ParseRSAPublicKeyFromPEM(publicKeyByte)
	privateKeyByte, err := ioutil.ReadFile("private.key")
	if err != nil {
		log.Println(err.Error())
	}
	privateKey, _ = jwt.ParseRSAPrivateKeyFromPEM(privateKeyByte)
}

// 创建一个token
func (j *JWT) CreateToken(claims request.CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	return token.SignedString(privateKey)
}

// 解析 token
func (j *JWT) ParseToken(tokenString string) (*request.CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &request.CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return publicKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*request.CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid

	} else {
		return nil, TokenInvalid

	}

}
