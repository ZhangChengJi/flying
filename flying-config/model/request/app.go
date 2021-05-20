package request

import "flying-config/model"

type AppSearch struct {
	model.App
	PageInfo
}
