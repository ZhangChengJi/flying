package api

import (
	"flying-admin/global"
	"flying-admin/logic"
	"flying-admin/model"
	"flying-admin/model/request"
	"flying-admin/model/response"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// @Tags Release
// @Summary 创建Release
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Release true "创建Release"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /release/createRelease [post]
func CreateRelease(c *gin.Context) {
	var release request.Release
	_ = c.ShouldBindJSON(&release)
	node, _ := logic.GetNode(release.NodeId)
	code, _ := logic.CreateRelease(node, release)
	response.BuildResult(code, c)
	return

}

// @Tags Release
// @Summary 删除Release
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Release true "删除Release"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /release/deleteRelease [delete]
func DeleteRelease(c *gin.Context) {
	var release model.Release
	_ = c.ShouldBindJSON(&release)
	if err := logic.DeleteRelease(release); err != nil {
		global.LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// @Tags Release
// @Summary 批量删除Release
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Release"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /release/deleteReleaseByIds [delete]
func DeleteReleaseByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := logic.DeleteReleaseByIds(IDS); err != nil {
		global.LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags Release
// @Summary 更新Release
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Release true "更新Release"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /release/updateRelease [put]
//func UpdateRelease(c *gin.Context) {
//	var release model.Release
//	_ = c.ShouldBind(&release)
//	now := time.Now()
//	release.CreateTime = &now
//	release.UpdateTime = &now
//	if err := logic.UpdateRelease(release); err != nil {
//		global.LOG.Error("更新失败!", zap.Any("err", err))
//		response.FailWithMessage("更新失败", c)
//	} else {
//		listener.AddListener(release.AppId)
//		response.OkWithMessage("更新成功", c)
//	}
//}

// @Tags Release
// @Summary 用id查询Release
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Release true "用id查询Release"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /release/findRelease [get]
func FindRelease(c *gin.Context) {
	var release model.Release
	_ = c.ShouldBindQuery(&release)
	if err, rerelease := logic.GetRelease(release.Id); err != nil {
		global.LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rerelease": rerelease}, c)
	}
}

// @Tags Release
// @Summary 分页获取Release列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.ReleaseSearch true "分页获取Release列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /release/getReleaseList [get]
//func GetReleaseList(c *gin.Context) {
//	var pageInfo request.ReleaseSearch
//	_ = c.ShouldBindQuery(&pageInfo)
//	if err, list, total := logic.GetReleaseInfoList(pageInfo); err != nil {
//		global.LOG.Error("获取失败", zap.Any("err", err))
//		response.FailWithMessage("获取失败", c)
//	} else {
//		response.OkWithDetailed(response.PageResult{
//			List:     list,
//			Total:    total,
//			Page:     pageInfo.Page,
//			PageSize: pageInfo.PageSize,
//		}, "获取成功", c)
//	}
//}
