package model

type User struct {
	Id        uint   `json:"id" form:"id" gorm:"column:id;comment:id;"`
	Username  string `json:"userName" gorm:"comment:用户登录名"`
	Password  string `json:"-"  gorm:"comment:用户登录密码"`
	NickName  string `json:"nickName" gorm:"default:系统用户;comment:用户昵称" `
	HeaderImg string `json:"headerImg" gorm:"default:http://qmplusimg.henrongyi.top/head.png;comment:用户头像"`
}
type SysUserResponse struct {
	User User `json:"user"`
}

type LoginResponse struct {
	User      User   `json:"user"`
	Token     string `json:"token"`
	ExpiresAt int64  `json:"expiresAt"`
}

func (User) TableName() string {
	return "user"
}
