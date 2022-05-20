package model

/*
	model包 用来封装model-数据库表
*/

type User struct {
	UserId        int64
	UserName      string
	Password      string
	FollowCount   int64
	FollowerCount int64
}
