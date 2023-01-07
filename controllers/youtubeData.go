package controllers

import (
	"errors"
	"fmt"

	models "youtubesync/models"
	repository "youtubesync/repository"
)

func SearchController(request map[string]interface{}) ([]models.YoutubeData, error) {
	title := ""
	description := ""
	var limit int64
	if val, ok := request["title"].(string); ok {
		title = val
	}
	if val, ok := request["description"].(string); ok {
		description = val
	}
	if val, ok := request["limit"].(float64); ok {
		limit = int64(val)
	}
	if title == "" && description == "" {
		return make([]models.YoutubeData, 0), errors.New("no title or description found")
	}
	// if no limit is set so by default value is 10
	if limit == int64(0) {
		limit = 10
	}

	searchData, err := repository.SearchByTitleOrDescrition(title, description, limit)
	if err != nil {
		return make([]models.YoutubeData, 0), err
	}
	fmt.Println("search resp", searchData)
	return searchData, nil
}
