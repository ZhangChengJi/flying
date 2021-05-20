package initialize

import (
	"flying-admin/utils"
)

func NewEvent() {
	utils.Dispatcher = utils.NewEventDispatcher()
	add := utils.NewEventListener(utils.AddNode)
	del := utils.NewEventListener(utils.DelNode)
	utils.Dispatcher.AddEventListener(utils.Add, add)
	utils.Dispatcher.AddEventListener(utils.Del, del)

}
