package utils

import (
	"encoding/json"
	"flying-admin/model"
	"fmt"
	"github.com/Allenxuxu/toolkit/convert"
	"github.com/coocood/freecache"
	"runtime/debug"
	"sync"
)

var _ Cache = &freeCache{}
var y sync.RWMutex

type Cache interface {
	Set(config *model.Node) (err error)
	Get(config *string) (v *model.Node, ok bool)
	Clear()
}

func NewCache(cacheSize int) Cache {
	return newFreeCache(cacheSize)
}

type freeCache struct {
	cache *freecache.Cache
}

func (f freeCache) Set(c *model.Node) (err error) {
	y.Lock()
	value, err := json.Marshal(c)
	if err != nil {
		return err
	}
	err = f.cache.Set(getKey(&c.Key), value, -1)
	y.Unlock()
	return
}

func (f freeCache) Get(key *string) (*model.Node, bool) {
	v, err := f.cache.Get(getKey(key))
	if err != nil {
		return nil, false
	}

	var value model.Node
	err = json.Unmarshal(v, &value)
	if err != nil {
		return nil, false
	}
	return &value, true
}

func (f freeCache) Del(c *model.Node) bool {
	return f.cache.Del(getKey(&c.Key))
}

func (f freeCache) Clear() {
	f.cache.Clear()
}

func newFreeCache(size int) Cache {
	c := &freeCache{
		cache: freecache.NewCache(size),
	}
	debug.SetGCPercent(10)

	return c
}

func getKey(appId *string) []byte {
	return convert.StringToBytes(fmt.Sprintf("node:%s", appId))
}
