package db

import (
	"TikTok/config"
	"TikTok/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDb() error {
	var err error
	dsn := config.UserName + ":" + config.Password + "@tcp(" + config.HOST + ")/" + config.DbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	err = DB.AutoMigrate(&model.User{}, &model.Video{}, &model.Comment{}, &model.FollowRelation{}, &model.FavoriteVideoRelation{})
	if err != nil {
		return err
	}
	return nil

}
