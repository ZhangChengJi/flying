package logic

import (
	"flying-config/global"
	"flying-config/model"
	proto "flying-config/proto/app"
	re "flying-config/proto/common"
)

func CreateApp(req *proto.App) (err error) {
	err = global.DB.Create(&model.App{AppId: req.AppId, Name: req.Name}).Error
	return err
}
func FindAppIds(appId string) (apps []model.App, err error) {
	err = global.DB.Where("app_id = ? ", appId).Find(&apps).Error
	return
}
func DeleteApp(app *model.App) (err error) {
	tx := global.DB.Begin()
	err = tx.Delete(&app).Error
	err = tx.Where("a_id=?", app.Id).Delete(&model.Namespace{}).Error
	err = tx.Where("app_id=?", app.AppId).Delete(&model.Release{}).Error
	if err != nil {
		tx.Rollback()
	} else {
		tx.Commit()
	}
	return
}

func DeleteByIdApp(req *proto.App) (*proto.App, error) {
	err := global.DB.First(&req).Delete(&req).Error
	return req, err
}

func UpdateApp(app *model.App) (err error) {
	err = global.DB.Save(&app).Error
	return
}
func GetById(req *proto.App) (*proto.App, error) {
	var aa proto.App
	err := global.DB.Debug().Where("name=?", "86").First(&aa).Error
	return &aa, err
}
func FindAppId(appId string) (app *model.App, err error) {
	err = global.DB.Debug().Where("app_id = ?", appId).First(&app).Error
	return
}

func QueryApp() (apps []*model.App, err error) {
	err = global.DB.Find(&apps).Error
	return
}

func GetListApp(req *re.Request) (result []*proto.App, total int64, err error) {

	limit := req.PageSize
	offset := req.PageSize * (req.Page - 1)
	// 创建db
	db := global.DB.Model(&model.App{})
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&result).Error
	return

}
