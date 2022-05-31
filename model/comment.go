package model

type Comment struct {
	Id         int64  `json:"id,omitempty" gorm:"id"`                   // 评论id
	VideoId    int64  `json:"-" gorm:"video_id"`                        //视频id作为外键
	UserId     int64  `json:"-" gorm:"userId"`                          //用户id，作为第二外键，json解析自动忽略
	User       User   `json:"user" gorm:"user"`                         // 评论用户信息
	Content    string `json:"content,omitempty" gorm:"content"`         // 评论内容
	CreateDate string `json:"create_date,omitempty" gorm:"create_date"` // 评论发布日期，格式 mm-dd
}
