package listener

import (
	"container/list"
	"fmt"
	"sync"
)

type Listener struct {
	s        sync.RWMutex
	cache    Cache
	watchers map[string]*list.List
}

var defaultListener *Listener

func Init(cacheSize int) {
	defaultListener = newConfig(cacheSize)
}
func newConfig(cacheSize int) *Listener {
	fmt.Println("初始化")
	return &Listener{
		cache:    New(cacheSize),
		watchers: make(map[string]*list.List),
	}
}

var i = 1

func AddListener(key string) {

	list := defaultListener.copyWatchers(key)
	for e := list.Front(); e != nil; e = e.Next() {
		w := e.Value.(*Qt)
		select {
		case w.updates <- &i:
		default:
		}
	}
}

func (l *Listener) copyWatchers(key string) *list.List {
	watcherList := list.New()
	l.s.RLock()
	watchers := l.watchers[key]
	if watchers != nil {
		watcherList.PushBackList(watchers)
	}
	l.s.RUnlock()

	return watcherList
}

func Watch(appId, namespace string) *Qt {
	return defaultListener.Watch(appId, namespace)
}

func (l *Listener) Watch(appId, namespace string) *Qt {
	w := &Qt{
		exit:    make(chan interface{}),
		updates: make(chan *int, 1),
	}
	l.s.Lock()
	watchers := l.watchers[appId]
	if watchers == nil {
		watchers = list.New()
		l.watchers[appId] = watchers
	}
	e := watchers.PushBack(w)
	l.s.Unlock()
	go func() {
		<-w.exit
		l.s.Lock()
		watchers.Remove(e)
		l.s.Unlock()
	}()
	return w
}
