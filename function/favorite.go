package function

import (
	"TikTok/db"
	"TikTok/model"
	"errors"
	"gorm.io/gorm"
	"time"
)

// AddFavoriteVideo 点赞
func AddFavoriteVideo(videoId, userId uint) error {
	var favoriteVideoRelation = model.FavoriteVideoRelation{
		Id:           0,
		VideoId:      videoId,
		UserId:       userId,
		FavoriteDate: time.Now(),
	}
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
	if err := tx.Model(&model.Video{}).Where("id = ? ", videoId).Update("favorite_count", gorm.Expr("favorite_count+ ?", 1)).Error; err != nil {
		tx.Rollback()
		return err
	}
	var FVR model.FavoriteVideoRelation
	tx.Model(&model.FavoriteVideoRelation{}).Where("video_id = ? and user_id = ?", videoId, userId).First(&FVR)
	if FVR.Id > 0 {
		tx.Rollback()
		return errors.New("重复点赞")
	}
	if err := tx.Create(&favoriteVideoRelation).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error

}

// DropFavoriteVideo 取消点赞
func DropFavoriteVideo(videoId, userId uint) error {
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
	if err := tx.Model(&model.Video{}).Where("id = ? ", videoId).Update("favorite_count", gorm.Expr("favorite_count+ ?", -1)).Error; err != nil {
		tx.Rollback()
		return err
	}
	var FVR model.FavoriteVideoRelation
	if err := tx.Model(&model.FavoriteVideoRelation{}).Where("video_id = ? and user_id = ?", videoId, userId).First(&FVR).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Where("video_id = ? and user_id = ?", videoId, userId).Delete(&model.FavoriteVideoRelation{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error

}

// GetFavoriteVideoList 查询点赞视频列表
func GetFavoriteVideoList(userId uint) (favoriteVideoList []model.VideoResp, err error) {
	var FVR []model.FavoriteVideoRelation
	err = db.DB.Model(&model.FavoriteVideoRelation{}).Where("user_id = ?", userId).Find(&FVR).Error
	if err != nil {
		return nil, err
	}
	var tempVideo model.Video
	for _, v := range FVR {
		db.DB.Where("id = ?", v.VideoId).First(&tempVideo)
		favoriteVideoList = append(favoriteVideoList, tempVideo.ToResp(userId))
	}
	return favoriteVideoList, nil
}
