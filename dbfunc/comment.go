package dbfunc

import (
	"TikTok/model"
	"TikTok/storage"
	"gorm.io/gorm"
)

// AddComment 添加评论
func AddComment(tempComment model.Comment) error {
	//开启事务
	tx := storage.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}
	//查询评论视频是否存在
	var v model.Video
	if err := tx.Model(&model.Video{}).Where("id = ?", tempComment.VideoId).First(&v).Error; err != nil {
		tx.Rollback()
		return err
	}
	//视频评论数comment_count+1
	if err := tx.Model(&model.Video{}).Where("id = ? ", tempComment.VideoId).Update("comment_count", gorm.Expr("comment_count+ ?", 1)).Error; err != nil {
		tx.Rollback()
		return err
	}
	//在评论表中添加评论
	if err := tx.Create(&tempComment).Error; err != nil {
		tx.Rollback()
		return err
	}
	//提交事务
	return tx.Commit().Error
}

// DeleteComment 删除评论
func DeleteComment(commentId, userId, videoId uint) error {
	//开始事务
	tx := storage.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}
	//查询评论视频是否存在
	var v model.Video
	if err := tx.Model(&model.Video{}).Where("id = ?", videoId).First(&v).Error; err != nil {
		tx.Rollback()
		return err
	}
	//视频评论数comment_count-1
	if err := tx.Model(&model.Video{}).Where("id = ? ", videoId).Update("comment_count", gorm.Expr("comment_count- ?", 1)).Error; err != nil {
		tx.Rollback()
		return err
	}
	//查询删除评论是否存在
	var c model.Comment
	if err := tx.Model(&model.Comment{}).Where("id = ? and video_id = ? and user_id = ?", commentId, videoId, userId).First(&c).Error; err != nil {
		tx.Rollback()
		return err
	}
	//删除评论
	if err := tx.Where("id = ? and video_id = ? and user_id = ?", commentId, videoId, userId).Delete(&model.Comment{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	//提交事务
	return tx.Commit().Error
}

// CommentList 查询评论列表
func CommentList(videoId, userId uint) (commentRespList []model.CommentResp, err error) {
	var commentList []model.Comment
	err = storage.DB.Model(&model.Comment{}).Where("video_id = ?", videoId).Find(&commentList).Error
	if err != nil {
		return nil, err
	}
	//转换为对应响应结构体
	for _, c := range commentList {
		commentRespList = append(commentRespList, c.ToResp(userId))
	}
	return commentRespList, nil
}
