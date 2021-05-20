// 自动生成模板Namespace
package model

import (
	"encoding/json"
	"flying-admin/global"
	"flying-admin/model/response"
	"fmt"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// 如果含有time.Time 请自行import time包
type Namespace struct {
	Id          uint       `json:"id" form:"id" gorm:"column:id;comment:id;"`
	AId         uint       `json:"aId" form:"aId" gorm:"column:a_id;comment:应用ID;type:int;size:10;"`
	AppId       string     `json:"appId" form:"appId" gorm:"column:app_id;comment:appId;type:varchar(60);size:60;"`
	Format      string     `json:"format" form:"format" gorm:"column:format;comment:格式;type:varchar(60);size:60;"`
	NodeId      uint       `json:"nodeId" form:"nodeId" gorm:"column:node_id;comment:环境ID;type:int;size:10;"`
	Name        string     `json:"name" form:"name" gorm:"column:name;comment:命名空间名称;type:varchar(255);size:255;"`
	Value       string     `json:"value" form:"value" gorm:"column:value;comment:配置信息;type:longtext;"`
	Status      bool       `json:"status" form:"status" gorm:"column:status;comment:发布状态;"`
	CreateTime  *time.Time `json:"createTime" form:"createTime" gorm:"column:create_time;comment:创建时间;default:null;"`
	UpdateTime  *time.Time `json:"updateTime" form:"updateTime" gorm:"column:update_time;comment:更新时间;default:null;"`
	ReleaseTime *time.Time `json:"releaseTime" form:"releaseTime" gorm:"column:release_time;comment:发布时间;default:null;"`
}

//func (Namespace) TableName() string {
//	return "namespace"
//}

type Client interface {
	RemoteRelease(do *ReleaseDo, node *Node) (err error)
	CreateRemoteNamespace(nodes []*Node) (err error)
	UpdateRemoteNamespace(node *Node) (err error)
	DeleteRemoteNamespace(node *Node) (err error)
	SyncRemoteNamespace(node *Node) (err error)
	FindRemoteNamespace(node *Node) (err error, namespaces Namespace)
}

func (namespace *Namespace) CreateRemoteNamespace(nodes []*Node) (err error) {
	contentType := "application/json"
	namespace.AId = 0
	jsonByte, _ := json.Marshal(namespace)
	for _, v := range nodes {

		resp, err := http.Post(v.Url+"/namespace/createNamespace", contentType, strings.NewReader(string(jsonByte)))
		if err != nil {
			global.LOG.Error("远程插入失败!", zap.Any("err", err))
			return err
		}
		resp.Body.Close()
	}

	return err

}

type ReleaseDo struct {
	AppId         string `json:"appId" form:"appId"`
	NamespaceName string `json:"namespaceName" form:"namespaceName"`
	ReleaseName   string `json:"releaseName" form:"releaseName"`
}

func (namespace *Namespace) RemoteRelease(do *ReleaseDo, node *Node) (err error) {
	contentType := "application/json"
	jsonByte, _ := json.Marshal(do)
	resp, err := http.Post(node.Url+"/release/createRelease", contentType, strings.NewReader(string(jsonByte)))
	if err != nil {
		global.LOG.Error("远程插入失败!", zap.Any("err", err))
		return err
	}
	defer resp.Body.Close()
	return
}

func (namespace *Namespace) DeleteRemoteNamespace(node *Node) (err error) {
	contentType := "application/json"
	jsonByte, _ := json.Marshal(namespace)
	url := node.Url + "/namespace/deleteNamespace"

	resp, err := http.NewRequest(http.MethodDelete, url, strings.NewReader(string(jsonByte)))
	resp.Header.Add("Content-Type", contentType)
	_, err = http.DefaultClient.Do(resp)
	if err != nil {
		global.LOG.Error("远程删除失败!", zap.Any("err", err))
		return err
	}
	defer resp.Body.Close()
	return
}

func (namespace *Namespace) SyncRemoteNamespace(node *Node) (err error) {
	contentType := "application/json"
	jsonByte, _ := json.Marshal(namespace)
	resp, err := http.Post(node.Url+"/namespace/syncNamespace", contentType, strings.NewReader(string(jsonByte)))
	if err != nil {
		global.LOG.Error("远程同步配置失败!", zap.Any("err", err))
		return err
	}
	defer resp.Body.Close()
	return
}
func (namespace *Namespace) FindRemoteNamespace(nodes []*Node) (err error, ok bool, message string) {
	for _, node := range nodes {
		apiUrl := node.Url + "/namespace/findNamespace"
		data := url.Values{}
		data.Set("appId", namespace.AppId)
		data.Set("namespaceName", namespace.Name)
		u, err := url.ParseRequestURI(apiUrl)
		if err != nil {
			fmt.Printf("get failed, err:%v\n", err)
			break
		}
		u.RawQuery = data.Encode()
		resp, err := http.Get(u.String())
		var namespace2 Namespace
		var response response.Response
		b, err := ioutil.ReadAll(resp.Body)
		if err == nil {
			json.Unmarshal(b, &response)
			data, _ := json.Marshal(response.Data)
			json.Unmarshal(data, &namespace2)
			resp.Body.Close()
			if namespace2 != (Namespace{}) {
				ok = true
				message += node.Name + "环境已经存在当前namespace，格式：" + namespace2.Format + ","
			}
		} else {
			break
		}
	}
	return
}
func (namespace *Namespace) UpdateRemoteNamespace(node *Node) (err error) {
	contentType := "application/json"
	jsonByte, _ := json.Marshal(namespace)
	url := node.Url + "/namespace/updateNamespace"

	resp, err := http.NewRequest(http.MethodPut, url, strings.NewReader(string(jsonByte)))
	resp.Header.Add("Content-Type", contentType)
	_, err = http.DefaultClient.Do(resp)
	if err != nil {
		global.LOG.Error("远程更新失败!", zap.Any("err", err))
		return err
	}
	defer resp.Body.Close()
	return
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
