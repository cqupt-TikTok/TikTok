package router

import (
	"TikTok/api"
	"github.com/gin-gonic/gin"
)

func InitRouter() error {
	r := gin.Default()
	apiRouter := r.Group("/douyin")

	// basic apis
	//apiRouter.GET("/feed/", controller.Feed)
	apiRouter.GET("/user/", api.UserInfo)
	apiRouter.POST("/user/register/", api.Register)
	apiRouter.POST("/user/login/", api.Login)
	//apiRouter.POST("/publish/action/", controller.Publish)
	//apiRouter.GET("/publish/list/", controller.PublishList)

	// extra apis - I
	apiRouter.POST("/favorite/action/", api.FavoriteAction)
	apiRouter.GET("/favorite/list/", api.FavoriteVideoList)
	apiRouter.POST("/comment/action/", api.CommentAction)
	apiRouter.GET("/comment/list/", api.CommentList)

	// extra apis - II
	apiRouter.POST("/relation/action/", api.FollowAction)
	apiRouter.GET("/relation/follow/list/", api.FollowList)
	apiRouter.GET("/relation/follower/list/", api.FollowerList)

	err := r.Run(":8080")
	if err != nil {
		return err
	}
	return nil
}
