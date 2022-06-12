// @Title : relation
// @Description :
// @Author : MX
// @Update : 2022/6/6 17:36

package apifunc

import (
	"TikTok/dbfunc"
	"TikTok/util"
	"errors"
	"strconv"
	"time"

	"TikTok/model"
	"github.com/gin-gonic/gin"
)

func FollowAction(c *gin.Context) error {
	toUserIdStr := c.Query("to_user_id")
	toUserId64, err := strconv.Atoi(toUserIdStr)
	if err != nil {
		return errors.New("请求参数不规范")
	}
	ToUserId := uint(toUserId64)
	ActionType := c.Query("action_type")
	var key *util.MyClaims
	var Tid uint
	token := c.Query("token")
	if token == "" {
		Tid = 0
	} else {
		key, err = util.CheckToken(token)
		if err != nil {
			return err
		}
		Tid = key.UserId
	}
	relation := model.FollowRelation{
		FollowerId: ToUserId,
		UserId:     Tid,
		FollowDate: time.Now(),
	}
	if ActionType == "1" {
		err = dbfunc.CreateRelation(relation)
		return err
	} else if ActionType == "2" {
		err = dbfunc.DeleteRelation(relation)
		return err
	} else {
		return errors.New("请求参数不规范")
	}
}

func FollowList(c *gin.Context) ([]model.UserResp, error) {
	var FollowList []model.UserResp
	userIdStr := c.Query("user_id")
	userId64, err := strconv.Atoi(userIdStr)
	if err != nil {
		return nil, errors.New("请求参数不规范")
	}
	userId := uint(userId64)
	followCount, err := dbfunc.GetFollowCount(userId)
	if err != nil {
		return FollowList, err
	}
	userIds, err := dbfunc.GetFollowIds(userId, followCount)
	if err != nil {
		return FollowList, err
	}
	FollowList, err = GetUserList(userIds, followCount)
	if err != nil {
		return FollowList, err
	}
	return FollowList, nil
}

func FollowerList(c *gin.Context) (FollowerList []model.UserResp, err error) {
	userIdStr := c.Query("user_id")
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
