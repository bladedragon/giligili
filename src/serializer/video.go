package serializer

import "giligili/model"

type Video struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Info      string `json:"info"`
	URL       string `json:"url"`
	Avatar    string `json:"avatar"`
	View      uint64 `json:"view"`
	CreatedAt int64  `json:"created_at"`
}

func BuildVideo(item model.Video) Video {
	return Video{
		ID:        item.ID,
		Title:     item.Title,
		Info:      item.Info,
		URL:       item.URL,
		Avatar:    item.Avatar,
		View:      item.View(),
		CreatedAt: item.CreatedAt.Unix(),
	}
}

func BuildVideos(items []model.Video) (videos []Video) {
	for _, item := range items {
		video := BuildVideo(item)
		videos = append(videos, video)
	}
	return videos
}
