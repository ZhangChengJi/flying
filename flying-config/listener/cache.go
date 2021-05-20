package listener

import (
	"encoding/json"
	"flying-config/model/request"
	"flying-config/model/response"
	"fmt"
	"github.com/Allenxuxu/toolkit/convert"
	"github.com/coocood/freecache"
	"runtime/debug"
)

var _ Cache = &freeCache{}

type Cache interface {
	Set(config *response.FlyingConfig) error
	Get(config *request.Client) (v *response.FlyingConfig, ok bool)
	Clear()
}

func New(cacheSize int) Cache {
	return newFreeCache(cacheSize)
}

type freeCache struct {
	cache *freecache.Cache
}

func (f freeCache) Set(c *response.FlyingConfig) error {
	value, err := json.Marshal(c)
	if err != nil {
		return err
	}
	return f.cache.Set(getKey(c.AppId), value, -1)
}

func (f freeCache) Get(c *request.Client) (*response.FlyingConfig, bool) {
	v, err := f.cache.Get(getKey(c.AppId))
	if err != nil {
		return nil, false
	}

	var value response.FlyingConfig
	err = json.Unmarshal(v, &value)
	if err != nil {
		return nil, false
	}
	return &value, true
}

func (f freeCache) Del(c *request.Client) bool {
	return f.cache.Del(getKey(c.AppId))
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

func getKey(appId string) []byte {
	return convert.StringToBytes(fmt.Sprintf("%s", appId))
}
