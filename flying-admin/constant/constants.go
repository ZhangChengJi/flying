package constant

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
	//AppId已经存在AppId:%s重复
	APPID_DUPLICATE = 10010
	//未找到要更新的app
	NOT_FIND_APP = 10011
	//环境已经存在此项目
	APP_EXISTS_NODE = 10012

	CREATE_SUCCESS     = 20010
	UPDATE_SUCCESS     = 20011
	DELETE_SUCCESS     = 20012
	CONNECTION_SUCCESS = 20013
	RELEASE_SUCCESS    = 20014
	SYNC_SUCCESS       = 20015
	CREATE_FAILED      = 50010
	UPDATE_FAILED      = 50011
	DELETE_FAILED      = 50012
	CONNECTION_FAILED  = 50013
	RELEASE_FAILED     = 50014
	SYNC_FAILED        = 50015

	//环境无法连接
	NODE_FAILED_CONN = 10020
	//环境名称重复
	NODE_NAME_DUPLICATE = 10021
	//环境地址名称重复
	NODE_ADDRESS_DUPLICATE = 10022
	//client节点存在以往app,请清理数据库再使用
	HISTORYLEGACY = 10023
	//未找到要更新的环境
	NODE_NOT_FOUND = 10024
	//环境名称重复
	NODE_ADDRESS_DUPLICATE_UPDATE = 10025
	//不可以删除有%v应用在使用此环境
	NODE_NOT_DELETED = 10026

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

	CURRENT_PASSWORD_FAILED       = 10050
	UPDATE_PASSWORD_FAILED        = 10051
	NEW_CURRENT_PASSWORD_NOT_SAME = 10052

	INVALID_USERNAME_OR_PASSWORD = 40100
	INVALID_ACCESS               = 40101
	TOKEN_EXPIRED                = 40102
	LOGIN_STATUS_FAILED          = 40103
	GET_TOKEN_FAILED             = 40104
	JWT_INVALIDATION_FAILED      = 40105
)
