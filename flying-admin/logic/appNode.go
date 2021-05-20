package logic

import (
	"flying-admin/global"
	"flying-admin/model"
	"flying-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateAppNode
//@description: 创建AppNode记录
//@param: appNode model.AppNode
//@return: err error

func CreateAppNode(appNode model.AppNode) (err error) {
	err = global.DB.Create(&appNode).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteAppNode
//@description: 删除AppNode记录
//@param: appNode model.AppNode
//@return: err error

func DeleteAppNode(appNode model.AppNode) (err error) {
	err = global.DB.Delete(appNode).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteAppNodeByIds
//@description: 批量删除AppNode记录
//@param: ids request.IdsReq
//@return: err error

func DeleteAppNodeByIds(ids request.IdsReq) (err error) {
	err = global.DB.Delete(&[]model.AppNode{}, "id in ?", ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateAppNode
//@description: 更新AppNode记录
//@param: appNode *model.AppNode
//@return: err error

func UpdateAppNode(appNode model.AppNode) (err error) {
	err = global.DB.Save(&appNode).Error
	return err
}
func GetNodeApp(id uint) (err error, count int64) {
	err = global.DB.Model(&model.AppNode{}).Where("node_id = ?", id).Count(&count).Error
	return
}
func GetAppAllNodes(id uint) (nodes []*model.Node, err error) {
	var ids []uint
	err = global.DB.Where("id in (?)", global.DB.Table("app_node").Select("node_id").Where("a_id = ?", id).Find(&ids)).Find(&nodes).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetAppNode
//@description: 根据id获取AppNode记录
//@param: id uint
//@return: err error, appNode model.AppNode

func GetAppNode(appId string) (err error, node []*model.Node, app *model.App) {
	var appNodes []model.AppNode
	var ids []uint

	err = global.DB.Where("app_id = ?", appId).Find(&appNodes).Error
	if err == nil {
		for _, node := range appNodes {
			ids = append(ids, node.NodeId)
		}
		node, err = GetAppNodes(ids)
		err = global.DB.Where("app_id = ?", appId).First(&app).Error
	}

	return
}

func GetSyncNode(sync *request.SyncId) (err error, node []*model.Node) {
	var appNodes []model.AppNode
	var ids []uint
	err = global.DB.Where("app_id = ?", sync.AppId).Not("node_id = ?", sync.NodeId).Find(&appNodes).Error
	if err == nil {
		for _, node := range appNodes {
			ids = append(ids, node.NodeId)
		}
		node, err = GetAppNodes(ids)
	}
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetAppNodeInfoList
//@description: 分页获取AppNode记录
//@param: info request.AppNodeSearch
//@return: err error, list interface{}, total int64

func GetAppNodeInfoList(info request.AppNodeSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.DB.Model(&model.AppNode{})
	var appNodes []model.AppNode
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&appNodes).Error
	return err, appNodes, total
}
func GetDeleteAppNode(id uint) (err error, node []*model.Node) {
	var appNodes []model.AppNode
	var ids []uint

	err = global.DB.Where("a_id = ?", id).Find(&appNodes).Error
	if err == nil {
		for _, node := range appNodes {
			ids = append(ids, node.NodeId)
		}
		node, err = GetAppNodes(ids)
	}
	return
}
