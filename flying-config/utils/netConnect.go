package utils

import (
	"flying-config/global"
	"fmt"
	"go.uber.org/zap"
	http "net/http"
	"strings"
	"time"
)

const re = "/test"

var (
	Client http.Client
)

func init() {
	Client = http.Client{Timeout: 3 * time.Second}

}

func Connect(url string) (ok bool) {
	if strings.HasSuffix(url, "/") {
		url = url[:len(url)-1]
	}
	url = fmt.Sprintf("%s%s", url, re)

	rep, err := Client.Get(url)

	if err != nil {
		global.LOG.Error(url+"连接出错", zap.Any(" error:", err))
		ok = false
		return
	} else {
		if rep.StatusCode == http.StatusOK {
			ok = true
		} else {
			ok = false
		}
	}
	return
}
