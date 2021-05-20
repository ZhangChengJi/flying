package utils

import "log"

const (
	Add    = "add"
	Update = "update"
	Del    = "delete"
)

var (
	Dispatcher *EventDispatcher
)

//事件基类
type Event struct {
	IDispatcher iEventDispatcher //事件响应
	EventType   string           //类型
	Object      interface{}      //数据
}

//调度类
type EventDispatcher struct {
	cells []*eventCell
}

//调试器里存放单元
type eventCell struct {
	eventType string
	listeners []*EventListener
}

//监听
type EventListener struct {
	Haldler EventHandler
}

// 监听器函数
type EventHandler func(event Event)

//事件调试接口
type iEventDispatcher interface {
	//注册事件
	AddEventListener(eventType string, listener *EventListener)
	//移除事件
	RemoveEventListener(eventType string, listener *EventListener) bool
	//是否包含事件
	HasEventListener(eventType string) bool
	//事件派发
	DispatchEvent(e string, o interface{}) bool
	//删除所有事件
	RemoveAllEvent()
}

//调试器
func NewEventDispatcher() *EventDispatcher {
	e := &EventDispatcher{}
	return e
}

//监听器
func NewEventListener(h EventHandler) *EventListener {
	l := &EventListener{h}
	return l
}

//创建事件
func NewEvent(eventType string, obj interface{}) Event {
	e := Event{EventType: eventType, Object: obj}
	return e
}

//克隆事件
func (e *Event) Clone() *Event {
	n := &Event{}
	n.EventType = e.EventType
	n.Object = e.Object
	return n
}

//显示
func (e *Event) ToString() string {
	return ("Event eventType:" + e.EventType)
}

//添加事件
func (e *EventDispatcher) AddEventListener(eventType string, listener *EventListener) {
	// event := Event{name,data}
	//重复添加，待定
	for _, ev := range e.cells {
		if ev.eventType == eventType {
			ev.listeners = append(ev.listeners, listener)
			return
		}
	}
	//空的集合，要添加
	cell := &eventCell{eventType: eventType, listeners: []*EventListener{listener}}
	e.cells = append(e.cells, cell)
}

//取消
func (e *EventDispatcher) RemoveEventListener(eventType string, listener *EventListener) bool {
	// event := Event{name,data}
	//重复添加，待定
	for _, ev := range e.cells {
		if ev.eventType == eventType {
			for i, subListener := range ev.listeners {
				if subListener == listener {
					//找到了
					ev.listeners = append(ev.listeners[:i], ev.listeners[i+1:]...)
					return true
				}
			}
		}
	}
	return false
}

//是否包含某个事件
func (e *EventDispatcher) HasEventListener(eventType string) bool {
	for _, cell := range e.cells {
		if cell.eventType == eventType {
			return true
		}
	}
	return false
}

//事件分发
func (e *EventDispatcher) DispatchEvent(eventType string, o interface{}) bool {
	event := Event{EventType: eventType, Object: o}
	for _, cell := range e.cells {
		if cell.eventType == event.EventType {
			for _, listener := range cell.listeners {
				event.IDispatcher = e
				listener.Haldler(event)
			}
			return true
		}
	}
	log.Println("DispatchEvent error event:", event.EventType)
	return false
}

//删除所有事件
func (e *EventDispatcher) RemoveAllEvent() {
	e.cells = e.cells[0:0]
}
