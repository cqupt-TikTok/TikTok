package model

import (
	"TikTok/storage"
	"gorm.io/gorm"
)

// User 用户
type User struct {
	gorm.Model
	Name            string `gorm:"column:name;type:varchar(20);not null"`     // 用户名称
	Password        string `gorm:"column:password;type:varchar(20);not null"` //用户密码
	Signature       string `gorm:"column:signature;type:varchar(20)"`         //个性签名
	Avatar          string `gorm:"column:avatar;type:varchar(100)"`           //用户头像链接
	BackgroundImage string `gorm:"column:background_image;type:varchar(100)"` //背景图链接
	FollowCount     int64  `gorm:"column:follow_count;type:int;default:0"`    // 关注总数
	FollowerCount   int64  `gorm:"column:follower_count;type:int;default:0"`  // 粉丝总数
	TotalFavorited  int64  `gorm:"column:total_favorited;type:int;default:0"` //被赞总次数
	FavoriteCount   int64  `gorm:"column:favorite_count;type:int;default:0"`  //喜欢总数量
}

// UserResp 响应结构体
type UserResp struct {
	Id              uint   `json:"id"`               // 用户id
	Name            string `json:"name"`             // 用户名称
	Signature       string `json:"signature"`        //个性签名
	Avatar          string `json:"avatar"`           //用户头像链接
	BackgroundImage string `json:"background_image"` //背景图链接
	FollowCount     int64  `json:"follow_count"`     // 关注总数
	FollowerCount   int64  `json:"follower_count"`   // 粉丝总数
	TotalFavorited  int64  `json:"total_favorited"`  //被赞总次数
	FavoriteCount   int64  `json:"favorite_count"`   //喜欢总数量
	IsFollow        bool   `json:"is_follow"`        // true-已关注，false-未关注
}

// ToResp 转化为响应结构体，默认关注
func (U User) ToResp() (UR UserResp) {
	UR.Id = U.ID
	UR.Name = U.Name
	UR.FollowCount = U.FollowerCount
	UR.FollowerCount = U.FollowerCount
	UR.TotalFavorited = U.TotalFavorited
	UR.FavoriteCount = U.FavoriteCount
	UR.Signature = U.Signature
	UR.BackgroundImage = U.BackgroundImage
	UR.Avatar = U.Avatar
	UR.IsFollow = true
	return UR
}

// IsFollowJudge 关注校验，视情况调用
func (UR *UserResp) IsFollowJudge(UserId uint) {
	var FR FollowRelation
	storage.DB.Where("follower_id = ? AND user_id = ?", UR.Id, UserId).First(&FR)
	if FR.Id <= 0 {
		(*UR).IsFollow = false
	}
}
