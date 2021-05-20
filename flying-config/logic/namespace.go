package logic

import (
	"bytes"
	"encoding/gob"
	"errors"
	"flying-config/global"
	"flying-config/model"
	proto "flying-config/proto/namespace"
	"flying-config/utils"
)

func CreateNamespace(namespace *model.Namespace) (err error) {
	err = global.DB.Create(&namespace).Error
	return
}

func DeleteNamespace(namespace *model.Namespace) (err error) {
	err = global.DB.Delete(&namespace).Error
	err = DeleteAllRelease(namespace.AppId, namespace.Name)
	return
}

func DeleteByIdNamespace(req *proto.Namespace) (*proto.Namespace, error) {
	err := global.DB.First(&req).Delete(&req).Error
	return req, err
}

func FindNamespace(id uint) (namespace *model.Namespace, err error) {
	err = global.DB.Where("id = ?", id).First(&namespace).Error
	return
}

//
func UpdateNamespace(namespace *model.Namespace) (err error) {
	err = global.DB.Save(&namespace).Error
	return
}

func RepeatName(appId, name string) (err error) {
	var namespaces []*model.Namespace
	err = global.DB.Where("app_id = ? and name = ?", appId, name).Find(&namespaces).Error
	if len(namespaces) > 0 {
		return errors.New("duplicate namespace name")
	}
	return
}

func DeepCopy(dst, src interface{}) error {
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(src); err != nil {
		return err
	}
	return gob.NewDecoder(bytes.NewBuffer(buf.Bytes())).Decode(dst)
}

func GetAppNamespace(appId string) (namespaces []*proto.Namespace, err error) {
	var nms []*model.Namespace
	err = global.DB.Where("app_id = ? ", appId).Find(&nms).Error
	for _, nm := range nms {

		c, _ := utils.TimeToTimestamp(nm.CreateTime)
		u, _ := utils.TimeToTimestamp(nm.UpdateTime)
		r, _ := utils.TimeToTimestamp(nm.ReleaseTime)
		pm := &proto.Namespace{
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
		namespaces = append(namespaces, pm)
	}
	return
}
func Sync(req *proto.Namespace) (err error) {
	var namespace model.Namespace

	err = global.DB.Where("app_id = ? and name = ?", req.AppId, req.Name).First(&namespace).Error
	if namespace != (model.Namespace{}) {
		namespace.AppId = req.AppId
		namespace.Format = req.Format
		namespace.Name = req.Name
		namespace.Value = req.Value
		namespace.Status = req.Status
		err = global.DB.Save(namespace).Error
	} else {
		namespace := &model.Namespace{
			AppId:  req.AppId,
			Format: req.Format,
			Name:   req.Name,
			Value:  req.Value,
			Status: req.Status,
		}
		err = global.DB.Create(&namespace).Error

	}
	return
}

//
//func FindNamespace(req *proto.Namespace) (*proto.Namespace, error) {
//	err := global.DB.Where(&req).First(&req).Error
//	return req, err
//}
//
//func GetListNamespace(req *proto.Request) (result []*proto.Namespace, total int64, err error) {
//	if req.PageSize == 0 {
//		req.PageSize = constant.PAGESIZE
//	}
//	if req.Page == 0 {
//		req.Page = constant.PAGE
//	}
//	limit := req.PageSize
//	offset := req.PageSize * (req.Page - 1)
//
//    unmarshal := &proto.Namespace{}
//	err = ptypes.UnmarshalAny(req.Query, unmarshal)
//	db := global.DB.Model(&result)
//
//
//    if unmarshal.AId != 0 {
//            db = db.Where("a_id = ?", unmarshal.AId)
//    }
//
//
//
//    if unmarshal.AppId != "" {
//        db = db.Where("app_id LIKE ?", "%"+unmarshal.AppId+"%")
//    }
//
//
//
//    if unmarshal.CreateTime != 0 {
//            db = db.Where("create_time = ?", unmarshal.CreateTime)
//    }
//
//
//
//    if unmarshal.Format != "" {
//        db = db.Where("format LIKE ?", "%"+unmarshal.Format+"%")
//    }
//
//
//
//    if unmarshal.Id != 0 {
//            db = db.Where("id = ?", unmarshal.Id)
//    }
//
//
//
//    if unmarshal.Name != "" {
//        db = db.Where("name LIKE ?", "%"+unmarshal.Name+"%")
//    }
//
//
//
//    if unmarshal.NodeId != 0 {
//            db = db.Where("node_id = ?", unmarshal.NodeId)
//    }
//
//
//
//    if unmarshal.ReleaseTime != 0 {
//            db = db.Where("release_time = ?", unmarshal.ReleaseTime)
//    }
//
//
//
//    if unmarshal.Status != 0 {
//            db = db.Where("status = ?", unmarshal.Status)
//    }
//
//
//
//    if unmarshal.UpdateTime != 0 {
//            db = db.Where("update_time = ?", unmarshal.UpdateTime)
//    }
//
//
//
//    if unmarshal.Value != 0 {
//            db = db.Where("value = ?", unmarshal.Value)
//    }
//
//
//
//	err = db.Count(&total).Error
//
//	if err != nil {
//		return result, total, err
//	} else {
//		db = db.Limit(limit).Offset(offset)
//		if req.OrderKey != "" {
//			var OrderStr string
//			if req.OrderDesc != "" {
//				OrderStr = req.OrderKey + " desc"
//			} else {
//				OrderStr = req.OrderKey
//			}
//			err = db.Order(OrderStr).Find(&result).Error
//		} else {
//			err = db.Order("id").Find(&result).Error
//		}
//	}
//
//	return result, total, err
//}
