package model

/*
	model包 用来封装model-数据库表
*/

type User struct {
	UserId        int64		`json:"id" gorm:"column:user_id"`
	UserName      string	`json:"name" gorm:"column:user_name"`
	Password      string	`gorm:"-"`
	FollowCount   int64		`json:"follow_count" gorm:"column:follow_count"`
	FollowerCount int64		`json:"follower_count" gorm:"column:follower_count"`
	IsFollow	  bool		`json:"is_follow" gorm:"column:is_follow"`
}


func (u User) TableName() string {
	return "user"
}