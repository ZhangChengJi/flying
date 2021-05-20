package logic

import (
	"bytes"
	"context"
	"flying-admin/constant"
	"flying-admin/global"
	"flying-admin/model"
	"flying-admin/model/request"
	pb "flying-admin/proto/app"
	"fmt"
	"go.uber.org/zap"
)

func newClient(node *model.Node) (p pb.AppServiceClient, err error) {

	if err != nil {
		global.LOG.Error("node环境："+node.Url+"连接失败", zap.Any("err", err))
		return
	}

	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateApp
//@description: 创建App记录
//@param: app model.App
//@return: err error

func CreateApp(app *model.App, nodes []*model.Node) (code int32, err error) {

	tx := global.DB.Begin()
	var appNodes = make([]*model.AppNode, 0)
	for _, node := range nodes {
		resp, err := node.App().Create(context.Background(), &pb.App{AppId: app.AppId, Name: app.Name})
		if err != nil {
			return constant.CREATE_FAILED, err
		}
		if resp.Code != constant.CREATE_SUCCESS {
			global.LOG.Error(fmt.Sprintf("%s 环境"+resp.Msg, node.Name), zap.Any("err", err))
			return resp.Code, err

		} else {
			if app.Id != 0 {
				appNode := &model.AppNode{
					AppId:  app.AppId,
					NodeId: node.Id,
					AId:    app.Id,
				}
				appNodes = append(appNodes, appNode)
			} else {
				appNode := &model.AppNode{
					NodeId: node.Id,
				}
				appNodes = append(appNodes, appNode)
			}
		}
	}

	if app.Id == 0 {
		if len(appNodes) > 0 {
			if err1 := tx.Create(&app).Error; err1 == nil {
				for i := 0; i < len(appNodes); i++ {
					appNodes[i].AppId = app.AppId
					appNodes[i].AId = app.Id
				}
				if err1 = tx.Create(&appNodes).Error; err == nil {
					tx.Commit()
					code = constant.CREATE_SUCCESS
				} else {
					tx.Rollback()
					global.LOG.Error("本地创建失败", zap.Any("err", err))
					code = constant.CREATE_FAILED
				}
			} else {
				tx.Rollback()
				global.LOG.Error("本地创建失败", zap.Any("err", err))
				code = constant.CREATE_FAILED
			}
		}
	} else {
		if err = tx.Create(&appNodes).Error; err == nil {
			tx.Commit()
			code = constant.CREATE_SUCCESS
		} else {
			tx.Rollback()
			global.LOG.Error("本地创建失败", zap.Any("err", err))
			code = constant.CREATE_FAILED
		}
	}
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteApp
//@description: 删除App记录
//@param: app model.App
//@return: err error

func DeleteApp(app model.App, nodes []*model.Node) (code int32, err error) {
	tx := global.DB.Begin()
	for _, node := range nodes {
		resp, err := node.App().Delete(context.Background(), &pb.App{AppId: app.AppId})
		if err != nil {
			return constant.DELETE_FAILED, err
		}
		if resp.Code != constant.DELETE_SUCCESS {
			return resp.Code, err

		}
	}
	err = tx.Delete(app).Error
	err = tx.Where("a_id=?", app.Id).Delete(&model.AppNode{}).Error
	if err != nil {
		tx.Rollback()
		code = constant.DELETE_FAILED
	} else {
		tx.Commit()
		code = constant.DELETE_SUCCESS
	}
	return
}
func ClearAppNode(app *model.App, id []uint, nodes []*model.Node) (code int32, err error) {
	tx := global.DB.Begin()
	for _, node := range nodes {
		resp, err := node.App().Delete(context.Background(), &pb.App{AppId: app.AppId, Name: app.Name})
		if err != nil {
			global.LOG.Error("远程更新失败,"+node.Name+"环境", zap.Any("err", err))
			return constant.DELETE_FAILED, err
		}
		if resp.Code != constant.DELETE_SUCCESS {
			return resp.Code, err
		}
	}
	//删除本地 appNode表和namespace

	err = tx.Where("a_id = ? and node_id in ?", app.Id, id).Delete(&model.AppNode{}).Error
	if err != nil {
		tx.Rollback()
		return constant.CREATE_FAILED, err
	} else {
		tx.Commit()
	}

	return constant.CREATE_SUCCESS, err
}
func UpdateRemoteApp(app *model.App, nodes []*model.Node) (code int32, err error) {
	tx := global.DB.Begin()
	for _, node := range nodes {
		resp, err := node.App().Update(context.Background(), &pb.App{AppId: app.AppId, Name: app.Name})
		if err != nil {
			global.LOG.Error("远程更新失败,"+node.Name+"环境", zap.Any("err", err))
			return constant.UPDATE_FAILED, err
		}
		if resp.Code != constant.UPDATE_SUCCESS {
			return resp.Code, err
		}

	}
	err = tx.Save(&app).Error
	if err != nil {
		tx.Rollback()
		code = constant.UPDATE_FAILED
	} else {
		tx.Commit()
		code = constant.UPDATE_SUCCESS
	}
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteAppByIds
//@description: 批量删除App记录
//@param: ids request.IdsReq
//@return: err error

func DeleteAppByIds(ids request.IdsReq) (err error) {
	err = global.DB.Delete(&[]model.App{}, "id in ?", ids.Ids).Error
	return err
}

//求交集
func intersect(slice1, slice2 []uint) []uint {
	m := make(map[uint]int)
	nn := make([]uint, 0)
	for _, v := range slice1 {
		m[v]++
	}

	for _, v := range slice2 {
		times, _ := m[v]
		if times == 1 {
			nn = append(nn, v)
		}
	}
	return nn
}

//求差集 slice1-并集
func difference(slice1, slice2 []uint) []uint {
	m := make(map[uint]int)
	nn := make([]uint, 0)
	inter := intersect(slice1, slice2)
	for _, v := range inter {
		m[v]++
	}

	for _, value := range slice1 {
		times, _ := m[value]
		if times == 0 {
			nn = append(nn, value)
		}
	}
	return nn
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateApp
//@description: 更新App记录
//@param: app *model.App
//@return: err error

func UpdateApp(app model.AppDetail, nodes []*model.Node) (code int32, err error) {

	//清理
	var nodeId []uint
	global.DB.Table("app_node").Select("node_id").Where("a_id = ? ", app.Id).Find(&nodeId)
	var nodeList []*model.Node
	global.DB.Where("id  in (?)", nodeId).Find(&nodeList)
	//1.求差异
	clearId := difference(nodeId, app.Nodes)
	if len(clearId) > 0 {
		var uNodes []*model.Node
		//本地清理appNode和namespace 远程删除app和namespace2
		for _, node := range nodeList {
			for _, id := range clearId {
				if node.Id == id {
					uNodes = append(uNodes, node)
				}
			}
		}
		code, err = ClearAppNode(&app.App, clearId, uNodes)
		if code != constant.DELETE_SUCCESS {
			return constant.UPDATE_FAILED, err
		}
	}
	var appt model.App
	global.DB.Where("app_id = ? and name = ?", app.AppId, app.Name).Find(&appt)
	if appt == (model.App{}) {
		//2. 更新已经存在的
		updateIds := intersect(nodeId, app.Nodes)
		if len(updateIds) > 0 {
			var uNode []*model.Node
			for _, node := range nodes {
				for _, id := range updateIds {
					if node.Id == id {
						uNode = append(uNode, node)
					}
				}
			}
			//更新现在已有的
			code, err = UpdateRemoteApp(&app.App, nodes)
		}
	}

	//3.插入不存在的
	insertId := difference(app.Nodes, nodeId)
	if len(insertId) > 0 {
		var nodeins []*model.Node
		global.DB.Where("id in (?)", insertId).Find(&nodeins)
		//远程插入
		code, err = CreateApp(&app.App, nodeins)
		if code != constant.CREATE_SUCCESS {
			return constant.UPDATE_FAILED, err
		}
	}

	return code, err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetApp
//@description: 根据id获取App记录
//@param: id uint
//@return: err error, app model.App

func GetApp(id string) (err error, app model.App) {
	err = global.DB.Where("app_id = ?", id).First(&app).Error
	return
}
func GetAppDetail(id uint) (err error, app model.AppDetail) {
	err = global.DB.Where("id = ?", id).First(&app.App).Error
	an := global.DB.Model(&model.AppNode{})
	an.Where("app_id=?", app.AppId).Select("node_id").Find(&app.Nodes)
	return
}

func GetIf(appNode *request.AppNodeList) (code int32, msg string) {
	var buffer bytes.Buffer
	var app model.App
	global.DB.Where(model.App{AppId: appNode.App.AppId, Name: appNode.App.Name}).First(&app)
	var app1 model.App

	if app != (model.App{}) {
		for _, id := range appNode.Nodes {
			var appNodes model.AppNode
			var node model.Node
			global.DB.Where(model.Node{Id: id}).First(&node)
			global.DB.Where(model.AppNode{AppId: app.AppId, NodeId: id}).Find(&appNodes)
			if appNodes != (model.AppNode{}) {
				buffer.WriteString(fmt.Sprintf("%s ", node.Name))
				code = constant.APP_EXISTS_NODE
			}
		}
	} else {
		global.DB.Where(model.App{AppId: appNode.App.AppId}).First(&app1)
		if app1 != (model.App{}) {
			buffer.WriteString(app1.AppId)
			code = constant.APPID_DUPLICATE
		}
	}
	msg = buffer.String()
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetAppInfoList
//@description: 分页获取App记录
//@param: info request.AppSearch
//@return: err error, list interface{}, total int64

func GetAppInfoList(info request.AppSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.DB.Model(&model.App{})
	var apps []model.App
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&apps).Error
	var name []string

	appList := make([]interface{}, 0)
	for _, app := range apps {
		global.DB.Debug().Model(&model.Node{}).Where("id in (?)", global.DB.Debug().Model(&model.AppNode{}).Select("node_id").Where("a_id=?", app.Id)).Select("name").Find(&name)
		as := model.AppList{
			app,
			name,
		}
		appList = append(appList, as)
	}
	return err, appList, total
}
