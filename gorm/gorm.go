package gorm

import (
	"TikTok/config"
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
	return nil
}
