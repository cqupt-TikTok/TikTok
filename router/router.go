package router

import (
	"TikTok/api"
	"github.com/gin-gonic/gin"
)

func InitRouter() error {
	r := gin.Default()
	r.POST("/douyin/comment/action/", api.CommentAction)
	r.GET("/douyin/comment/list/", api.CommentList)
	err := r.Run(":8080")
	if err != nil {

	}
	return nil
}
