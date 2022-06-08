package api

import (
	"TikTok/function"
	"TikTok/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// FavoriteAction 点赞操作
func FavoriteAction(c *gin.Context) {
	userId := uint(1)
	videoId64, _ := strconv.ParseUint(c.Query("video_id"), 10, 64)
	actionType, _ := strconv.Atoi(c.Query("action_type"))
	videoId := uint(videoId64)
	if actionType == 1 {
		err := function.AddFavoriteVideo(videoId, userId)
		if err != nil {
			c.JSON(http.StatusOK, model.BaseResponse{
				StatusCode: 1,
				StatusMsg:  "false",
			})
			return
		}
		c.JSON(http.StatusOK, model.BaseResponse{
			StatusCode: 0,
			StatusMsg:  "success",
		})
		return
	} else if actionType == 2 {
		err := function.DropFavoriteVideo(videoId, userId)
		if err != nil {
			c.JSON(http.StatusOK, model.BaseResponse{
				StatusCode: 1,
				StatusMsg:  "false",
			})
			return
		}
		c.JSON(http.StatusOK, model.BaseResponse{
			StatusCode: 0,
			StatusMsg:  "success",
		})
		return
	}
	c.JSON(http.StatusOK, model.BaseResponse{
		StatusCode: 1,
		StatusMsg:  "false",
	})
	return
}

//
func FavoriteVideoList(c *gin.Context) {
	userId64, _ := strconv.ParseUint(c.Query("user_id"), 10, 64)
	userId := uint(userId64)
	videoList, err := function.GetFavoriteVideoList(userId)
	if err != nil {
		c.JSON(http.StatusOK, model.FavoriteListResponse{
			BaseResponse: model.BaseResponse{
				StatusCode: 1,
				StatusMsg:  "false",
			},
			VideoList: videoList,
		})
		return
	}
	c.JSON(http.StatusOK, model.FavoriteListResponse{
		BaseResponse: model.BaseResponse{
			StatusCode: 0,
			StatusMsg:  "success",
		},
		VideoList: videoList,
	})
	return
}
