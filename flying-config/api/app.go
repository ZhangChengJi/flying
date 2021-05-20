package api

import (
	"context"
	"flying-config/constant"
	"flying-config/global"
	"flying-config/logic"
	"flying-config/model/response"
	proto "flying-config/proto/app"
	re "flying-config/proto/common"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"
)

type App struct{}

func (s *App) Lists(ctx context.Context, request *re.Request) (*re.Responses, error) {
	panic("implement me")
}

// 生成curd代码
func (s *App) Create(ctx context.Context, req *proto.App) (*re.Response, error) {

	if apps, _ := logic.FindAppIds(req.AppId); len(apps) > 0 {
		global.LOG.Error(constant.APPID_DUPLICATE.String(), zap.Any("err", nil))
		return response.BuildResult(constant.APPID_DUPLICATE), nil
	}
	err := logic.CreateApp(req)
	if err != nil {
		global.LOG.Error(constant.ERROR.String(), zap.Any("err", err))
		return response.BuildResult(constant.CREATE_FAILED), nil
	}

	return response.BuildResult(constant.CREATE_SUCCESS), nil
}

func (s *App) Delete(ctx context.Context, req *proto.App) (*re.Response, error) {
	app, err := logic.FindAppId(req.AppId)
	if err == nil {
		if err := logic.DeleteApp(app); err != nil {
			global.LOG.Error(constant.ERROR.String(), zap.Any("err", err))
			return response.BuildResult(constant.DELETE_FAILED), nil
		} else {
			return response.BuildResult(constant.DELETE_SUCCESS), nil
		}

	} else {
		global.LOG.Error(constant.NOT_FIND_APP.String(), zap.Any("err", err))
		return response.BuildResult(constant.NOT_FIND_APP), nil
	}

}

func (s *App) DeleteById(ctx context.Context, req *proto.App) (*re.Response, error) {
	data, err := logic.DeleteByIdApp(req)
	return response.OkWithData(data), err
}

func (s *App) Update(ctx context.Context, req *proto.App) (*re.Response, error) {
	app, err := logic.FindAppId(req.AppId)
	if err == nil {
		app.Name = req.Name
		if err := logic.UpdateApp(app); err != nil {
			global.LOG.Error(constant.UpdateError, zap.Any("err", err))
			return response.BuildResult(constant.UPDATE_FAILED), nil
		} else {
			return response.BuildResult(constant.UPDATE_SUCCESS), nil
		}
	} else {
		global.LOG.Error(err.Error(), zap.Any("err", err))
		return response.BuildResult(constant.NOT_FIND_APP), nil
	}

}

func (s *App) Find(ctx context.Context, req *emptypb.Empty) (*re.Response, error) {
	//data, err := logic.FindApp(req)
	return response.OkWithData(nil), nil
}

//func (s *App) Lists(ctx context.Context, req *proto.Request) (*proto.Response, error) {
//    data, total, err := logic.GetListApp(req)
//
//	var any = make([]*anypb.Any, len(data))
//	for k, r := range data {
//		any[k], err = ptypes.MarshalAny(r)
//	}
//
//	return response.OkWithData(any), err
//}
//
//
