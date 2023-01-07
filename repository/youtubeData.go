package repository

import (
	"fmt"
	"time"
	"youtubesync/models"
	database "youtubesync/utils/database"

	"gorm.io/gorm"
)

// models "youtubesync/models"
// database "youtubesync/utils"

func GetDataByFilters(filters map[string]interface{}) error {
	// youtubeData := models.YoutubeData{}
	// val := database.GDB.Table("youtube_data").Where(filters).First(&youtubeData)

	return nil
}

func CreateYoutubeData(tx *gorm.DB, youtubeData interface{}) error {
	if tx == nil {
		tx = database.GDB
	}
	result := tx.Table("youtube_data").Create(youtubeData)
	return result.Error
}

func GetLastPublished(tx *gorm.DB) (*time.Time, error) {
	if tx == nil {
		tx = database.GDB
	}
	query := "select published_date_time from youtube_data order by published_date_time desc limit 1"
	result := tx.Raw(query)
	if result.Error != nil {
		return nil, result.Error
	}
	fmt.Println("result", result)
	return nil, result.Error
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
