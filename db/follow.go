// @Title : relation
// @Description :
// @Author : MX
// @Update : 2022/6/6 17:35

package db

import "TikTok/model"

func CreateRelation(relation model.FollowRelation) (err error) {
	err = DB.Create(&relation).Error
	if err != nil {
		return err
	}
	return
}

func DeleteRelation(relation model.FollowRelation) (err error) {
	err = DB.Where("follower_id = ? AND user_id = ?", relation.FollowerId, relation.UserId).
		Delete(&model.FollowRelation{}).Error
	if err != nil {
		return
	}
	return
}

func GetFollowCount(uid uint) (FollowCount int64, err error) {
	var user model.User
	err = DB.Model(&user).Select("follow_count").First(uid).Error
	if err != nil {
		return
	}
	return user.FollowCount, err
}

func GetFollowerCount(uid uint) (FollowCount int64, err error) {
	var user model.User
	err = DB.Model(&user).Select("follower_count").First(uid).Error
	if err != nil {
		return
	}
	return user.FollowerCount, err
}

func GetFollowIds(uid uint, size int64) (uids []uint, err error) {
	// 通过提前获取切片大小, 提前为切片分配空间, 避免重复分配内存影响性能
	relations := make([]model.FollowRelation, 0, size)
	uids = make([]uint, 0, size)

	err = DB.Model(&model.FollowRelation{}).Select("follower_id").
		Where("user_id = ?", uid).Find(&relations).Error

	for _, relation := range relations {
		uid := relation.FollowerId
		uids = append(uids, uid)
	}
	return
}

func GetFollowerIds(uid uint, size int64) (uids []uint, err error) {
	// 通过提前获取切片大小, 提前为切片分配空间, 避免重复分配内存影响性能
	relations := make([]model.FollowRelation, 0, size)
	uids = make([]uint, 0, size)

	err = DB.Model(&model.FollowRelation{}).Select("user_id").
		Where("follower_id = ?", uid).Find(&relations).Error

	for _, relation := range relations {
		uid := relation.FollowerId
		uids = append(uids, uid)
	}
	return
}

func GetUsers(uids []uint, size int64) (users []model.User, err error) {
	users = make([]model.User, 0, size)
	err = DB.Model(&model.User{}).Where("id IN ?", uids).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return
}
