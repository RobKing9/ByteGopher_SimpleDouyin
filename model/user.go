package model

/*
	model包 用来封装model-数据库表
*/

// User 用户结构体
type User struct {
	UserId        uint
	UserName      string
	Password      string
	FollowCount   int64
	FollowerCount int64
}
