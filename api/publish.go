package api

import (
	"TikTok/apifunc"
	"TikTok/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

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
