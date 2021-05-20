package api

import (
	"flying-admin/constant"
	"flying-admin/global"
	"flying-admin/logic"
	"flying-admin/model"
	"flying-admin/model/request"
	"flying-admin/model/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"strings"
)

func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

// @Tags App
// @Summary 创建App
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.App true "创建App"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /app/createApp [post]

func CreateApp(c *gin.Context) {
	var app request.AppNodeList
	err := c.ShouldBind(&app)
	app.AppId = strings.Trim(app.AppId, " ")
	code, msg := logic.GetIf(&app)
	if len(msg) <= 0 {
		nodes, _ := logic.GetAppNodes(app.Nodes)
		code, _ := logic.CreateApp(&app.App, nodes)
		response.BuildResult(code, c)

	} else {
		global.LOG.Error(msg, zap.Any("err", err))
		response.BuildResultMsg(code, msg, c)
	}
}

// @Tags App
// @Summary 删除App
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.App true "删除App"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /app/deleteApp [delete]
func DeleteApp(c *gin.Context) {
	var app model.App
	_ = c.ShouldBindJSON(&app)
	_, nodes := logic.GetDeleteAppNode(app.Id)
	code, _ := logic.DeleteApp(app, nodes)
	response.BuildResult(code, c)
}

// @Tags App
// @Summary 批量删除App
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除App"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /app/deleteAppByIds [delete]
func DeleteAppByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := logic.DeleteAppByIds(IDS); err != nil {
		global.LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags App
// @Summary 更新App
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.App true "更新App"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /app/updateApp [put]
func UpdateApp(c *gin.Context) {
	var app model.AppDetail
	_ = c.ShouldBindJSON(&app)
	nodes, _ := logic.GetAppAllNodes(app.App.Id)
	code, _ := logic.UpdateApp(app, nodes)
	if code == 0 {
		response.BuildResult(constant.UPDATE_SUCCESS, c)
		return
	}
	response.BuildResult(code, c)

}

// @Tags App
// @Summary 用id查询App
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.App true "用id查询App"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /app/findApp [get]
func FindApp(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err, app := logic.GetAppDetail(uint(id)); err != nil {
		global.LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"app": app}, c)
	}
}

// @Tags App
// @Summary 分页获取App列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.AppSearch true "分页获取App列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /app/getAppList [get]
func GetAppList(c *gin.Context) {
	var pageInfo request.AppSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := logic.GetAppInfoList(pageInfo); err != nil {
		global.LOG.Error("获取失败", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
