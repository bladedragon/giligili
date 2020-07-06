package model

import (
	"giligili/cache"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/jinzhu/gorm"
	"os"
	"strconv"
	"strings"
)

type Video struct {
	gorm.Model
	Title  string
	Info   string
	URL    string
	Avatar string
}

func (video *Video) AvatarURL() string {
	client, _ := oss.New(os.Getenv("OSS_END_POINT"), os.Getenv("OSS_ACCESS_KEY_ID"), os.Getenv("OSS_ACCESS_KEY_SECRET"))
	bucket, _ := client.Bucket(os.Getenv("OSS_BUCKET"))
	signedGetURL, _ := bucket.SignURL(video.Avatar, oss.HTTPGet, 60)
	if strings.Contains(signedGetURL, "http://giligili-img-av.oss-cn-hangzhou.aliyuncs.com/?Exp") {
		signedGetURL = "https://giligili-img-av.oss-cn-hangzhou.aliyuncs.com/img/noface.png"
	}
	return signedGetURL
}

func (video *Video) VideoURL() string {
	return ""
}

func (video *Video) View() uint64 {
	countStr, _ := cache.RedisClient.Get(cache.VideoViewKey(video.ID)).Result()
	count, _ := strconv.ParseUint(countStr, 10, 64)
	return count
}

func (video *Video) AddView() {
	cache.RedisClient.Incr(cache.VideoViewKey(video.ID))

	cache.RedisClient.ZIncrBy(cache.DailyRankKey, 1, strconv.Itoa(int(video.ID)))

}
