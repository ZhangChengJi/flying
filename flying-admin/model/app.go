// 自动生成模板App
package model

import (
	"encoding/json"
	"flying-admin/global"
	"go.uber.org/zap"
	"net/http"
	"strings"
	"time"
)

// 如果含有time.Time 请自行import time包
type App struct {
	Id         uint       `json:"id" form:"id" gorm:"column:id;comment:id;"`
	AppId      string     `json:"appId" form:"appId" gorm:"column:app_id;comment:AppId;type:varchar(60);size:60;"`
	Name       string     `json:"name" form:"name" gorm:"column:name;comment:应用名称;type:varchar(255);size:255;"`
	CreateTime *time.Time `json:"createTime" form:"createTime" gorm:"column:create_time;comment:创建时间;default:null;"`
	UpdateTime *time.Time `json:"updateTime" form:"updateTime" gorm:"column:update_time;comment:更新时间;default:null;"`
}
type AppList struct {
	App
	Node []string `json:"node"`
}
type AppDetail struct {
	App
	Nodes []uint `json:"nodes" form:"nodes"`
}

func (App) TableName() string {
	return "app"
}

type FlyingClient interface {
	CreateRemoteApp(node []*Node)
	DeleteRemoteApp(nodes []*Node) (err error)
	UpdateRemoteApp(node *Node) (err error)
}

//func (app *App) CreateRemoteApp(nodes Node) (err error) {
//
//
//		c,_:=nodes.New()
//		a,err := pb.NewAppServiceClient(c).Create(context.Background(), app)
//		fmt.Println(a)
//
//
//	return err
//
//}
func (app *App) DeleteRemoteApp(nodes []*Node) (err error) {
	contentType := "application/json"
	jsonByte, _ := json.Marshal(app)
	for _, v := range nodes {
		url := v.Url + "/app/deleteApp"
		resp, err := http.NewRequest(http.MethodDelete, url, strings.NewReader(string(jsonByte)))
		resp.Header.Add("Content-Type", contentType)
		_, err = http.DefaultClient.Do(resp)
		if err != nil {
			global.LOG.Error("远程删除失败!", zap.Any("err", err))
			return err
		}
		resp.Body.Close()
	}
	return err

}
func (app *App) UpdateRemoteApp(nodes []*Node) (err error) {
	contentType := "application/json"
	jsonByte, _ := json.Marshal(app)
	for _, v := range nodes {
		url := v.Url + "/app/updateApp"
		resp, err := http.NewRequest(http.MethodPut, url, strings.NewReader(string(jsonByte)))
		resp.Header.Add("Content-Type", contentType)
		_, err = http.DefaultClient.Do(resp)
		if err != nil {
			global.LOG.Error("远程更新失败!", zap.Any("err", err))
			return err
		}
		resp.Body.Close()
	}

	return
}

// 如果使用工作流功能 需要打开下方注释 并到initialize的workflow中进行注册 且必须指定TableName
// type AppWorkflow struct {
// 	// 工作流操作结构体
// 	WorkflowBase      `json:"wf"`
// 	App   `json:"business"`
// }

// func (App) TableName() string {
// 	return "app"
// }

// 工作流注册代码

// initWorkflowModel内部注册
// model.WorkflowBusinessStruct["app"] = func() model.GVA_Workflow {
//   return new(model.AppWorkflow)
// }

// initWorkflowTable内部注册
// model.WorkflowBusinessTable["app"] = func() interface{} {
// 	return new(model.App)
// }
