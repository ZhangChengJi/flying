package utils

import (
	"flying-admin/global"
	"flying-admin/model"
	app "flying-admin/proto/app"
	"flying-admin/proto/client"
	"flying-admin/proto/namespace"
	"flying-admin/proto/release"
	"fmt"
	"sync"
)

var (
	Y sync.RWMutex
)

func AddNode(event Event) {
	Y.Lock()
	defer Y.Unlock()
	node := event.Object.(model.Node)
	con, err := model.NewConn(node.Url)
	if err != nil {
		global.LOG.Error("获取连接错误")
	}
	model.Conn[node.Key] = con
	svc := &model.GrpcService{
		App:       app.NewAppServiceClient(con),
		Namespace: namespace.NewNamespaceServiceClient(con),
		Release:   release.NewReleaseServiceClient(con),
		Client:    client.NewClientServiceClient(con),
	}
	model.M1[node.Key] = svc
	fmt.Println(model.Conn)
}

func DelNode(event Event) {
	y.Lock()
	node := event.Object.(model.Node)
	defer y.Unlock()
	delete(model.Conn, node.Key)
	delete(model.M1, node.Key)

}

//func GetNode() []*model.Node{
//	y.RLock()
//	defer  y.RUnlock()
//	return model.NodeList
//}

/**
 */
