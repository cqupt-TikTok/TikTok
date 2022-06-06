package model

import (
	"TikTok/db"
	"gorm.io/gorm"
)

// Video 视频
type Video struct {
	gorm.Model
	AuthorId      uint   `gorm:"column:author_id"`      //作者id，作为外键
	Title         string `gorm:"column:title"`          //视频标题
	PlayUrl       string `gorm:"column:play_url"`       // 视频播放地址
	CoverUrl      string `gorm:"column:cover_url"`      // 视频封面地址
	FavoriteCount int64  `gorm:"column:favorite_count"` // 视频的点赞总数
	CommentCount  int64  `gorm:"column:comment_count"`  // 视频的评论总数
}

// VideoResp 响应结构体
type VideoResp struct {
	Id            uint     `json:"id,omitempty"`             // 视频唯一标识
	Title         string   `json:"title"`                    //视频标题
	Author        UserResp `json:"author"`                   // 视频作者信息
	PlayUrl       string   `json:"play_url,omitempty"`       // 视频播放地址
	CoverUrl      string   `json:"cover_url,omitempty"`      // 视频封面地址
	FavoriteCount int64    `json:"favorite_count,omitempty"` // 视频的点赞总数
	CommentCount  int64    `json:"comment_count,omitempty"`  // 视频的评论总数
	IsFavorite    bool     `json:"is_favorite,omitempty"`    // true-已点赞，false-未点赞
}

// ToResp 转化为响应结构体，默认点赞
func (V Video) ToResp(UserId uint) (VR VideoResp) {
	VR.Id = V.ID
	VR.Title = V.Title
	VR.PlayUrl = V.PlayUrl
	VR.CoverUrl = V.CoverUrl
	VR.FavoriteCount = V.FavoriteCount
	VR.CommentCount = V.CommentCount
	var u User
	db.DB.Where("id = ?", V.AuthorId).First(&u)
	if u.ID > 0 {
		VR.Author = u.ToResp()
		VR.Author.IsFollowJudge(UserId)
	} else {
		VR.Author = UserResp{}
	}
	VR.IsFavorite = true
	return VR
}

// IsFavoriteJudge 点赞校验，视情况调用
func (VR *VideoResp) IsFavoriteJudge(UserId uint) {
	var FVR FavoriteVideoRelation
	db.DB.Where("video_id = ? AND user_id = ?", VR.Id, UserId).First(&FVR)
	if FVR.Id <= 0 {
		(*VR).IsFavorite = false
	}
}
