// 自动生成模板Namespace
package model

import "time"

// 如果含有time.Time 请自行import time包
type Namespace struct {
	Id          uint       `json:"id" form:"id" gorm:"column:id;comment:id;"`
	AId         uint       `json:"aId" form:"aId" gorm:"column:a_id;comment:应用ID;type:int;size:10;"`
	AppId       string     `json:"appId" form:"appId" gorm:"column:app_id;comment:appId;type:varchar(60);size:60;"`
	Format      string     `json:"format" form:"format" gorm:"column:format;comment:格式;type:varchar(60);size:60;"`
	NodeId      int        `json:"nodeId" form:"nodeId" gorm:"column:node_id;comment:环境ID;type:int;size:10;"`
	Name        string     `json:"name" form:"name" gorm:"column:name;comment:命名空间名称;type:varchar(255);size:255;"`
	Value       string     `json:"value" form:"value" gorm:"column:value;comment:配置信息;type:longtext;"`
	Status      bool       `json:"status" form:"status" gorm:"column:status;comment:发布状态;"`
	CreateTime  *time.Time `json:"createTime" form:"createTime" gorm:"column:create_time;comment:创建时间;default:null;"`
	UpdateTime  *time.Time `json:"updateTime" form:"updateTime" gorm:"column:update_time;comment:更新时间;default:null;"`
	ReleaseTime *time.Time `json:"releaseTime" form:"releaseTime" gorm:"column:release_time;comment:发布时间;default:null;"`
}

func (Namespace) TableName() string {
	return "namespace"
}

// 如果使用工作流功能 需要打开下方注释 并到initialize的workflow中进行注册 且必须指定TableName
// type NamespaceWorkflow struct {
// 	// 工作流操作结构体
// 	WorkflowBase      `json:"wf"`
// 	Namespace   `json:"business"`
// }

// func (Namespace) TableName() string {
// 	return "namespace"
// }

// 工作流注册代码

// initWorkflowModel内部注册
// model.WorkflowBusinessStruct["namespace"] = func() model.GVA_Workflow {
//   return new(model.NamespaceWorkflow)
// }

// initWorkflowTable内部注册
// model.WorkflowBusinessTable["namespace"] = func() interface{} {
// 	return new(model.Namespace)
// }
