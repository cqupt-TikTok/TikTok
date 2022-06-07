package model

import (
	"TikTok/db"
	"fmt"
	"gorm.io/gorm"
)

// Comment 评论
type Comment struct {
	gorm.Model
	VideoId uint   `gorm:"column:video_id;not null"` //视频id作为外键
	UserId  uint   `gorm:"column:user_id;not null"`  //用户id，作为第二外键
	Content string `gorm:"column:content;type:text"` //评论内容
}

// CommentResp 响应结构体
type CommentResp struct {
	Id         uint     `json:"id,omitempty"`          // 评论id
	User       UserResp `json:"user"`                  // 评论用户信息
	Content    string   `json:"content,omitempty"`     // 评论内容
	CreateDate string   `json:"create_date,omitempty"` // 评论发布日期，格式 mm-dd
}

// ToResp 转化为响应结构体
func (C Comment) ToResp(UserId uint) (CR CommentResp) {
	CR.Id = C.ID
	CR.Content = C.Content
	year, month, day := C.CreatedAt.Date()
	CR.CreateDate = fmt.Sprintf("%d%02d%02d", year, month, day)
	var U User
	db.DB.Where("id=?", C.UserId).First(&U)
	if U.ID > 0 {
		CR.User = UserResp{}
		CR.User.IsFollowJudge(UserId)
	} else {
		CR.User = U.ToResp()
	}
	return CR
}
