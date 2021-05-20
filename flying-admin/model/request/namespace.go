package request

import "flying-admin/model"

type NamespaceSearch struct {
	model.Namespace
	PageInfo
}
type AddNamespace struct {
	model.Namespace
	Nodes []uint `json:"nodes" form:"nodes"`
}
type IdNodeId struct {
	Id     int64 `json:"id" form:"id"`
	NodeId uint  `json:"nodeId" form:"nodeId"`
}
type Release struct {
	Id          uint   `json:"id" form:"id"`
	NodeId      uint   `json:"nodeId" form:"nodeId"`
	ReleaseName string `json:"releaseName" form:"releaseName"`
}
