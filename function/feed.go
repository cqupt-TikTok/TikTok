package function

import (
	"TikTok/dbfunc"
	"TikTok/model"
	"TikTok/util"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func Feed(c *gin.Context) (resp model.FeedResponse, err error) {
	var key *util.MyClaims
	var Tid uint
	token := c.Query("token")
	if token == "" {
		Tid = 0
	} else {
		key, err = util.CheckToken(token)
		if err != nil {
			return resp, err
		}
		Tid = key.UserId
	}
	var lastTime int64
	lastTimeStr := c.Query("last_time")
	if lastTimeStr == "" || lastTimeStr == "0" {
		lastTime = time.Now().Unix()
	} else {
		lastTime, err = strconv.ParseInt(lastTimeStr, 10, 64)
	}
	resp.VideoList, resp.NextTime, err = dbfunc.Feed(lastTime, Tid)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
