package function

import (
	"TikTok/dbfunc"
	"TikTok/model"
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
)

// FavoriteAction 点赞操作
func FavoriteAction(c *gin.Context) (err error) {
	userId := uint(1)

	videoId64, _ := strconv.ParseUint(c.Query("video_id"), 10, 64)
	actionType := c.Query("action_type")
	videoId := uint(videoId64)
	if actionType == "1" {
		err = dbfunc.AddFavoriteVideo(videoId, userId)
		if err != nil {
			return err
		}
		return nil
	} else if actionType == "2" {
		err = dbfunc.DropFavoriteVideo(videoId, userId)
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("action_type错误")
}

//
func FavoriteVideoList(c *gin.Context) (model.FavoriteListResponse, error) {
	var favoriteListResponse model.FavoriteListResponse
	userId64, _ := strconv.ParseUint(c.Query("user_id"), 10, 64)
	userId := uint(userId64)
	videoList, err := dbfunc.GetFavoriteVideoList(userId)
	favoriteListResponse.VideoList = videoList
	if err != nil {
		return favoriteListResponse, err
	}
	return favoriteListResponse, nil
}
