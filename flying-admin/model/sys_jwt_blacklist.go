package model

type JwtBlacklist struct {
	Id  uint   `json:"id" form:"id" gorm:"column:id;comment:id;"`
	Jwt string `gorm:"type:text;comment:jwt"`
}

func (JwtBlacklist) TableName() string {
	return "jwt_blacklist"
}
