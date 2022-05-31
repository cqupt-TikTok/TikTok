package model

type Video struct {
	Id            int64  `json:"id,omitempty" gorm:"primaryKey"` // 视频唯一标识
	AuthorId      int64  `json:"-" gorm:"author_id"`             //作者id，作为外键,json解析自动忽略
	Title         string `json:"title" gorm:"title"`
	Author        User   //`json:"author" gorm:"-"`                                // 视频作者信息，gorm解析自动忽略
	PlayUrl       string `json:"play_url,omitempty" gorm:"play_url"`             // 视频播放地址
	CoverUrl      string `json:"cover_url,omitempty" gorm:"cover_url"`           // 视频封面地址
	FavoriteCount int64  `json:"favorite_count,omitempty" gorm:"favorite_count"` // 视频的点赞总数
	CommentCount  int64  `json:"comment_count,omitempty" gorm:"comment_count"`   // 视频的评论总数
	IsFavorite    bool   `json:"is_favorite,omitempty" gorm:"-"`                 // true-已点赞，false-未点赞，gorm解析自动忽略
}
