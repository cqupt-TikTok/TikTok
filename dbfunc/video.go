package dbfunc

import (
	"TikTok/model"
	"TikTok/storage"
	"time"
)

func Feed(lastTime int64, Tid uint) (videoList []model.VideoResp, nextTime int64, err error) {
	var videos []model.Video
	var Vresp model.VideoResp
	err = storage.DB.Order("created_at").Where("created_at < FROM_UNIXTIME(?)", lastTime).Limit(30).Find(&videos).Error
	if err != nil {
		return videoList, time.Now().Unix(), err
	}
	for _, v := range videos {
		Vresp = v.ToResp(Tid)
		Vresp.IsFavoriteJudge(Tid)
		Vresp.Author.IsFollowJudge(Tid)
		videoList = append(videoList, Vresp)
	}
	if len(videos) == 0 {
		return videoList, time.Now().Unix(), nil
	}
	nextTime = videos[0].CreatedAt.Unix()
	return videoList, nextTime, nil
}
func Publish(Tid uint, Title string, playUrl string, coverUrl string) error {
	var video = model.Video{
		AuthorId: Tid,
		Title:    Title,
		PlayUrl:  playUrl,
		CoverUrl: coverUrl,
	}
	err := storage.DB.Create(&video).Error
	if err != nil {
		return err
	}
	return nil
}
func PostList(userId, Tid uint) (videoList []model.VideoResp, err error) {
	var videos []model.Video
	var Vresp model.VideoResp
	err = storage.DB.Where("author_id = ?", userId).Find(&videos).Error
	if err != nil {
		return videoList, err
	}
	for _, v := range videos {
		Vresp = v.ToResp(userId)
		Vresp.IsFavoriteJudge(Tid)
		videoList = append(videoList, Vresp)
	}
	return videoList, nil
}
