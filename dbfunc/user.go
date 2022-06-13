package dbfunc

import (
	"TikTok/model"
	"TikTok/storage"
	"TikTok/util"
	"errors"
)

// Register 注册储存用户信息
func Register(username, password string) (uint, error) {
	var tempUser model.User
	storage.DB.Where("name = ?", username).First(&tempUser)
	if tempUser.ID > 0 {
		return 0, errors.New("用户名已存在")
	}
	tempUser.Name = username
	tempUser.Password = util.ScryptPw(password)
	tempUser.Signature = "测试！测试！！测试！！！"
	tempUser.BackgroundImage = "http://zbwzbw.top/background-image-test2.jpeg"
	tempUser.Avatar = "http://zbwzbw.top/avatar-test-2.jpg.jpg"
	tx := storage.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return 0, err
	}
	err := tx.Create(&tempUser).Error
	if err != nil {
		tx.Rollback()
		return 0, errors.New("创建失败")
	}
	err = tx.Where("name = ?", username).First(&tempUser).Error
	if err != nil || tempUser.ID <= 0 {
		tx.Rollback()
		return 0, err
	}
	return tempUser.ID, tx.Commit().Error
}

// Login 登录验证
func Login(username, password string) (uint, error) {
	var tempUser model.User
	storage.DB.Where("name = ?", username).First(&tempUser)
	if tempUser.ID <= 0 {
		return 0, errors.New("用户名不存在")
	}
	if tempUser.Password != util.ScryptPw(password) {
		return 0, errors.New("密码错误")
	}
	return tempUser.ID, nil
}

// UserInfo 获取用户详情userid要查询用户的id，Tid登录用户token中的id
func UserInfo(userId, Tid uint) (model.UserInfoResponse, error) {
	var user model.User
	var userResp model.UserResp
	var userInfoResponse model.UserInfoResponse
	err := storage.DB.Where("id = ?", userId).First(&user).Error
	if err != nil {
		return userInfoResponse, err
	}
	userResp = user.ToResp()
	userResp.IsFollowJudge(Tid)
	userInfoResponse.UserResp = userResp
	return userInfoResponse, nil
}
