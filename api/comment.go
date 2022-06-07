package api

import (
	"TikTok/function"
	"TikTok/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

// CommentAction 评论操作
func CommentAction(c *gin.Context) {
	userId := uint(1)
	videoId64, _ := strconv.ParseUint(c.Query("video_id"), 10, 64)
	actionType, _ := strconv.Atoi(c.Query("action_type"))
	videoId := uint(videoId64)
	if actionType == 1 {
		commentText := c.Query("comment_text")
		tempComment := model.Comment{
			VideoId: videoId,
			UserId:  userId,
			Content: commentText,
		}
		err := function.AddComment(tempComment)
		if err != nil {
			c.JSON(http.StatusOK, model.BaseResponse{
				StatusCode: 1,
				StatusMsg:  "false",
			})
			return
		}
		tempComment.CreatedAt = time.Now()
		c.JSON(http.StatusOK, model.CommentActionResponse{
			BaseResponse: model.BaseResponse{
				StatusCode: 0,
				StatusMsg:  "success",
			},
			CommentResp: tempComment.ToResp(userId),
		})
		return
	} else if actionType == 2 {
		commentId64, _ := strconv.ParseUint(c.Query("video_id"), 10, 64)
		commentId := uint(commentId64)
		err := function.DeleteComment(commentId, userId, videoId)
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

// CommentList 评论列表
func CommentList(c *gin.Context) {

}
