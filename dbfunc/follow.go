// @Title : relation
// @Description :
// @Author : MX
// @Update : 2022/6/6 17:35

package dbfunc

import (
	"TikTok/model"
	"TikTok/storage"
	"errors"
	"gorm.io/gorm"
)

// CreateRelation 创建用户关系，关注用户
func CreateRelation(relation model.FollowRelation) (err error) {
	//开始事务
	tx := storage.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}
	//关注用户是否存在
	var u model.User
	if err := tx.Model(&model.User{}).Where("id = ?", relation.FollowerId).First(&u).Error; err != nil {
		tx.Rollback()
		return err
	}
	//查询是否已经关注
	var FR model.FollowRelation
	tx.Model(&model.FollowRelation{}).Where("follower_id = ? and user_id = ?", relation.FollowerId, relation.UserId).First(&FR)
	if FR.Id > 0 {
		tx.Rollback()
		return errors.New("重复关注")
	}
	//用户关注总数follow_count+1
	if err := tx.Model(&model.User{}).Where("id = ? ", relation.UserId).Update("follow_count", gorm.Expr("follow_count+ ?", 1)).Error; err != nil {
		tx.Rollback()
		return err
	}
	//被关注用户粉丝总数follower_count+1
	if err := tx.Model(&model.User{}).Where("id = ? ", relation.FollowerId).Update("follower_count", gorm.Expr("follower_count+ ?", 1)).Error; err != nil {
		tx.Rollback()
		return err
	}
	//关注表中写入数据
	if err := tx.Create(&relation).Error; err != nil {
		tx.Rollback()
		return err
	}
	//提交事务
	return tx.Commit().Error
}

// DeleteRelation 删除用户关系，取消关注
func DeleteRelation(relation model.FollowRelation) (err error) {
	//开始事务
	tx := storage.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}
	//关注用户是否存在
	var u model.User
	if err := tx.Model(&model.User{}).Where("id = ?", relation.FollowerId).First(&u).Error; err != nil {
		tx.Rollback()
		return err
	}
	//查询是否已经关注
	var FR model.FollowRelation
	tx.Model(&model.FollowRelation{}).Where("follower_id = ? and user_id = ?", relation.FollowerId, relation.UserId).First(&FR)
	if FR.Id <= 0 {
		tx.Rollback()
		return errors.New("还未关注")
	}
	//用户关注总数follow_count-1
	if err := tx.Model(&model.User{}).Where("id = ? ", relation.UserId).Update("follow_count", gorm.Expr("follow_count- ?", 1)).Error; err != nil {
		tx.Rollback()
		return err
	}
	//被关注用户粉丝总数follower_count-1
	if err := tx.Model(&model.User{}).Where("id = ? ", relation.FollowerId).Update("follower_count", gorm.Expr("follower_count- ?", 1)).Error; err != nil {
		tx.Rollback()
		return err
	}
	//关注表中写入数据
	err = tx.Where("follower_id = ? AND user_id = ?", relation.FollowerId, relation.UserId).Delete(&model.FollowRelation{}).Error
	if err != nil {
		return err
	}
	//提交事务
	return tx.Commit().Error
}

// GetFollowCount 获取用户关注总数
func GetFollowCount(uid uint) (FollowCount int64, err error) {
	var user model.User
	err = storage.DB.Model(&user).Where("id = ?", uid).First(&user).Error
	if err != nil {
		return
	}
	return user.FollowCount, err
}

// GetFollowerCount 获取用户粉丝总数
func GetFollowerCount(uid uint) (FollowCount int64, err error) {
	var user model.User
	err = storage.DB.Model(&user).Where("id = ?", uid).First(&user).Error
	if err != nil {
		return
	}
	return user.FollowerCount, err
}

// GetFollowIds 获取用户关注的用户id
func GetFollowIds(uid uint, size int64) (uids []uint, err error) {
	// 通过提前获取切片大小, 提前为切片分配空间, 避免重复分配内存影响性能
	relations := make([]model.FollowRelation, 0, size)
	uids = make([]uint, 0, size)

	err = storage.DB.Model(&model.FollowRelation{}).Select("follower_id").
		Where("user_id = ?", uid).Find(&relations).Error

	for _, relation := range relations {
		uid := relation.FollowerId
		uids = append(uids, uid)
	}
	return
}

// GetFollowerIds 获取用户粉丝的id
func GetFollowerIds(uid uint, size int64) (uids []uint, err error) {
	// 通过提前获取切片大小, 提前为切片分配空间, 避免重复分配内存影响性能
	relations := make([]model.FollowRelation, 0, size)
	uids = make([]uint, 0, size)

	err = storage.DB.Model(&model.FollowRelation{}).Where("follower_id = ?", uid).Find(&relations).Error
	for _, relation := range relations {
		uid := relation.UserId
		uids = append(uids, uid)
	}
	return
}

// GetUsers 查询绑定用户信息
func GetUsers(uids []uint, size int64) (users []model.User, err error) {
	users = make([]model.User, 0, size)
	err = storage.DB.Model(&model.User{}).Where("id IN ?", uids).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return
}
