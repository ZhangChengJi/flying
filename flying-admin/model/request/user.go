package request

type User struct {
	Id        uint   `json:"id" form:"id" gorm:"column:id;comment:id;"`
	Username  string `json:"userName" gorm:"comment:用户登录名"`
	Password  string `json:"-"  gorm:"comment:用户登录密码"`
	NickName  string `json:"nickName" gorm:"default:系统用户;comment:用户昵称" `
	HeaderImg string `json:"headerImg" gorm:"default:http://qmplusimg.henrongyi.top/head.png;comment:用户头像"`
}

// User login structure
type Login struct {
	Username string `json:"username"  form:"username"`
	Password string `json:"password" form:"password"`
}
type UpPass struct {
	Username string `json:"username"  form:"username"`
	Password string `json:"password" form:"password"`
	Pass     string `json:"pass"`
}
