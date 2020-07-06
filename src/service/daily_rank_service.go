package service

import (
	"fmt"
	"giligili/cache"
	"giligili/model"
	"giligili/serializer"
	"strings"
)

type DailyRankService struct{}

func (service *DailyRankService) Get() serializer.Response {

	var videos []model.Video

	vids, _ := cache.RedisClient.ZRange(cache.DailyRankKey, 0, 5).Result()

	if len(vids) > 1 {
		order := fmt.Sprintf("FIELD(id,%s", strings.Join(vids, ","))
		err := model.DB.Where("id in (?)", vids).Order(order).Find(&videos).Error
		if err != nil {
			return serializer.Response{
				Status: 50000,
				Msg:    "数据库连接错误",
				Error:  err.Error(),
			}
		}
	}

	return serializer.Response{
		Data: serializer.BuildVideos(videos),
	}

}
