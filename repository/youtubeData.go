package repository

import (
	"errors"
	"fmt"
	"time"
	models "youtubesync/models"
	database "youtubesync/utils/database"
)

// models "youtubesync/models"
// database "youtubesync/utils"

func GetDataPaginated(limit int64, page int64) ([]models.YoutubeData, error) {
	// youtubeData := models.YoutubeData{}
	// val := database.GDB.Table("youtube_data").Where(filters).First(&youtubeData)
	youtubeData := []models.YoutubeData{}
	query := fmt.Sprintf("select * from youtube_data order by published_date_time limit %v,%v", limit*(page-1), limit)
	result := database.GDB.Debug().Raw(query).Find(&youtubeData)
	if result.Error != nil {
		return youtubeData, result.Error
	}
	fmt.Println("result", result)
	return youtubeData, result.Error
}

func CreateYoutubeData(youtubeData *models.YoutubeData) error {
	query := fmt.Sprintf("insert into youtube_data(published_date_time,title,description,channel_id,channel_title,thumbnail_url,video_id) values(%v,%v,%v,%v,%v,%v,%v) WHERE NOT EXISTS (select * from youtube_data where video_id=%v)",
		*youtubeData.PublishedDateTime,
		*youtubeData.Title,
		*youtubeData.Description,
		*youtubeData.ChannelId,
		*youtubeData.ChannelTitle,
		*youtubeData.ThumbnailUrl,
		*youtubeData.VideoId,
		*youtubeData.VideoId)

	result := database.GDB.Debug().Raw(query)
	// result := database.GDB.Table("youtube_data").Create(youtubeData)
	return result.Error
}

func GetLastPublished() (*time.Time, error) {
	youtubeData := models.YoutubeData{}
	query := "select * from youtube_data order by published_date_time desc limit 1"
	result := database.GDB.Debug().Raw(query).Find(&youtubeData)
	fmt.Println("result", youtubeData.PublishedDateTime)
	if result.Error != nil {
		return nil, result.Error
	}
	if youtubeData.PublishedDateTime == nil {
		return nil, errors.New("last published time not present")
	}

	return youtubeData.PublishedDateTime, result.Error
}

func SearchByTitleOrDescrition(title string, description string, limit int64) ([]models.YoutubeData, error) {
	youtubeData := []models.YoutubeData{}
	query := fmt.Sprintf("select * from youtube_data where  description like '%%%v%%' or title like '%%%v%%' order by published_date_time desc limit %v", description, title, limit)
	result := database.GDB.Debug().Raw(query).Find(&youtubeData)
	if result.Error != nil {
		return youtubeData, result.Error
	}
	fmt.Println("result", result)
	return youtubeData, result.Error
}
