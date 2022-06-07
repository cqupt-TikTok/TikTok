package function

import (
	"TikTok/db"
	"TikTok/model"
	"fmt"
	"gorm.io/gorm"
)

func AddComment(tempComment model.Comment) error {
	tx := db.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}
	var v model.Video
	if err := tx.Model(&model.Video{}).Where("id = ?", tempComment.VideoId).First(&v).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Model(&model.Video{}).Where("id = ? ", tempComment.VideoId).Update("comment_count", gorm.Expr("comment_count+ ?", 1)).Error; err != nil {
		tx.Rollback()
		return err
	}
	fmt.Println(tempComment)
	if err := tx.Create(&tempComment).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func DeleteComment(commentId, userId, videoId uint) error {
	tx := db.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}
	var v model.Video
	if err := tx.Model(&model.Video{}).Where("id = ?", videoId).First(&v).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Model(&model.Video{}).Where("id = ? ", videoId).Update("comment_count", gorm.Expr("comment_count+ ?", -1)).Error; err != nil {
		tx.Rollback()
		return err
	}
	var c model.Comment
	if err := tx.Model(&model.Comment{}).Where("id = ? and video_id = ? and user_id = ?", commentId, videoId, userId).First(&c).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Where("id = ? and video_id = ? and user_id = ?", commentId, videoId, userId).Delete(&model.Comment{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}
