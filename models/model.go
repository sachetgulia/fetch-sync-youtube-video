package models

import "time"

type YoutubeData struct {
	Id                int64     `gorm:"type:BIGINT UNSIGNED AUTO_INCREMENT;INDEX" json:"id"`
	Created           time.Time `gorm:"type:DATETIME DEFAULT CURRENT_TIMESTAMP; autoCreateTime" json:"created"`
	PublishedDateTime *time.Time `gorm:"type:DATETIME DEFAULT CURRENT_TIMESTAMP;INDEX" json:"published_date_time"`
	Title             *string   `gorm:"type:VARCHAR(300)" json:"title"`
	Description       *string   `gorm:"type:VARCHAR(500)" json:"description"`
	ChannelId         *string   `gorm:"type:VARCHAR(100)" json:"channel_id"`
	ChannelTitle      *string   `gorm:"type:VARCHAR(200)" json:"channel_title"`
	ThumbnailUrl      *string   `gorm:"type:VARCHAR(500)" json:"thumbnail_url"`
	VideoId           *string   `gorm:"type:VARCHAR(100)" json:"video_id"`
}
