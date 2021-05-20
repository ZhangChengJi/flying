package response

type FlyingConfig struct {
	AppId         string `json:"appId" form:"appId" gorm:"column:app_id;comment:应用ID;type:varchar(500);"`
	NamespaceName string `json:"namespaceName" form:"namespaceName" gorm:"column:namespace_name;comment:命名空间;type:varchar(64);size:64;"`
	ReleaseKey    string `json:"releaseKey" form:"releaseKey" gorm:"column:release_key;comment:发布key;type:varchar(64);size:64;"`
	Value         string `json:"value" form:"value" gorm:"column:value;comment:配置信息;type:longtext;"`
}
