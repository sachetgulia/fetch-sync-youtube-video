package dtos

import (
	"time"
)

type YoutubeDataDtos struct {
	PublishedDateTime time.Time `json:"published_date_time"`
	Title             *string    `json:"title"`
	Description       *string    `json:"description"`
	ChannelId         *string    `json:"channel_id"`
	ChannelTitle      *string    `json:"channel_title"`
	ThumbnailUrl      *string    `json:"thumbnail_url"`
	VideoId           *string    `json:"video_id"`
}
