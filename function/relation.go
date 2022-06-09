// @Title : relation
// @Description :
// @Author : MX
// @Update : 2022/6/6 17:36

package function

import (
	"errors"
	"strconv"

	"TikTok/db"
	"TikTok/model"
	"github.com/gin-gonic/gin"
)

func FollowAction(c *gin.Context) (err error) {
	ToUserIdStr := c.Param("to_user_id")
	ToUserIdInt, err := strconv.Atoi(ToUserIdStr)
	if err != nil {
		return errors.New("请求参数不规范")
	}
	ToUserId := uint(ToUserIdInt)

	ActionType := c.Param("action_type")

	UserIdValue, _ := c.Get("uid")
	UserId := UserIdValue.(uint)

	relation := model.FollowRelation{
		FollowerId: ToUserId,
		UserId:     UserId,
	}
	if ActionType == "1" {
		err = db.CreateRelation(relation)
		return
	} else if ActionType == "2" {
		err = db.DeleteRelation(relation)
		return
	} else {
		return errors.New("请求参数不规范")
	}
}

func FollowList(c *gin.Context) (resp model.FollowListResponse, err error) {

	return
}

func GetFollowList(uid uint) (UserList []model.UserResp, err error) {

	return
}

func FollowerList(c *gin.Context) (resp model.FollowerListResponse, err error) {

	return
}
