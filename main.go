package main

import (
	"TikTok/model"
	"TikTok/router"
	"TikTok/storage"
)

func main() {
	err := storage.InitDb()
	if err != nil {
		panic(err)
	}

	err = storage.DB.AutoMigrate(&model.User{}, &model.Video{}, &model.Comment{}, &model.FollowRelation{}, &model.FavoriteVideoRelation{})
	if err != nil {
		panic(err)
	}

	err = router.InitRouter()
	if err != nil {
		panic(err)
	}
}
