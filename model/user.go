package model

/*
	model包 用来封装model-数据库表
*/

type User struct {
	UserId        int64		`gorm:"column:user_id"`
	UserName      string	`gorm:"column:user_name"`
	Password      string	`gorm:"-"`
	FollowCount   int64		`gorm:"column:follow_count"`
	FollowerCount int64		`gorm:"column:follower_count"`
	IsFollow	  bool		`gorm:"column:is_follow"`
}


func (u User) TableName() string {
	return "user"
}