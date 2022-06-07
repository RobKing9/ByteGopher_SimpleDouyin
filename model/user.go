package model

import (
	"database/sql"
)

type UserModel struct {
	UserID        int            `gorm:"column:user_id;primaryKey;unique;not null;autoIncrement" json:"user_id"`
	UserName      sql.NullString `gorm:"column:user_name" json:"user_name"`
	Password      sql.NullString `gorm:"column:password" json:"password"`
	FollowCount   sql.NullInt64  `gorm:"column:follow_count" json:"follow_count"`
	FollowerCount sql.NullInt64  `gorm:"column:follower_count" json:"follower_count"`
}

