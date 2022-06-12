package apifunc

import (
	"TikTok/config"
	"TikTok/dbfunc"
	"TikTok/model"
	"TikTok/storage"
	"TikTok/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

// Publish 发布视频
func Publish(c *gin.Context) error {
	var key *util.MyClaims
	var playUrl string
	var coverUrl string
	var err error
	//form, err := c.MultipartForm()
	//if err != nil {
	//	return err
	//}
	//token := form.Value["token"]
	//key, err = util.CheckToken(token[0])
	//title := form.Value["title"][0]
	//data := form.File["data"]
	token := c.PostForm("token")
	key, err = util.CheckToken(token)
	if err != nil {
		return err
	}
	Tid := key.UserId
	title := c.PostForm("title")
	data, _ := c.FormFile("data")
	file, err := data.Open()
	if err != nil {
		return err
	}
	videoName := strconv.FormatInt(time.Now().Unix(), 10) + strconv.Itoa(int(key.UserId)) + ".mp4"
	playUrl = config.ImgUrl + videoName
	coverUrl = playUrl + "?vframe/jpg/offset/1"
	err = storage.UpLoadFile(file, videoName, data.Size)
	if err != nil {
		fmt.Println(err)
		fmt.Println("upLoade "+videoName+" false: ", err.Error())
		return err
	} else {
		fmt.Println("upLoade " + videoName + " finished")
	}
	err = dbfunc.Publish(Tid, title, playUrl, coverUrl)
	if err != nil {
		return err
	}
	return nil
}

// PublishList 获取发布列表
func PublishList(c *gin.Context) (resp model.PostListResponse, err error) {
	var key *util.MyClaims
	token := c.Query("token")
	key, err = util.CheckToken(token)
	if err != nil {
		return resp, err
	}
	Tid := key.UserId
	userId64, _ := strconv.ParseUint(c.Query("user_id"), 10, 64)
	userId := uint(userId64)
	resp.VideoList, err = dbfunc.PostList(userId, Tid)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
