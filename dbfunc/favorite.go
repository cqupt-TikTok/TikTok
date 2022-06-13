package dbfunc

import (
	"TikTok/model"
	"TikTok/storage"
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
	//查询点赞视频是否存在
	var v model.Video
	if err := tx.Model(&model.Video{}).Where("id = ?", videoId).First(&v).Error; err != nil {
		tx.Rollback()
		return err
	}
	//查询是否已经点赞
	var FVR model.FavoriteVideoRelation
	tx.Model(&model.FavoriteVideoRelation{}).Where("video_id = ? and user_id = ?", videoId, userId).First(&FVR)
	if FVR.Id > 0 {
		tx.Rollback()
		return errors.New("重复点赞")
	}
	//视频点赞总数favorite_count+1
	if err := tx.Model(&model.Video{}).Where("id = ? ", videoId).Update("favorite_count", gorm.Expr("favorite_count+ ?", 1)).Error; err != nil {
		tx.Rollback()
		return err
	}
	//用户被喜欢总数(获赞总数)TotalFavorited+1
	if err := tx.Model(&model.User{}).Where("id = ? ", v.AuthorId).Update("total_favorited", gorm.Expr("total_favorited+ ?", 1)).Error; err != nil {
		tx.Rollback()
		return err
	}
	//用户喜欢总数(点赞总数)FavoriteCount+1
	if err := tx.Model(&model.User{}).Where("id = ? ", userId).Update("favorite_count", gorm.Expr("favorite_count+ ?", 1)).Error; err != nil {
		tx.Rollback()
		return err
	}
	//点赞表中写入数据
	if err := tx.Create(&favoriteVideoRelation).Error; err != nil {
		tx.Rollback()
		return err
	}
	//提交事务
	return tx.Commit().Error

}

// DropFavoriteVideo 取消点赞
func DropFavoriteVideo(videoId, userId uint) error {
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
	//查询点赞视频是否存在
	var v model.Video
	if err := tx.Model(&model.Video{}).Where("id = ?", videoId).First(&v).Error; err != nil {
		tx.Rollback()
		return err
	}
	//检查是否已经点赞
	var FVR model.FavoriteVideoRelation
	if err := tx.Model(&model.FavoriteVideoRelation{}).Where("video_id = ? and user_id = ?", videoId, userId).First(&FVR).Error; err != nil {
		tx.Rollback()
		return err
	}
	//视频点赞总数favorite_count-1
	if err := tx.Model(&model.Video{}).Where("id = ? ", videoId).Update("favorite_count", gorm.Expr("favorite_count- ?", 1)).Error; err != nil {
		tx.Rollback()
		return err
	}
	//用户被喜欢总数(获赞总数)TotalFavorited-1
	if err := tx.Model(&model.User{}).Where("id = ? ", v.AuthorId).Update("total_favorited", gorm.Expr("total_favorited- ?", 1)).Error; err != nil {
		tx.Rollback()
		return err
	}
	//用户喜欢总数(点赞总数)FavoriteCount-1
	if err := tx.Model(&model.User{}).Where("id = ? ", userId).Update("favorite_count", gorm.Expr("favorite_count- ?", 1)).Error; err != nil {
		tx.Rollback()
		return err
	}
	//删除数据，取消点赞
	if err := tx.Where("video_id = ? and user_id = ?", videoId, userId).Delete(&model.FavoriteVideoRelation{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	//提交事务
	return tx.Commit().Error

}

// GetFavoriteVideoList 查询点赞视频列表
func GetFavoriteVideoList(userId uint) (favoriteVideoList []model.VideoResp, err error) {
	var FVR []model.FavoriteVideoRelation
	err = storage.DB.Where("user_id = ?", userId).Find(&FVR).Error
	if err != nil {
		return nil, err
	}
	for _, v := range FVR {
		var tempVideo model.Video
		storage.DB.Where("id = ?", v.VideoId).First(&tempVideo)
		favoriteVideoList = append(favoriteVideoList, tempVideo.ToResp(userId))
	}
	return favoriteVideoList, nil
}
