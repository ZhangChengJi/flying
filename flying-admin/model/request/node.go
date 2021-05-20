package request

import "flying-admin/model"

type NodeSearch struct {
	model.Node
	PageInfo
}

type SyncId struct {
	AppId  string `json:"appId" form:"appId"`
	NodeId uint   `json:"nodeId" form:"nodeId"`
}
