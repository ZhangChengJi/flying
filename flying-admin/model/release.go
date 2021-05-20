// 自动生成模板Release
package model

import (
	"time"
)

// 如果含有time.Time 请自行import time包
type Release struct {
	Id            uint       `json:"id" form:"id" gorm:"column:id;comment:id;"`
	AppId         string     `json:"appId" form:"appId" gorm:"column:app_id;comment:应用ID;type:varchar(500);"`
	Comment       string     `json:"comment" form:"comment" gorm:"column:comment;comment:发布说明;type:varchar(255);size:255;"`
	Name          string     `json:"name" form:"name" gorm:"column:name;comment:发布名称;type:varchar(64);size:64;"`
	NamespaceName string     `json:"namespaceName" form:"namespaceName" gorm:"column:namespace_name;comment:命名空间;type:varchar(64);size:64;"`
	ReleaseKey    string     `json:"releaseKey" form:"releaseKey" gorm:"column:release_key;comment:发布key;type:varchar(64);size:64;"`
	Value         string     `json:"value" form:"value" gorm:"column:value;comment:配置信息;type:longtext;"`
	CreateTime    *time.Time `json:"createTime" form:"createTime" gorm:"column:create_time;comment:创建时间;default:null;"`
	UpdateTime    *time.Time `json:"updateTime" form:"updateTime" gorm:"column:update_time;comment:更新时间;default:null;"`
}

func (Release) TableName() string {
	return "release"
}

// 如果使用工作流功能 需要打开下方注释 并到initialize的workflow中进行注册 且必须指定TableName
// type ReleaseWorkflow struct {
// 	// 工作流操作结构体
// 	WorkflowBase      `json:"wf"`
// 	Release   `json:"business"`
// }

// func (Release) TableName() string {
// 	return "release"
// }

// 工作流注册代码

// initWorkflowModel内部注册
// model.WorkflowBusinessStruct["release"] = func() model.GVA_Workflow {
//   return new(model.ReleaseWorkflow)
// }

// initWorkflowTable内部注册
// model.WorkflowBusinessTable["release"] = func() interface{} {
// 	return new(model.Release)
// }
