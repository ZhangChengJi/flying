package api

import (
	"context"
	"flying-admin/constant"
	"flying-admin/global"
	"flying-admin/logic"
	"flying-admin/model"
	"flying-admin/model/request"
	"flying-admin/model/response"
	"flying-admin/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func Ping(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	node, _ := logic.GetNode(uint(id))
	_, err := node.Client().Ping(context.Background(), &emptypb.Empty{})
	if err == nil {
		node.Status = true
		if err := logic.UpdateNode(node); err != nil {
			global.LOG.Error(err.Error(), zap.Any("err", err))
			response.BuildResult(constant.UPDATE_FAILED, c)
		} else {
			response.BuildResult(constant.CONNECTION_SUCCESS, c)
		}
	} else {
		response.BuildResult(constant.CONNECTION_FAILED, c)
	}
}

// @Tags Node
// @Summary 创建Node
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Node true "创建Node"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /node/createNode [post]
func CreateNode(c *gin.Context) {
	var node model.Node
	_ = c.ShouldBindJSON(&node)
	node.Url = strings.Trim(node.Url, " ")
	node.Name = strings.Trim(node.Name, " ")
	i := logic.IfExists(&node)
	if i == 1 {
		response.BuildResult(constant.NODE_NAME_DUPLICATE, c)
		return
	} else if i == 2 {
		response.BuildResult(constant.NODE_ADDRESS_DUPLICATE, c)
		return
	}
	t := &utils.Telnet{
		Node: &node,
	}
	resp, err := t.QueryApp()
	if err != nil {
		response.BuildResult(constant.NODE_FAILED_CONN, c)
		return
	}
	if resp.Code != http.StatusOK {
		response.BuildResult(resp.Code, c)
		return
	} else {
		node.Key = strconv.FormatInt(time.Now().UnixNano(), 10)
		node.Status = true
		if err := logic.CreateNode(&node); err != nil {
			global.LOG.Error(err.Error(), zap.Any("err", err))
			response.BuildResult(constant.CREATE_FAILED, c)
			return
		} else {
			//事件
			utils.Dispatcher.DispatchEvent(utils.Add, node)
			response.BuildResult(constant.CREATE_SUCCESS, c)
			return
		}
	}
}

// @Tags Node
// @Summary 删除Node
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Node true "删除Node"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /node/deleteNode [delete]
func DeleteNode(c *gin.Context) {
	var node model.Node
	_ = c.ShouldBindJSON(&node)
	err, count := logic.GetNodeApp(node.Id)
	if err == nil {
		if count > 0 {
			response.BuildResultMsg(constant.NODE_NOT_DELETED, strconv.Itoa(int(count)), c)
			return
		}
		if err := logic.DeleteNode(node); err != nil {
			global.LOG.Error(err.Error(), zap.Any("err", err))
			response.BuildResult(constant.DELETE_FAILED, c)
		} else {
			//事件
			utils.Dispatcher.DispatchEvent(utils.Del, node)
			response.BuildResult(constant.DELETE_SUCCESS, c)
		}
	} else {
		response.BuildResult(constant.NODE_NOT_FOUND, c)
	}
}

// @Tags Node
// @Summary 批量删除Node
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Node"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /node/deleteNodeByIds [delete]
func DeleteNodeByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := logic.DeleteNodeByIds(IDS); err != nil {
		global.LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags Node
// @Summary 更新Node
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Node true "更新Node"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /node/updateNode [put]
func UpdateNode(c *gin.Context) {
	var node model.Node
	_ = c.ShouldBindJSON(&node)
	_, err := logic.GetNode(node.Id)
	if err == nil {
		if !logic.IfRepeat(&node) {
			if err := logic.UpdateNode(&node); err != nil {
				global.LOG.Error(err.Error(), zap.Any("err", err))
				response.BuildResult(constant.UPDATE_FAILED, c)
			} else {
				//事件
				utils.Dispatcher.DispatchEvent(utils.Update, node)
				response.BuildResult(constant.UPDATE_SUCCESS, c)
			}

		} else {
			response.BuildResult(constant.NODE_ADDRESS_DUPLICATE_UPDATE, c)
		}
	} else {
		global.LOG.Error(err.Error(), zap.Any("err", err))
		response.BuildResult(constant.NODE_NOT_FOUND, c)
	}

}

// @Tags Node
// @Summary 用id查询Node
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Node true "用id查询Node"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /node/findNode [get]
func FindNode(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if node, err := logic.GetNode(uint(id)); err != nil {
		global.LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"node": node}, c)
	}
}

// @Tags Node
// @Summary 分页获取Node列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.NodeSearch true "分页获取Node列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /node/getNodeList [get]
func GetNodeList(c *gin.Context) {
	var pageInfo request.NodeSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := logic.GetNodeInfoList(pageInfo); err != nil {
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
func GetNodeAll(c *gin.Context) {
	if err, reapp := logic.GetAll(); err != nil {
		global.LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"list": reapp}, c)
	}
}
func GetAppNode(c *gin.Context) {
	if err, node, app := logic.GetAppNode(strings.Trim(c.Param("appId"), " ")); err != nil {
		global.LOG.Error("获取失败", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithData(gin.H{"node": node, "app": app}, c)
	}

}

func GetSyncNode(c *gin.Context) {
	var sync request.SyncId
	_ = c.ShouldBindQuery(&sync)
	if err, node := logic.GetSyncNode(&sync); err != nil {
		global.LOG.Error("获取node失败", zap.Any("err", err))
		response.FailWithMessage("获取node失败", c)
	} else {
		response.OkWithData(gin.H{"node": node}, c)
	}
}
