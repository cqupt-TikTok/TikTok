package function

import (
	"TikTok/db"
	"TikTok/model"
	"gorm.io/gorm"
)

func AddComment(tempComment model.Comment) error {
	if err := db.DB.Model(&model.Video{}).Where("id = ? ", tempComment.VideoId).Update("comment_count", gorm.Expr("comment_count+ ?", 1)).Error; err != nil {
		return err
	}
	if err := db.DB.Create(&tempComment).Error; err != nil {
		return err
	}
	return nil
}

func DeleteComment(commentId, userId, videoId uint) error {

	return nil
}
