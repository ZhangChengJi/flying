package api

import (
	"context"
	"flying-config/constant"
	"flying-config/global"
	"flying-config/listener"
	"flying-config/logic"
	"flying-config/model"
	"flying-config/model/response"
	proto "flying-config/proto/client"
	re "flying-config/proto/common"
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"
	"time"
)

type Client struct{}

//Ping
func (s *Client) Ping(ctx context.Context, req *emptypb.Empty) (resp *emptypb.Empty, err error) {
	return &emptypb.Empty{}, nil
}
func (s *Client) QueryApp(ctx context.Context, req *emptypb.Empty) (*re.Response, error) {
	apps, err := logic.QueryApp()
	if err != nil {
		global.LOG.Error(constant.QueryAppError)
	}
	if len(apps) > 0 {
		return response.BuildResult(constant.HISTORYLEGACY), nil
	} else {
		return response.Ok(), nil
	}
}

func (s *Client) Config(ctx context.Context, req *proto.Client) (*re.Response, error) {
	app, err := logic.FindAppId(req.AppId)
	if err != nil || app == (&model.App{}) {
		global.LOG.Error(fmt.Sprintf(constant.GetConfigNotAppId, req.AppId), zap.Any("err", err))
		return response.FailWithMessage(fmt.Sprintf(constant.GetConfigNotAppId, app.AppId)), nil

	} else {

		release, err := logic.GetReleaseValue(req)
		if err != nil {
			global.LOG.Error(fmt.Sprintf(constant.GetConfigNotNamespace, app.AppId, req.Namespace))
			return response.FailWithMessage(fmt.Sprintf(constant.GetConfigNotNamespace, app.AppId, req.Namespace)), nil
		} else {
			global.LOG.Info(fmt.Sprintf(constant.GetConfigOK, app.AppId, req.Namespace))
			flying := &proto.FlyingConfig{
				AppId:         release.AppId,
				NamespaceName: release.NamespaceName,
				ReleaseKey:    release.ReleaseKey,
				Format:        release.Format,
				Value:         release.Value,
			}
			return response.OkWithData(flying), nil
		}
	}

}

func (s *Client) Listener(ctx context.Context, req *proto.Client) (*re.Response, error) {
	w := listener.Watch(req.AppId, req.Namespace)
	time.AfterFunc(time.Second*45, func() {
		_ = w.Stop()
	})

	_, err := w.Next()
	if err != nil {
		return response.NotUpdated("no need to update"), nil

	}
	return response.OkWithMessage("need to be updated"), nil
}
