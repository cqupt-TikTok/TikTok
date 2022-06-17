package router

import (
	"TikTok/api"
	"TikTok/log"
	"TikTok/util"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

func InitRouter() error {
	gin.SetMode("debug")
	r := gin.New()
	r.Use(log.ApiLogger())
	r.Use(gin.Recovery())
	apiRouter := r.Group("/douyin")
	pprof.Register(r)
	// basic apis
	apiRouter.GET("/feed/", api.Feed)
	apiRouter.GET("/user/", util.JWT(), api.UserInfo)
	apiRouter.POST("/user/register/", api.Register)
	apiRouter.POST("/user/login/", api.Login)
	apiRouter.POST("/publish/action/", api.Publish)
	apiRouter.GET("/publish/list/", util.JWT(), api.PublishList)

	// extra apis - I
	apiRouter.POST("/favorite/action/", util.JWT(), api.FavoriteAction)
	apiRouter.GET("/favorite/list/", util.JWT(), api.FavoriteVideoList)
	apiRouter.POST("/comment/action/", util.JWT(), api.CommentAction)
	apiRouter.GET("/comment/list/", util.JWT(), api.CommentList)

	// extra apis - II
	apiRouter.POST("/relation/action/", util.JWT(), api.FollowAction)
	apiRouter.GET("/relation/follow/list/", util.JWT(), api.FollowList)
	apiRouter.GET("/relation/follower/list/", util.JWT(), api.FollowerList)

	err := r.Run(":8080")
	if err != nil {
		return err
	}
	return nil
}
