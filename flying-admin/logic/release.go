package logic

import (
	"context"
	"flying-admin/constant"
	"flying-admin/global"
	"flying-admin/model"
	"flying-admin/model/request"
	pb "flying-admin/proto/release"
	"github.com/golang/protobuf/ptypes"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateRelease
//@description: 创建Release记录
//@param: release model.Release
//@return: err error

func CreateRelease(node *model.Node, release request.Release) (code int32, err error) {

	timeProto := ptypes.TimestampNow()
	re := &pb.Release{
		Id:         int64(release.Id),
		Name:       release.ReleaseName,
		CreateTime: timeProto,
	}
	resp, err := node.Release().Create(context.Background(), re)
	if err != nil {
		code = constant.RELEASE_FAILED
		return
	}
	code = resp.Code
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteRelease
//@description: 删除Release记录
//@param: release model.Release
//@return: err error

func DeleteRelease(release model.Release) (err error) {
	err = global.DB.Delete(release).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteReleaseByIds
//@description: 批量删除Release记录
//@param: ids request.IdsReq
//@return: err error

func DeleteReleaseByIds(ids request.IdsReq) (err error) {
	err = global.DB.Delete(&[]model.Release{}, "id in ?", ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateRelease
//@description: 更新Release记录
//@param: release *model.Release
//@return: err error

func UpdateRelease(release model.Release) (err error) {
	err = global.DB.Save(&release).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetRelease
//@description: 根据id获取Release记录
//@param: id uint
//@return: err error, release model.Release

func GetRelease(id uint) (err error, release model.Release) {
	err = global.DB.Where("id = ?", id).First(&release).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetReleaseInfoList
//@description: 分页获取Release记录
//@param: info request.ReleaseSearch
//@return: err error, list interface{}, total int64

//func GetReleaseInfoList(info request.ReleaseSearch) (err error, list interface{}, total int64) {
//	limit := info.PageSize
//	offset := info.PageSize * (info.Page - 1)
//	// 创建db
//	db := global.DB.Model(&model.Release{})
//	var releases []model.Release
//	// 如果有条件搜索 下方会自动创建搜索语句
//	err = db.Count(&total).Error
//	err = db.Limit(limit).Offset(offset).Find(&releases).Error
//	return err, releases, total
//}
//func GetReleaseValue(client request.Client) (err error, release model.Release) {
//	err = global.DB.Where(&model.Release{AppId: client.AppId, NamespaceName: client.Namespace}).Order("update_time").First(&release).Error
//	return
//}
