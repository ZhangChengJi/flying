package request

import "flying-config/model"

type ReleaseSearch struct {
	model.Release
	PageInfo
}
type ReleaseDo struct {
	AppId         string `json:"appId" form:"appId"`
	NamespaceName string `json:"namespaceName" form:"namespaceName"`
	ReleaseName   string `json:"releaseName" form:"releaseName"`
}
