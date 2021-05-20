package api

import (
	"context"
	"flying-config/constant"
	"flying-config/global"
	"flying-config/listener"
	"flying-config/logic"
	"flying-config/model"
	"flying-config/model/response"
	re "flying-config/proto/common"
	proto "flying-config/proto/namespace"
	"flying-config/utils"
	"github.com/golang/protobuf/ptypes"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/anypb"
)

type Namespace struct{}

func (s *Namespace) RepeatName(ctx context.Context, req *proto.Namespace) (*re.Response, error) {
	if err := logic.RepeatName(req.AppId, req.Name); err != nil {
		return response.BuildResult(constant.NAMESPACE_REPEAT), nil
	} else {
		return response.Ok(), nil
	}
}

// 生成curd代码
func (s *Namespace) Create(ctx context.Context, req *proto.Namespace) (*re.Response, error) {
	app, err := logic.FindAppId(req.AppId)
	if err == nil {
		namespace := &model.Namespace{
			AId:    app.Id,
			AppId:  req.AppId,
			Format: req.Format,
			Name:   req.Name,
			Value:  req.Value,
			Status: req.Status,
		}
		err := logic.CreateNamespace(namespace)
		if err != nil {
			return response.BuildResult(constant.CREATE_FAILED), nil
		}
	} else {
		return response.BuildResult(constant.NOT_FIND_APP), nil
	}
	return response.BuildResult(constant.CREATE_SUCCESS), err
}

func (s *Namespace) Delete(ctx context.Context, req *proto.Namespace) (*re.Response, error) {
	namespace, err := logic.FindNamespace(uint(req.Id))
	if err == nil || *namespace != (model.Namespace{}) {
		if err := logic.DeleteNamespace(namespace); err != nil {
			global.LOG.Error(constant.DeleteError, zap.Any("err", err))
			return response.BuildResult(constant.DELETE_FAILED), nil
		} else {
			listener.AddListener(req.AppId)
			logic.DeleteAllRelease(namespace.AppId, namespace.Name)
			return response.BuildResult(constant.DELETE_SUCCESS), nil
		}
	} else {
		global.LOG.Error(constant.NotFindNamespace, zap.Any("err", err))
		return response.BuildResult(constant.NAMESPACE_DELETE_NOT_FIND), nil
	}

}

func (s *Namespace) DeleteById(ctx context.Context, req *proto.Namespace) (*re.Response, error) {
	logic.DeleteByIdNamespace(req)
	return response.FailWithMessage("创建失败"), nil
}

func (s *Namespace) Update(ctx context.Context, req *proto.Namespace) (*re.Response, error) {
	namespace, err := logic.FindNamespace(uint(req.Id))
	if err != nil {
		global.LOG.Error(constant.NotFindNamespace, zap.Any("err", err))
		return response.BuildResult(constant.NAMESPACE_UPDATE_NOT_FIND), nil

	}
	namespace.Status = false
	namespace.Value = req.Value
	if err := logic.UpdateNamespace(namespace); err != nil {
		global.LOG.Error(constant.UpdateError, zap.Any("err", err))
		return response.BuildResult(constant.UPDATE_FAILED), nil
	} else {
		return response.BuildResult(constant.UPDATE_SUCCESS), nil
	}

}

func (s *Namespace) Find(ctx context.Context, req *proto.Namespace) (*re.Response, error) {
	nm, err := logic.FindNamespace(uint(req.Id))
	if err == nil {
		c, _ := utils.TimeToTimestamp(nm.CreateTime)
		u, _ := utils.TimeToTimestamp(nm.UpdateTime)
		r, _ := utils.TimeToTimestamp(nm.ReleaseTime)
		namespace := &proto.Namespace{
			Id:          int64(nm.Id),
			AppId:       nm.AppId,
			CreateTime:  c,
			UpdateTime:  u,
			ReleaseTime: r,
			Format:      nm.Format,
			Status:      nm.Status,
			Name:        nm.Name,
			Value:       nm.Value,
		}
		return response.OkWithData(namespace), nil
	} else {
		return response.BuildResult(constant.RELEASE_NAMESPACE_SYNC_NOT_FOUND), nil
	}

}

func (s *Namespace) Lists(ctx context.Context, req *proto.Namespace) (*re.Responses, error) {
	nms, err := logic.GetAppNamespace(req.AppId)
	if err == nil {
		var any = make([]*anypb.Any, len(nms))
		for k, r := range nms {
			any[k], err = ptypes.MarshalAny(r)
		}
		return response.OkWithDatas(any), nil
	}
	return response.OkWithDatas(nil), nil

}
func (s *Namespace) Sync(ctx context.Context, req *proto.Namespace) (*re.Response, error) {
	_, err := logic.FindAppId(req.AppId)
	if err == nil {
		err := logic.Sync(req)
		if err != nil {
			return response.BuildResult(constant.SYNC_FAILED), nil
		}
	} else {
		return response.BuildResult(constant.RELEASE_APP_NO_FIND), nil
	}
	return response.BuildResult(constant.SYNC_SUCCESS), err
}
