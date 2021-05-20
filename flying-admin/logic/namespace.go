package logic

import (
	"bytes"
	"context"
	"flying-admin/constant"
	"flying-admin/global"
	"flying-admin/model"
	"flying-admin/model/request"
	pb "flying-admin/proto/namespace"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"go.uber.org/zap"
	"net/http"
	"time"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateNamespace
//@description: 创建Namespace记录
//@param: namespace model.Namespace
//@return: err error

func CreateNamespace(namespace *model.Namespace, nodes []*model.Node) (code int32, msg string) {
	tx := global.DB.Begin()
	var err error
	var buffer bytes.Buffer
	var app model.App
	err = tx.Where("app_id=?", namespace.AppId).First(&app).Error
	if err != nil {
		code = constant.NOT_FIND_APP
		return
	}
	for _, node := range nodes {
		resp, err := node.Namespace().Create(context.Background(), &pb.Namespace{AppId: app.AppId, Name: namespace.Name, Format: namespace.Format, Value: namespace.Value})
		if err != nil {
			global.LOG.Error(err.Error(), zap.Any("err", err))
			code = constant.CREATE_FAILED
			return
		}
		if resp.Code == constant.NOT_FIND_APP {
			buffer.WriteString(node.Name + " ")
		}

	}
	msg = buffer.String()
	if len(msg) > 0 {
		return constant.NOT_FIND_APP, msg
	} else {
		return constant.CREATE_SUCCESS, ""
	}

}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteNamespace
//@description: 删除Namespace记录
//@param: namespace model.Namespace
//@return: err error

func DeleteNamespace(namespace *model.Namespace, node *model.Node) (code int32, err error) {

	resp, err := node.Namespace().Delete(context.Background(), &pb.Namespace{Id: int64(namespace.Id)})
	if err != nil {
		code = constant.DELETE_FAILED
		return
	}
	code = resp.Code
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteNamespaceByIds
//@description: 批量删除Namespace记录
//@param: ids request.IdsReq
//@return: err error

func DeleteNamespaceByIds(ids request.IdsReq) (err error) {
	err = global.DB.Delete(&[]model.Namespace{}, "id in ?", ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateNamespace
//@description: 更新Namespace记录
//@param: namespace *model.Namespace
//@return: err error

func UpdateNamespace(namespace *model.Namespace, node *model.Node) (code int32, err error) {

	resp, err := node.Namespace().Update(context.Background(), &pb.Namespace{Id: int64(namespace.Id), AppId: namespace.AppId, Name: namespace.Name, Format: namespace.Format, Value: namespace.Value})
	if err != nil {
		global.LOG.Error("远程更新失败,"+node.Name+"环境", zap.Any("err", err))
		code = constant.UPDATE_FAILED
		return
	}
	code = resp.Code
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetNamespace
//@description: 根据id获取Namespace记录
//@param: id uint
//@return: err error, namespace model.Namespace

func GetNamespace(id int64, node *model.Node) (namespace *model.Namespace, err error) {
	resp, err := node.Namespace().Find(context.Background(), &pb.Namespace{Id: id})
	if resp.Code == http.StatusOK {
		var nm = new(pb.Namespace)
		err = ptypes.UnmarshalAny(resp.Data, nm)
		c, _ := TimestampToTime(nm.CreateTime)
		u, _ := TimestampToTime(nm.UpdateTime)
		r, _ := TimestampToTime(nm.ReleaseTime)
		namespace = &model.Namespace{
			Id:          uint(nm.Id),
			AppId:       nm.AppId,
			CreateTime:  c,
			UpdateTime:  u,
			ReleaseTime: r,
			Format:      nm.Format,
			Status:      nm.Status,
			Name:        nm.Name,
			Value:       nm.Value,
		}

	}
	return

}
func RepeatName(namespace *model.Namespace, nodes []*model.Node) (code int32, msg string, err error) {
	var buffer bytes.Buffer
	for _, node := range nodes {
		resp, _ := node.Namespace().RepeatName(context.Background(), &pb.Namespace{AppId: namespace.AppId, Name: namespace.Name})

		if resp.Code != http.StatusOK {
			buffer.WriteString(node.Name + " ")
			code = resp.Code
		}
	}
	msg = buffer.String()
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetNamespaceInfoList
//@description: 分页获取Namespace记录
//@param: info request.NamespaceSearch
//@return: err error, list interface{}, total int64

func GetNamespaceInfoList(info request.NamespaceSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.DB.Model(&model.Namespace{})
	var namespaces []model.Namespace
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&namespaces).Error
	return err, namespaces, total
}

func GetAppNamespace(node *model.Node, namespace *model.Namespace) (err error, namespaces []*model.Namespace) {
	resp, _ := node.Namespace().Lists(context.Background(), &pb.Namespace{AppId: namespace.AppId})
	if resp.Code == http.StatusOK {
		var nms = make([]pb.Namespace, len(resp.Data))
		for k, r := range resp.Data {
			err = ptypes.UnmarshalAny(r, &nms[k])
		}
		for _, nm := range nms {

			c, _ := TimestampToTime(nm.CreateTime)
			u, _ := TimestampToTime(nm.UpdateTime)
			r, _ := TimestampToTime(nm.ReleaseTime)
			pm := &model.Namespace{
				Id:          uint(nm.Id),
				AppId:       nm.AppId,
				CreateTime:  c,
				UpdateTime:  u,
				ReleaseTime: r,
				Format:      nm.Format,
				Status:      nm.Status,
				Name:        nm.Name,
				Value:       nm.Value,
			}
			namespaces = append(namespaces, pm)
		}
	}

	return
}
func TimestampToTime(ts *timestamp.Timestamp) (t *time.Time, err error) {
	if ts != nil {
		t1, _ := ptypes.Timestamp(ts)
		t = &t1
	} else {
		return nil, nil
	}
	return
}
func FindRemoteNamespace(namespace *model.Namespace, nodes []*model.Node) (message string, err error) {
	var buffer bytes.Buffer
	for _, node := range nodes {
		resp, _ := node.Namespace().RepeatName(context.Background(), &pb.Namespace{AppId: namespace.AppId, Name: namespace.Name})
		if resp.Code != http.StatusOK {
			buffer.WriteString(node.Name + " ")
		}
	}
	message = buffer.String()
	return
}
func SyncNamespace(namespace *model.Namespace, nodes []*model.Node) (code int32, err error) {
	for _, node := range nodes {
		resp, err := node.Namespace().Sync(context.Background(), &pb.Namespace{AppId: namespace.AppId, Name: namespace.Name, Format: namespace.Format, Value: namespace.Value})
		if err != nil {
			code = constant.SYNC_FAILED
			break
		}
		if resp.Code != constant.SYNC_SUCCESS {
			return resp.Code, err
		}
		code = resp.Code
	}

	return
}
