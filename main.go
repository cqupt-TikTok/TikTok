package main

import (
	"TikTok/gorm"
	"TikTok/model"
	"TikTok/router"
)

func main() {
	err := gorm.InitDb()
	if err != nil {
		panic(err)
	}

	err = gorm.DB.AutoMigrate(&model.User{}, &model.Video{}, &model.Comment{}, &model.FollowRelation{}, &model.FavoriteVideoRelation{})
	if err != nil {
		panic(err)
	}

	err = router.InitRouter()
	if err != nil {
		panic(err)
	}
}
