package utils

import (
	"context"
	"flying-admin/global"
	"flying-admin/logic"
	"flying-admin/model"
	client "flying-admin/proto/client"
	"flying-admin/proto/common"
	"google.golang.org/protobuf/types/known/emptypb"
	"time"
)

type Telnet struct {
	Node *model.Node
}

func ConnectTestTask() {
	go func() {
		var err error
		for {
			//TODO 12秒检查一次连接
			time.Sleep(time.Second * 12)
			for key, val := range model.M1 {
				_, err = val.Client.Ping(context.Background(), &emptypb.Empty{})
				if err != nil {
					global.LOG.Debug("环境无法连接修改状态")
					if val.B {
						Y.Lock()
						logic.UpdateStatus(key, false)
						val.B = false
						Y.Unlock()
					}
				} else {
					if !val.B {
						Y.Lock()
						logic.UpdateStatus(key, true)
						val.B = true
						Y.Unlock()
					}
				}
			}
		}
	}()

}

// Ping 只是用于 创建node 时候进行测试使用
func (t *Telnet) QueryApp() (resp *common.Response, err error) {
	c, err := model.NewConn(t.Node.Url)
	defer c.Close()
	if err != nil {
		global.LOG.Error("获取连接错误")
	}
	resp, err = client.NewClientServiceClient(c).QueryApp(context.Background(), &emptypb.Empty{})
	return
}

//func (g GrpcConn) Ping() (err error) {
//	_, err = g.Client.Ping(context.Background(), &emptypb.Empty{})
//	return
//}
