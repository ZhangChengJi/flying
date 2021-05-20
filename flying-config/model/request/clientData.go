package request

type Client struct {
	AppId     string `json:"appId" form:"appId" binding:"required"`
	Namespace string `json:"namespace" form:"namespace" binding:"required"`
}
