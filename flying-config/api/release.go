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
	proto "flying-config/proto/release"
	"github.com/golang/protobuf/ptypes"
	"go.uber.org/zap"
	"math/rand"
	"strconv"
	"time"
)

type Release struct{}

func (s *Release) Create(ctx context.Context, req *proto.Release) (*re.Response, error) {
	namespace, err := logic.FindNamespace(uint(req.Id))
	if err == nil {
		if !namespace.Status {
			rand.Seed(time.Now().UnixNano())
			namespace.Status = true
			timeGo, _ := ptypes.Timestamp(req.CreateTime)
			namespace.ReleaseTime = &timeGo
			release := &model.Release{
				AppId:         namespace.AppId,
				Name:          req.Name,
				NamespaceName: namespace.Name,
				ReleaseKey:    strconv.Itoa(rand.Intn(1000000000)),
				Format:        namespace.Format,
				Value:         namespace.Value,
				CreateTime:    &timeGo,
			}

			if err := logic.CreateRelease(release, namespace); err != nil {
				global.LOG.Error(constant.ReleaseError, zap.Any("err", err))
				return response.BuildResult(constant.RELEASE_FAILED), nil
			} else {
				listener.AddListener(release.AppId)
				return response.BuildResult(constant.RELEASE_SUCCESS), nil
			}
		} else {
			return response.BuildResult(constant.RELEASE_OTHER), nil
		}
	} else {
		return response.BuildResult(constant.NAMESPACE_RELEASE_NOT_FIND), nil
	}
}

func (s *Release) Delete(ctx context.Context, req *proto.Release) (*re.Response, error) {
	//data, err := logic.DeleteRelease(req)
	return nil, nil
}

func (s *Release) DeleteById(ctx context.Context, req *proto.Release) (*re.Response, error) {
	//data, err := service.DeleteByIdRelease(req)
	return nil, nil
}

func (s *Release) Update(ctx context.Context, req *proto.Release) (*re.Response, error) {
	//data, err := service.UpdateRelease(req)
	return nil, nil
}

func (s *Release) Find(ctx context.Context, req *proto.Release) (*re.Response, error) {
	//data, err := service.FindRelease(req)
	return nil, nil
}

func (s *Release) Lists(ctx context.Context, req *re.Request) (*re.Responses, error) {
	//data, total, err := service.GetListRelease(req)
	//
	//var any = make([]*anypb.Any, len(data))
	//for k, r := range data {
	//	any[k], err = ptypes.MarshalAny(r)
	//}

	return nil, nil
}
