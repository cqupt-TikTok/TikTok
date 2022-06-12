package api

import (
	"TikTok/apifunc"
	"TikTok/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Feed 视频流
func Feed(c *gin.Context) {
	var resp model.FeedResponse
	var err error
	resp, err = apifunc.Feed(c)
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
