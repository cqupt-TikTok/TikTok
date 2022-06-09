// @Title : relation
// @Description :
// @Author : MX
// @Update : 2022/6/6 17:36

package function

import (
	"TikTok/dbfunc"
	"errors"
	"strconv"

	"TikTok/model"
	"github.com/gin-gonic/gin"
)

func FollowAction(c *gin.Context) (err error) {
	toUserIdStr := c.Param("to_user_id")
	toUserIdInt, err := strconv.Atoi(toUserIdStr)
	if err != nil {
		return errors.New("请求参数不规范")
	}
	ToUserId := uint(toUserIdInt)

	ActionType := c.Param("action_type")

	UserIdValue, _ := c.Get("uid")
	UserId := UserIdValue.(uint)

	relation := model.FollowRelation{
		FollowerId: ToUserId,
		UserId:     UserId,
	}
	if ActionType == "1" {
		err = dbfunc.CreateRelation(relation)
		return
	} else if ActionType == "2" {
		err = dbfunc.DeleteRelation(relation)
		return
	} else {
		return errors.New("请求参数不规范")
	}
}

func FollowList(c *gin.Context) (FollowList []model.UserResp, err error) {
	userIdStr := c.Param("user_id")
	userIdInt, err := strconv.Atoi(userIdStr)
	if err != nil {
		return nil, errors.New("请求参数不规范")
	}
	userId := uint(userIdInt)

	followCount, err := dbfunc.GetFollowCount(userId)
	if err != nil {
		return
	}

	userIds, err := dbfunc.GetFollowIds(userId, followCount)
	if err != nil {
		return
	}

	FollowList, err = GetUserList(userIds, followCount)
	if err != nil {
		return
	}

	return
}

func FollowerList(c *gin.Context) (FollowerList []model.UserResp, err error) {
	userIdStr := c.Param("user_id")
	userIdInt, err := strconv.Atoi(userIdStr)
	if err != nil {
		return nil, errors.New("请求参数不规范")
	}
	userId := uint(userIdInt)

	followerCount, err := dbfunc.GetFollowerCount(userId)
	if err != nil {
		return
	}

	userIds, err := dbfunc.GetFollowerIds(userId, followerCount)
	if err != nil {
		return
	}

	FollowerList, err = GetUserList(userIds, followerCount)
	if err != nil {
		return
	}

	return
}

func GetUserList(uids []uint, size int64) (UserList []model.UserResp, err error) {
	UserList = make([]model.UserResp, 0, size)
	Users, err := dbfunc.GetUsers(uids, size)
	if err != nil {
		return nil, err
	}

	for _, user := range Users {
		UserList = append(UserList, user.ToResp())
	}

	return
}
