package constant

import (
	"errors"
)

type PolicyType int32

const (
	yaml       = 1
	properties = 2
	json       = 3

	QueryAppOK    = "查询成功"
	QueryAppError = "查询成功"

	QueryAppIdError     = "未找到该appId的项目"
	QueryNamespaceError = "未找到该namespace"

	CreateOK    = "创建成功"
	CreateError = "创建失败"

	UpdateOK    = "修改成功"
	UpdateError = "修改失败"

	DeleteOK    = "删除成功"
	DeleteError = "删除失败"

	NotFindApp       = "未找到此应用"
	NotFindNamespace = "未找到此namespace"

	ReleaseOK    = "发布成功"
	ReleaseError = "发布失败"
	SyncOK       = "同步成功"
	SyncError    = "同步失败"

	RepeatApp       = "appId重复"
	RepeatNamespace = "此项目已经存在该namespace"
	ReleaseOther    = "当前namespace已经被其他管理员发布"

	HistoryLegacy = "节点存在以往app,请清理再使用"

	GetConfigErrors = "获取配置失败"

	GetConfigNotAppId     = GetConfigErrors + ",未查询到该AppId：%v 的配置"
	GetConfigNotNamespace = GetConfigErrors + ",未查询到该AppId：%v 的Namespace: %v 的配置,请检查是否存在此namespace或发布状态"

	GetConfigOK = "AppId: %v Namespace: %v 的配置获取成功"

	ERROR   PolicyType = 500
	SUCCESS PolicyType = 200
	UPDATES PolicyType = 304

	CREATE_SUCCESS     PolicyType = 20010
	UPDATE_SUCCESS     PolicyType = 20011
	DELETE_SUCCESS     PolicyType = 20012
	CONNECTION_SUCCESS PolicyType = 20013
	RELEASE_SUCCESS    PolicyType = 20014
	SYNC_SUCCESS                  = 20015
	CREATE_FAILED      PolicyType = 50010
	UPDATE_FAILED      PolicyType = 50011
	DELETE_FAILED      PolicyType = 50012
	CONNECTION_FAILED  PolicyType = 50013
	RELEASE_FAILED     PolicyType = 50014
	SYNC_FAILED                   = 50015

	APPID_DUPLICATE PolicyType = 10010
	NOT_FIND_APP    PolicyType = 10011

	//client节点存在以往app,请清理数据库再使用
	HISTORYLEGACY PolicyType = 10021

	NAMESPACE_REPEAT                = 10030
	NAMESPACE_UPDATE_NOT_FIND       = 10031
	NAMESPACE_DELETE_NOT_FIND       = 10032
	NAMESPACE_RELEASE_NOT_FIND      = 10033
	NAMESPACE_NO_CHANGE             = 10034
	NAMESPACE_NODE_UPDATE_NOT_FOUND = 10035
	NAMESPACE_NODE_DELETE_NOT_FOUND = 10036

	RELEASE_OTHER                       = 10040
	RELEASE_NAMESPACE_SYNC_NOT_FOUND    = 10041
	RELEASE_NAMESPACE_NODE_SYNC_HISTORY = 10042
	RELEASE_APP_NO_FIND                 = 10043
)

func (p PolicyType) String() string {
	switch p {
	case ERROR:
		return "failed!"
	case APPID_DUPLICATE:
		return "duplicate appId!"
	case NOT_FIND_APP:
		return "This app was not found!"

	default:
		return "unknown"
	}
}

var (
	DUPLICATE = errors.New(APPID_DUPLICATE.String())
)
