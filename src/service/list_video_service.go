package service

import (
	"giligili/model"
	"giligili/serializer"
)

type ListVideoService struct {
	Limit int `form:"limit"`
	Start int `form:"start"`
}

func (service *ListVideoService) List() serializer.Response {
	videos := []model.Video{}
	total := 0

	if service.Limit == 0 {
		service.Limit = 6
	}

	if err := model.DB.Limit(service.Limit).Offset(service.Start).Find(&videos).Error; err != nil {
		return serializer.Response{
			Status: 50000,
			Msg:    "数据库连接错误",
			Error:  err.Error(),
		}
	}
	return serializer.BuildListResponse(serializer.BuildVideos(videos), uint(total))
}
