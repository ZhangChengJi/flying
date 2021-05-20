package request

import (
	"flying-admin/model"
)

type AppSearch struct {
	model.App
	PageInfo
}
type AppNodeList struct {
	model.App
	Nodes []uint `json:"nodes" form:"nodes"`
}
