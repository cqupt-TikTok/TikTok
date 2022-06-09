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
