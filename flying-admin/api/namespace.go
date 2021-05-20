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
	"strings"
	"sync"
)

var (
	s sync.RWMutex
)

// @Tags Namespace
// @Summary 创建Namespace
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Namespace true "创建Namespace"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /namespace/createNamespace [post]
func CreateNamespace(c *gin.Context) {
	var namespace request.AddNamespace
	_ = c.ShouldBindJSON(&namespace)
	namespace.Name = strings.Trim(namespace.Name, " ")
	nodes, err := logic.GetAppNodes(namespace.Nodes)

	err, app := logic.GetApp(namespace.AppId)
	if err != nil {
		response.BuildResult(constant.NOT_FIND_APP, c)
	}
	namespace.Namespace.AppId = app.AppId
	code, msg, err := logic.RepeatName(&namespace.Namespace, nodes)
	if code <= 0 {
		code, msg := logic.CreateNamespace(&namespace.Namespace, nodes)
		response.BuildResultMsg(code, msg, c)
	} else {
		global.LOG.Error(msg, zap.Any("err", err))
		response.BuildResultMsg(code, msg, c)
	}

}

// @Tags Namespace
// @Summary 删除Namespace
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Namespace true "删除Namespace"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /namespace/deleteNamespace [delete]
func DeleteNamespace(c *gin.Context) {
	var ind request.IdNodeId
	_ = c.ShouldBind(&ind)
	node, err := logic.GetNode(ind.NodeId)
	if err == nil {
		if namespace, err := logic.GetNamespace(ind.Id, node); err == nil {
			code, _ := logic.DeleteNamespace(namespace, node)
			response.BuildResult(code, c)
		} else {
			global.LOG.Error("未查询到该namespace", zap.Any("err", err))
			response.BuildResult(constant.NAMESPACE_DELETE_NOT_FIND, c)
		}
	} else {
		global.LOG.Error("未找到到该namespace的环境", zap.Any("err", err))
		response.BuildResult(constant.NAMESPACE_NODE_DELETE_NOT_FOUND, c)
	}

}

// @Tags Namespace
// @Summary 批量删除Namespace
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Namespace"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /namespace/deleteNamespaceByIds [delete]
func DeleteNamespaceByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := logic.DeleteNamespaceByIds(IDS); err != nil {
		global.LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// @Tags Namespace
// @Summary 更新Namespace
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Namespace true "更新Namespace"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /namespace/updateNamespace [put]
func UpdateNamespace(c *gin.Context) {
	var namespace model.Namespace
	_ = c.ShouldBindJSON(&namespace)
	node, err := logic.GetNode(namespace.NodeId)
	if err == nil {
		namespaces, err := logic.GetNamespace(int64(namespace.Id), node)
		if err == nil {
			if namespace.Value == namespaces.Value {
				response.BuildResult(constant.NAMESPACE_NO_CHANGE, c)
				return
			}
			s.Lock()
			namespaces.Value = namespace.Value
			namespaces.Status = false

			code, _ := logic.UpdateNamespace(namespaces, node)
			s.Unlock()
			response.BuildResult(code, c)
			return
		} else {
			global.LOG.Error("未查询到该namespace", zap.Any("err", err))
			response.BuildResult(constant.NAMESPACE_UPDATE_NOT_FIND, c)
		}
	} else {
		global.LOG.Error("未找到当前namespace的环境", zap.Any("err", err))
		response.BuildResult(constant.NAMESPACE_NODE_UPDATE_NOT_FOUND, c)
	}
	//} else {
	//	global.LOG.Error("更新失败,未查询到当前namespace", zap.Any("err", err))
	//	response.FailWithMessage("更新失败,未查询到当前namespace", c)
	//}

}

// @Tags Namespace
// @Summary 用id查询Namespace
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Namespace true "用id查询Namespace"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /namespace/findNamespace [get]
func FindNamespace(c *gin.Context) {
	var ind request.IdNodeId
	_ = c.ShouldBind(&ind)
	node, err := logic.GetNode(ind.NodeId)
	if err != nil {
		global.LOG.Error("查询环境错误", zap.Any("err", err))
		response.FailWithMessage("查询环境错误", c)
		return
	} else {
		if namespace, err := logic.GetNamespace(ind.Id, node); err != nil {
			global.LOG.Error("查询失败!", zap.Any("err", err))
			response.FailWithMessage("查询失败", c)
			return
		} else {
			response.OkWithData(gin.H{"node": node, "namespace": namespace}, c)
		}

	}
}

// @Tags Namespace
// @Summary 分页获取Namespace列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.NamespaceSearch true "分页获取Namespace列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /namespace/getNamespaceList [get]
func GetNamespaceList(c *gin.Context) {
	var pageInfo request.NamespaceSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := logic.GetNamespaceInfoList(pageInfo); err != nil {
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
func GetAppNamespace(c *gin.Context) {
	var namespace model.Namespace
	_ = c.ShouldBindQuery(&namespace)
	node, err := logic.GetNode(namespace.NodeId)
	if err == nil {
		if err, namespaces := logic.GetAppNamespace(node, &namespace); err != nil {
			global.LOG.Error("获取失败", zap.Any("err", err))
			response.FailWithMessage("获取失败", c)
		} else {
			response.OkWithData(gin.H{"node": node, "list": namespaces}, c)
		}
	}
}
func FindRemoteNamespace(c *gin.Context) {
	var ids model.SyncId
	_ = c.ShouldBindJSON(&ids)
	nodes, err := logic.GetAppNodes(ids.Nodes)
	node, err := logic.GetNode(ids.NId)
	if err == nil {
		if namespace, err := logic.GetNamespace(int64(ids.Id), node); namespace == (&model.Namespace{}) {
			global.LOG.Error("要被同步的namespace没有找到", zap.Any("err", err))
			response.BuildResult(constant.NAMESPACE_RELEASE_NOT_FIND, c)
		} else {
			msg, _ := logic.FindRemoteNamespace(namespace, nodes)
			if len(msg) > 0 {
				response.BuildResultMsg(constant.RELEASE_NAMESPACE_NODE_SYNC_HISTORY, msg, c)
			} else {
				response.Ok(c)
			}
		}
	}
}
func SyncNamespace(c *gin.Context) {
	var ids model.SyncId
	_ = c.ShouldBindJSON(&ids)
	nodes, err := logic.GetAppNodes(ids.Nodes)
	node, err := logic.GetNode(ids.NId)
	if err == nil {
		if namespace, err := logic.GetNamespace(int64(ids.Id), node); namespace == (&model.Namespace{}) {
			global.LOG.Error("要被同步的namespace没有找到", zap.Any("err", err))
			response.BuildResult(constant.RELEASE_NAMESPACE_SYNC_NOT_FOUND, c)
		} else {
			code, _ := logic.SyncNamespace(namespace, nodes)
			response.BuildResult(code, c)
		}
	}
}
