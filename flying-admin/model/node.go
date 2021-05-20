// 自动生成模板Node
package model

import (
	"context"
	"flying-admin/global"
	app "flying-admin/proto/app"
	client "flying-admin/proto/client"
	namespace "flying-admin/proto/namespace"
	release "flying-admin/proto/release"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/protobuf/types/known/emptypb"
	"time"
)

// 如果含有time.Time 请自行import time包
type Node struct {
	Id         uint       `json:"id" form:"id" gorm:"column:id;comment:id;"`
	Name       string     `json:"name" form:"name" gorm:"column:name;comment:环境名称;type:varchar(30);size:30;"`
	Url        string     `json:"url" form:"url" gorm:"column:url;comment:环境地址;type:varchar(255);size:255;"`
	Key        string     `json:"key" form:"key" gorm:"column:key;comment:环境唯一标示;type:varchar(255);size:255;"`
	Status     bool       `json:"status" form:"status" gorm:"column:status;comment:环境连接状态;"`
	Content    string     `json:"content" form:"content" gorm:"column:content;comment:环境说明;type:varchar(255);size:255;"`
	CreateTime *time.Time `json:"createTime" form:"createTime" gorm:"column:create_time;comment:创建时间;default:null;"`
	UpdateTime *time.Time `json:"updateTime" form:"updateTime" gorm:"column:update_time;comment:更新时间;default:null;"`
}

var (
	M1   map[string]*GrpcService
	Conn map[string]*grpc.ClientConn
)

type GrpcService struct {
	App       app.AppServiceClient
	Namespace namespace.NamespaceServiceClient
	Release   release.ReleaseServiceClient
	Client    client.ClientServiceClient
	B         bool
}

type SyncId struct {
	Id    uint   `json:"id" form:"id"`
	Nodes []uint `json:"nodes" form:"nodes"`
	NId   uint   `json:"nId" form:"nId"`
}

func (Node) TableName() string {
	return "node"
}

func NewGrpcService() {
	Conn = make(map[string]*grpc.ClientConn)
	M1 = make(map[string]*GrpcService)
	var nodes []*Node
	global.DB.Find(&nodes)
	var bo = true

	for i := 0; i < len(nodes); i++ {
		con, err := NewConn(nodes[i].Url)
		if err != nil {
			global.LOG.Error("获取连接错误")
		}
		Conn[nodes[i].Key] = con
		t := client.NewClientServiceClient(con)
		_, err = t.Ping(context.Background(), &emptypb.Empty{})
		//TODO: 启动时候检查连接状态并且修改数据库
		if err != nil {
			bo = false
			global.LOG.Error(nodes[i].Url + ":测试无法连接")
			if nodes[i].Status {
				UpdateStatus(nodes[i].Key, false)
			}
		} else {
			bo = true
			if !nodes[i].Status {
				UpdateStatus(nodes[i].Key, true)
			}

		}
		svc := &GrpcService{
			App:       app.NewAppServiceClient(con),
			Namespace: namespace.NewNamespaceServiceClient(con),
			Release:   release.NewReleaseServiceClient(con),
			Client:    t,
			B:         bo,
		}
		M1[nodes[i].Key] = svc
	}

}

func UpdateStatus(key string, b bool) {
	var node Node
	var err error
	err = global.DB.Debug().Where(&Node{Key: key}).First(&node).Error
	if err != nil {
		global.LOG.Error("没 查询到该key的node")
		return
	}
	node.Status = b
	global.DB.Debug().Save(&node)
}

func NewConn(url string) (*grpc.ClientConn, error) {
	creds, err := credentials.NewClientTLSFromFile(global.GVA_CONFIG.System.PrivateKey, "flying.com")
	if err != nil {
		grpclog.Fatalf("Failed to create TLS credentials %v", err)
	}
	conn, err := grpc.Dial(url, grpc.WithTransportCredentials(creds))
	if err != nil {
		grpclog.Fatalln(err)
	}
	return conn, nil
}

/**   获取grpc的结构体，通过结构体调用具体的远程接口   **/

func (node *Node) App() app.AppServiceClient {
	return M1[node.Key].App
}
func (node *Node) Namespace() namespace.NamespaceServiceClient {
	return M1[node.Key].Namespace
}
func (node *Node) Release() release.ReleaseServiceClient {
	return M1[node.Key].Release
}
func (node *Node) Client() client.ClientServiceClient {
	return M1[node.Key].Client
}

/**  更新Conn 连接集合 **/
