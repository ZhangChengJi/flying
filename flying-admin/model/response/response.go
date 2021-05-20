package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

const (
	ERROR   = 500
	SUCCESS = 200
	UPDATES = 2
	CONFIRM = 3
	CONNECT = 4
)

func Result(code int, data interface{}, msg string, c *gin.Context) {
	// 开始时间
	c.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}
func AuthResult(code int32, data interface{}, msg string, c *gin.Context) {
	// 开始时间
	c.JSON(http.StatusUnauthorized, Response{
		int(code),
		data,
		msg,
	})
}
func BuildResult(code int32, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		int(code),
		map[string]interface{}{},
		"",
	})
}
func BuildResultMsg(code int32, message string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		int(code),
		map[string]interface{}{},
		message,
	})
}
func Ok(c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, "操作成功", c)
}

func OkWithMessage(message string, c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, message, c)
}
func OkWithUpdateMessage(message string, c *gin.Context) {
	Result(UPDATES, map[string]interface{}{}, message, c)
}
func OkWithData(data interface{}, c *gin.Context) {
	Result(SUCCESS, data, "操作成功", c)
}

func OkWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(SUCCESS, data, message, c)
}
func ConfirmWithMessage(message string, c *gin.Context) {
	Result(CONFIRM, map[string]interface{}{}, message, c)
}
func ConnectFailed(message string, c *gin.Context) {
	Result(CONNECT, map[string]interface{}{}, message, c)
}
func BreakFailed(message string, c *gin.Context) {
	Result(CONFIRM, map[string]interface{}{}, message, c)
}
func Fail(c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, "", c)
}

func FailWithMessage(message string, c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, message, c)
}

func FailWithDetailed(data interface{}, message string, c *gin.Context) {
	Result(ERROR, data, message, c)
}

func FailWithAuth(code int32, c *gin.Context) {
	AuthResult(code, map[string]interface{}{}, "", c)
}
