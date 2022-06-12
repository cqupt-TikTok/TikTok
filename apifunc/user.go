package apifunc

import (
	"TikTok/dbfunc"
	"TikTok/model"
	"TikTok/util"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

// Register 用户注册
func Register(c *gin.Context) (model.UserResponse, error) {
	var userResponse model.UserResponse
	var token string
	username := c.Query("username")
	password := c.Query("password")
	userId, err := dbfunc.Register(username, password)
	if err != nil {
		return userResponse, err
	}
	token, err = util.SetToken(username, userId, time.Now().Add(time.Hour*240))
	if err != nil {
		return userResponse, err
	}
	userResponse.UserId = userId
	userResponse.Token = token
	return userResponse, nil
}

// LoginCheck 用户登录
func LoginCheck(c *gin.Context) (model.UserResponse, error) {
	var userResponse model.UserResponse
	var token string
	username := c.Query("username")
	password := c.Query("password")
	userId, err := dbfunc.Login(username, password)
	if err != nil {
		return userResponse, err
	}
	token, err = util.SetToken(username, userId, time.Now().Add(time.Hour*240))
	if err != nil {
		return userResponse, err
	}
	userResponse.UserId = userId
	userResponse.Token = token
	return userResponse, nil
}

// UserInfo 用户信息
func UserInfo(c *gin.Context) (model.UserInfoResponse, error) {
	var userInfoResponse model.UserInfoResponse
	token := c.Query("token")
	key, err := util.CheckToken(token)
	if err != nil {
		return userInfoResponse, err
	}
	userId64, _ := strconv.ParseUint(c.Query("user_id"), 10, 64)
	userId := uint(userId64)
	userInfoResponse, err = dbfunc.UserInfo(userId, key.UserId)
	if err != nil {
		return userInfoResponse, err
	}
	return userInfoResponse, nil
}
