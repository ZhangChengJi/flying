// 自动生成模板App
package model

import (
	"time"
)

// 如果含有time.Time 请自行import time包
type App struct {
	Id         uint       `json:"id" form:"id" gorm:"column:id;comment:id;"`
	Name       string     `json:"name" form:"name" gorm:"column:name;comment:应用名称;type:varchar(255);size:255;"`
	AppId      string     `json:"appId" form:"appId" gorm:"column:app_id;comment:appId;type:varchar(255);size:255;"`
	CreatedAt  *time.Time `json:"createdAt" form:"createdAt" gorm:"column:created_at;comment:创建时间;default:null;"`
	UpdateTime *time.Time `json:"updateTime" form:"updateTime" gorm:"column:update_time;comment:更新时间;default:null;"`
}

func (App) TableName() string {
	return "app"
}
