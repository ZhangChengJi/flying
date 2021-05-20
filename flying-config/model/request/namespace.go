package request

import "flying-config/model"

type NamespaceSearch struct {
	model.Namespace
	PageInfo
}
