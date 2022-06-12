package storage

import (
	"TikTok/config"
	"context"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"mime/multipart"
)

// UpLoadFile 上传文件
func UpLoadFile(file multipart.File, fileName string, fileSize int64) error {
	putPolicy := storage.PutPolicy{
		Scope: config.Bucket,
	}
	mac := qbox.NewMac(config.AccessKey, config.SecretKey)
	upToken := putPolicy.UploadToken(mac)
	cfg := storage.Config{
		Zone:          &storage.ZoneHuadong,
		UseCdnDomains: true,
		UseHTTPS:      false,
	}
	resumeUploader := storage.NewResumeUploaderV2(&cfg)
	ret := storage.PutRet{}
	storage.SetSettings(&storage.Settings{
		TaskQsize: 40,
		Workers:   10,
		ChunkSize: 2 * 1024 * 1024,
		PartSize:  2 * 1024 * 1024,
		TryTimes:  2,
	})
	putExtra := storage.RputV2Extra{}
	var err error
	err = resumeUploader.Put(context.Background(), &ret, upToken, fileName, file, fileSize, &putExtra)
	if err != nil {
		return err
	}
	return nil
}
