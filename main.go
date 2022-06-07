package main

import (
	"TikTok/db"
	"TikTok/model"
	"TikTok/router"
	"fmt"
	"gorm.io/gorm"
)

func main() {
	err := db.InitDb()
	if err != nil {
		panic(err)
	}
	err = db.DB.AutoMigrate(&model.User{}, &model.Video{}, &model.Comment{}, &model.FollowRelation{}, &model.FavoriteVideoRelation{})
	if err != nil {
		return
	}
	err = router.InitRouter()
	if err != nil {
		panic(err)
	}
	//结构体方法调用示例
	var u model.User
	if err := db.DB.Where("id = ?", 1).First(&u).Error; err == gorm.ErrRecordNotFound {
		fmt.Println(err)
	} else {
		fmt.Println(u)
		r := u.ToResp()
		r.IsFollowJudge(2) //此方法根据应用场景自行调用。如：关注列表返回时，全应为true，无需调用，查看用户信息时，则需调用
		fmt.Println(r)
	}
}
