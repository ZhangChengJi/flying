// 自动生成模板AppNode
package model

// 如果含有time.Time 请自行import time包
type AppNode struct {
	Id     uint   `json:"id" form:"id" gorm:"column:id;comment:id;"`
	AppId  string `json:"appId" form:"appId" gorm:"column:app_id;comment:AppId;type:varchar(60);size:60;"`
	AId    uint   `json:"aId" form:"aId" gorm:"column:a_id;comment:项目id;type:int;size:10;"`
	NodeId uint   `json:"nodeId" form:"nodeId" gorm:"column:node_id;comment:环境id;type:int;size:10;"`
}

func (AppNode) TableName() string {
	return "app_node"
}

// 如果使用工作流功能 需要打开下方注释 并到initialize的workflow中进行注册 且必须指定TableName
// type AppNodeWorkflow struct {
// 	// 工作流操作结构体
// 	WorkflowBase      `json:"wf"`
// 	AppNode   `json:"business"`
// }

// func (AppNode) TableName() string {
// 	return "app_node"
// }

// 工作流注册代码

// initWorkflowModel内部注册
// model.WorkflowBusinessStruct["appNode"] = func() model.GVA_Workflow {
//   return new(model.AppNodeWorkflow)
// }

// initWorkflowTable内部注册
// model.WorkflowBusinessTable["appNode"] = func() interface{} {
// 	return new(model.AppNode)
// }
