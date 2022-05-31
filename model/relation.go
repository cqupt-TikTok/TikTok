package model

import "time"

// Follower 用户关注表
type Follower struct {
	Id         int64     `gorm:"id"`          //主键唯一id
	UserId     int64     `gorm:"user_id"`     //用户id
	FollowerId int64     `gorm:"follower_id"` //关注用户的id
	FollowDate time.Time `gorm:"follow_date"` //关注时间
}

// FavoriteVideo 视频点赞表
type FavoriteVideo struct {
	Id           int64     `gorm:"id"`            //主键唯一id
	VideoId      int64     `gorm:"video_id"`      //视频id
	UserId       int64     `gorm:"user_id"`       //用户id
	FavoriteDate time.Time `gorm:"favorite_date"` //点赞时间
}
