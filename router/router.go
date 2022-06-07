package router

import (
	"TikTok/api"
	"github.com/gin-gonic/gin"
)

func InitRouter() error {
	r := gin.Default()
	r.POST("/douyin/comment/action/", api.CommentAction)

	err := r.Run(":7070")
	if err != nil {

	}
	return nil
}
