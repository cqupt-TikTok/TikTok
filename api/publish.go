package api

import (
	"TikTok/apifunc"
	"TikTok/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Publish 视频投稿
func Publish(c *gin.Context) {
	var resp model.BaseResponse
	err := apifunc.Publish(c)
	if err != nil {
		resp.StatusCode = -1
		resp.StatusMsg = "发布失败"
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.StatusCode = 0
	resp.StatusMsg = "发布成功"
	c.JSON(http.StatusOK, resp)
	return
}

// PublishList 发布列表
func PublishList(c *gin.Context) {
	var resp model.PostListResponse
	var err error
	resp, err = apifunc.PublishList(c)
	if err != nil {
		resp.StatusCode = -1
		resp.StatusMsg = "查询失败"
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.StatusCode = 0
	resp.StatusMsg = "查询成功"
	c.JSON(http.StatusOK, resp)
	return
}
