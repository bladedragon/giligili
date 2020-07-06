package service

import (
	"giligili/model"
	"giligili/serializer"
)

type ShowVideoService struct{}

func (service *ShowVideoService) Show(id string) serializer.Response {
	var video model.Video
	err := model.DB.First(&video, id).Error
	if err != nil {
		return serializer.Response{
			Status: 404,
			Msg:    "视频不存在",
			Error:  err.Error(),
		}
	}
	//抽象出来处理观看相关问题
	video.AddView()

	return serializer.Response{
		Data: serializer.BuildVideo(video),
	}
}
