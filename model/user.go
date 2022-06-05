package model

import "time"

type User struct {
	Id            int64     `json:"id,omitempty" gorm:"column:id;primaryKey"`              // 用户id
	Name          string    `json:"name,omitempty" gorm:"column:name"`                     // 用户名称
	Password      string    `json:"-" gorm:"column:password"`                              //用户密码，json编码时自动忽略该字段
	FollowCount   int64     `json:"follow_count,omitempty" gorm:"column:follow_count"`     // 关注总数
	FollowerCount int64     `json:"follower_count,omitempty" gorm:"column:follower_count"` // 粉丝总数
	IsFollow      bool      `json:"is_follow,omitempty" gorm:"-"`                          // true-已关注，false-未关注，gorm解析自动忽略
	CreateDate    time.Time `json:"-" gorm:"column:create_date"`                           //注册时间
}
