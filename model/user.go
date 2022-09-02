package model

type UserModel struct {
	UserID        int64  `gorm:"column:user_id;primaryKey;unique;not null;autoIncrement" json:"id"`
	UserName      string `gorm:"column:user_name" json:"user_name"`
	Password      string `gorm:"column:password" json:"-"`
	FollowCount   int64  `gorm:"column:follow_count" json:"follow_count"`
	FollowerCount int64  `gorm:"column:follower_count" json:"follower_count"`
}

func (UserModel) TableName() string {
	return "user"
}
