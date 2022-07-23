package model

type User struct {
	Model
	//名称
	Name string `json:"name"`
	//电话号码
	Phone string `json:"phone" gorm:"index:,unique"`
	//密码
	Password string `json:"password"`
	//头像
	Avatar string `json:"avatar"`
}

func (User) IsEntity() {}
