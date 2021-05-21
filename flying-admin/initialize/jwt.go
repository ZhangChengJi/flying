package initialize

import (
	"flying-admin/global"
	"flying-admin/middleware"
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"log"
)

func Jwt() {
	publicKeyByte, err := ioutil.ReadFile(global.GVA_CONFIG.JWT.PublicKey)
	if err != nil {
		log.Println(err.Error())
	}
	middleware.PublicKey, err = jwt.ParseRSAPublicKeyFromPEM(publicKeyByte)
	privateKeyByte, err := ioutil.ReadFile(global.GVA_CONFIG.JWT.PrivateKey)
	if err != nil {
		log.Println(err.Error())
	}
	middleware.PrivateKey, _ = jwt.ParseRSAPrivateKeyFromPEM(privateKeyByte)
}
