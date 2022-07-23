package model

import "comm/database"

type Post struct {
	database.Model
	//文章标题
	Title string
	//文章论点
	Argument string
	//文章描述
	Describe string
	//文章热度
	Heat uint64
	//是否过期
	Enable bool
	//用户
	UserId uint64 `json:"user_id" gorm:"index"`
}

func (Post) IsEntity() {}
