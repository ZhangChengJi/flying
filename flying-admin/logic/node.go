package logic

import (
	"flying-admin/global"
	model "flying-admin/model"
	"flying-admin/model/request"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateNode
//@description: 创建Node记录
//@param: node model.Node
//@return: err error

func CreateNode(node *model.Node) (err error) {
	err = global.DB.Create(&node).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteNode
//@description: 删除Node记录
//@param: node model.Node
//@return: err error

func DeleteNode(node model.Node) (err error) {
	err = global.DB.Delete(node).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteNodeByIds
//@description: 批量删除Node记录
//@param: ids request.IdsReq
//@return: err error

func DeleteNodeByIds(ids request.IdsReq) (err error) {
	err = global.DB.Delete(&[]model.Node{}, "id in ?", ids.Ids).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateNode
//@description: 更新Node记录
//@param: node *model.Node
//@return: err error

func UpdateNode(node *model.Node) (err error) {
	err = global.DB.Save(&node).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetNode
//@description: 根据id获取Node记录
//@param: id uint
//@return: err error, node model.Node

func GetNode(id uint) (node *model.Node, err error) {
	err = global.DB.Where("id = ?", id).First(&node).Error
	return
}
func IfRepeat(node *model.Node) bool {
	var count int64
	global.DB.Model(&model.Node{}).Where("id!=? and name = ?", node.Id, node.Name).Count(&count)
	return count > 0
}
func IfExists(node *model.Node) (i int) {
	var nodes []*model.Node
	global.DB.Where(&model.Node{Name: node.Name}).Find(&nodes)
	if len(nodes) > 0 {
		return 1
	}
	global.DB.Where(&model.Node{Url: node.Url}).Find(&nodes)
	if len(nodes) > 0 {
		return 2
	}
	return
}
func GetAppNodes(ids []uint) (nodes []*model.Node, err error) {
	err = global.DB.Where("id IN (?)", ids).Find(&nodes).Error
	return
}
func GetAll() (err error, nodes []*model.Node) {
	err = global.DB.Find(&nodes).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetNodeInfoList
//@description: 分页获取Node记录
//@param: info request.NodeSearch
//@return: err error, list interface{}, total int64

func GetNodeInfoList(info request.NodeSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.DB.Model(&model.Node{})
	var nodes []model.Node
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&nodes).Error
	return err, nodes, total
}

func UpdateStatus(key string, b bool) {
	var node model.Node
	var err error
	err = global.DB.Debug().Where(&model.Node{Key: key}).First(&node).Error
	if err != nil {
		global.LOG.Error("没 查询到该key的node")
		return
	}
	node.Status = b
	global.DB.Debug().Save(&node)
}
