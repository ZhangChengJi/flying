package logic

import (
	"crypto/md5"
	"encoding/hex"
	"flying-admin/global"
	"flying-admin/model"
)

func Login(u *model.User) (err error, userInter *model.User) {
	var user model.User
	u.Password = MD5V([]byte(u.Password))
	err = global.DB.Where("username = ? AND password = ?", u.Username, u.Password).First(&user).Error
	return err, &user
}
func MD5V(str []byte) string {
	h := md5.New()
	h.Write(str)
	return hex.EncodeToString(h.Sum(nil))
}

func FindUserById(id uint) (err error, user model.User) {
	err = global.DB.Debug().Where("id=?", id).First(&user).Error
	return
}
func UpPass(u *model.User) (err error) {
	err = global.DB.Save(u).Error
	return
}
