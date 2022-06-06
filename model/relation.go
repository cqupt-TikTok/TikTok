package model

import "time"

// FollowRelation 用户关注表
type FollowRelation struct {
	Id         int64     `gorm:"column:id;primaryKey"` //主键唯一id
	FollowerId int64     `gorm:"column:follower_id"`   //被关注用户的id
	UserId     int64     `gorm:"column:user_id"`       //用户id
	FollowDate time.Time `gorm:"column:follow_date"`   //关注时间
}

// FavoriteVideoRelation 视频点赞表
type FavoriteVideoRelation struct {
	Id           int64     `gorm:"column:id;primaryKey"` //主键唯一id
	VideoId      int64     `gorm:"column:video_id"`      //视频id
	UserId       int64     `gorm:"column:user_id"`       //点赞用户id
	FavoriteDate time.Time `gorm:"column:favorite_date"` //点赞时间
}
