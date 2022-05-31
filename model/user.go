package model

type User struct {
	Id            int64  `json:"id,omitempty" gorm:"primaryKey"`                 // 用户id
	Name          string `json:"name,omitempty" gorm:"name"`                     // 用户名称
	Password      string `json:"-" gorm:"password"`                              //用户密码，json编码时自动忽略该字段
	FollowCount   int64  `json:"follow_count,omitempty" gorm:"follow_count"`     // 关注总数
	FollowerCount int64  `json:"follower_count,omitempty" gorm:"follower_count"` // 粉丝总数
	IsFollow      bool   `json:"is_follow,omitempty" gorm:"-"`                   // true-已关注，false-未关注，gorm解析自动忽略
}
