package model

import "comm/database"

type Discussion struct {
	database.Model
	//用户
	UserId uint64
	//文章
	PostId uint64
	//内容
	Content string
	//评论
	DiscussionId uint64
	//热度
	hot int
}

func (Discussion) IsEntity() {}
