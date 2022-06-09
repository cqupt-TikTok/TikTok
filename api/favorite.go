package api

import (
	"TikTok/function"
	"TikTok/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

// FavoriteAction 点赞操作
func FavoriteAction(c *gin.Context) {
	var resp model.BaseResponse
	var err error
	actionType := c.Query("action_type")
	err = function.FavoriteAction(c)
	if actionType == "1" {
		resp.StatusMsg = "点赞"
	} else if actionType == "2" {
		resp.StatusMsg = "取消点赞"
	} else {
		resp.StatusMsg = err.Error()
		resp.StatusCode = -1
		c.JSON(http.StatusOK, resp)
		return
	}
	if err != nil {
		resp.StatusCode = -1
		resp.StatusMsg += "失败"
	} else {
		resp.StatusCode = 0
		resp.StatusMsg += "成功"
	}
	c.JSON(http.StatusOK, resp)
	return
}

// FavoriteVideoList 点赞视频列表
func FavoriteVideoList(c *gin.Context) {
	favoriteListResponse, err := function.FavoriteVideoList(c)
	if err != nil {
		favoriteListResponse.StatusCode = -1
		favoriteListResponse.StatusMsg = err.Error()
		c.JSON(http.StatusOK, favoriteListResponse)
	}
	favoriteListResponse.StatusCode = 0
	favoriteListResponse.StatusMsg = "查询成功"
	c.JSON(http.StatusOK, favoriteListResponse)
	return
}
