package api

import (
	"TikTok/apifunc"
	"TikTok/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CommentAction 评论操作
func CommentAction(c *gin.Context) {
	var resp model.CommentActionResponse
	var commentResp model.CommentResp
	var err error
	commentResp, err = apifunc.CommentAction(c)
	resp.CommentResp = commentResp
	actionType := c.Query("action_type")
	if actionType == "1" {
		resp.StatusMsg = "评论"
	} else if actionType == "2" {
		resp.StatusMsg = "删除评论"
	} else {
		resp.StatusCode = -1
		resp.StatusMsg = err.Error()
		c.JSON(http.StatusOK, resp.BaseResponse)
		return
	}
	if err != nil {
		resp.StatusCode = -1
		resp.StatusMsg += "失败"
		c.JSON(http.StatusOK, resp.BaseResponse)
		return
	} else if actionType == "2" {
		resp.StatusCode = 0
		resp.StatusMsg += "成功"
		c.JSON(http.StatusOK, resp.BaseResponse)
		return
	} else {
		resp.StatusCode = 0
		resp.StatusMsg += "成功"
		c.JSON(http.StatusOK, resp)
		return
	}
}

// CommentList 评论列表
func CommentList(c *gin.Context) {
	var resp model.CommentListResponse
	var err error
	resp, err = apifunc.CommentList(c)
	if err != nil {
		resp.StatusCode = -1
		resp.StatusMsg = "查询失败:" + err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.StatusCode = 0
	resp.StatusMsg = "查询成功"
	c.JSON(http.StatusOK, resp)
	return
}
