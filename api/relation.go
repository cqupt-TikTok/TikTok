// @Title : relation
// @Description :
// @Author : MX
// @Update : 2022/6/6 17:35

package api

import (
	"net/http"

	"TikTok/function"
	"TikTok/model"
	"github.com/gin-gonic/gin"
)

func FollowAction(c *gin.Context) {
	err := function.FollowAction(c)
	var resp model.BaseResponse
	ActionType := c.Param("action_type")
	if ActionType == "1" {
		resp.StatusMsg = "关注"
	} else {
		resp.StatusMsg = "取消关注"
	}

	if err != nil {
		resp.StatusCode = 0
		resp.StatusMsg += "成功"
	} else {
		resp.StatusCode = 0
		resp.StatusMsg += "失败"
	}

	c.JSON(http.StatusOK, resp)
}
