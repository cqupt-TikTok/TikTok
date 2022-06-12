// @Title : relation
// @Description :
// @Author : MX
// @Update : 2022/6/6 17:35

package api

import (
	"TikTok/apifunc"
	"TikTok/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FollowAction(c *gin.Context) {
	var resp model.BaseResponse
	err := apifunc.FollowAction(c)
	ActionType := c.Query("action_type")
	if ActionType == "1" {
		resp.StatusMsg = "关注"
	} else {
		resp.StatusMsg = "取消关注"
	}

	if err != nil {
		resp.StatusCode = -1
		resp.StatusMsg += "失败"
	} else {
		resp.StatusCode = 0
		resp.StatusMsg += "成功"
	}

	c.JSON(http.StatusOK, resp)
}

func FollowList(c *gin.Context) {
	var resp model.FollowListResponse
	var err error
	resp.UserList, err = apifunc.FollowList(c)
	if err != nil {
		resp.StatusCode = -1
		resp.StatusMsg = "获取关注列表失败"
	} else {
		resp.StatusCode = 0
		resp.StatusMsg = "获取关注列表成功"
	}

	c.JSON(http.StatusOK, resp)
}

func FollowerList(c *gin.Context) {
	var resp model.FollowListResponse
	var err error
	resp.UserList, err = apifunc.FollowerList(c)
	if err != nil {
		resp.StatusCode = -1
		resp.StatusMsg = "获取粉丝列表失败"
	} else {
		resp.StatusCode = 0
		resp.StatusMsg = "获取粉丝列表成功"
	}

	c.JSON(http.StatusOK, resp)
}
