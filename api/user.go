package api

import (
	"TikTok/apifunc"
	"TikTok/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Register 用户注册
func Register(c *gin.Context) {
	var resp model.UserResponse
	var err error
	resp, err = apifunc.Register(c)
	if err != nil {
		resp.StatusCode = -1
		resp.StatusMsg = "注册失败:" + err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.StatusCode = 0
	resp.StatusMsg = "注册成功"
	c.JSON(http.StatusOK, resp)
	return
}

// Login 用户登录
func Login(c *gin.Context) {
	var resp model.UserResponse
	var err error
	resp, err = apifunc.LoginCheck(c)
	if err != nil {
		resp.StatusCode = -1
		resp.StatusMsg = "登录失败:" + err.Error()
		c.JSON(http.StatusOK, resp)
		return
	}
	resp.StatusCode = 0
	resp.StatusMsg = "登录成功"
	c.JSON(http.StatusOK, resp)
	return
}

// UserInfo 用户信息
func UserInfo(c *gin.Context) {
	var resp model.UserInfoResponse
	var err error
	resp, err = apifunc.UserInfo(c)
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
