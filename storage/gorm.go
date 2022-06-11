package storage

import (
	"TikTok/config"
	"TikTok/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDb() error {
	logger := log.NewMysqlLogger()
	var err error
	dsn := config.UserName + ":" + config.Password + "@tcp(" + config.HOST + ")/" + config.DbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger})
	if err != nil {
		return err
	}
	return nil
}
