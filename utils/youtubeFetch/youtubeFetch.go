package youtubeFetch

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
	dtos "youtubesync/dtos"
	models "youtubesync/models"
	repository "youtubesync/repository"
	constants "youtubesync/utils"
	parseTime "youtubesync/utils/helpers"
)

func FetchInit() error {

	for {
		// lastPublishedDataTime, _ := repository.GetLastPublished(nil)
		// fmt.Println("las t",lastPublishedDataTime)
		// fetch data in every 10 sec
		//get request call
		maxResults := 5
		secretKey := "AIzaSyB12amtMkVda3obVW1O-U39P-8t2NMH3v4"
		url := fmt.Sprintf(constants.YoutubeUrl, maxResults, secretKey)
		resp, err := http.Get(url)
		if err != nil {
			return err
		}
		//We Read the response body on the line below.
		responseRaw, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		var response map[string]interface{}
		err = json.Unmarshal(responseRaw, &response)
		// sb := string(responseRaw)
		// fmt.Println("respnose", sb)
		// fmt.Printf("resp %+v", response)
		youtubeData, err := processingData(response)
		if err != nil {
			time.Sleep(time.Second * 10)
			continue
		}
		fmt.Printf("data.. %+v", youtubeData)
		time.Sleep(time.Second * 10)
	}
	return nil
}

func processingData(data map[string]interface{}) (*dtos.YoutubeDataDtos, error) {
	youtubeDataDtos := dtos.YoutubeDataDtos{}
	var videos []interface{}
	if val, ok := data["items"].([]interface{}); ok {
		videos = val
	}
	for _, video := range videos {
		if videoMap, ok := video.(map[string]interface{}); ok {

			// snippet contains basic video details.
			if snippet, ok := videoMap["snippet"].(map[string]interface{}); ok {
				if title, ok := snippet["title"].(string); ok {
					youtubeDataDtos.Title = &title
				}

				if description, ok := snippet["description"].(string); ok {
					youtubeDataDtos.Description = &description
				}

				if channelId, ok := snippet["channelId"].(string); ok {
					youtubeDataDtos.ChannelId = &channelId
				}

				// may need time parsing...
				if publishTime, ok := snippet["publishTime"].(string); ok {
					parsedTime, _ := parseTime.ParseTimeFromStringToTime(publishTime)
					youtubeDataDtos.PublishedDateTime = parsedTime
				}

				if thumbnails, ok := snippet["thumbnails"].(map[string]interface{}); ok {
					if defaultThumbnail, ok := thumbnails["default"].(map[string]interface{}); ok {
						if url, ok := defaultThumbnail["url"].(string); ok {
							youtubeDataDtos.ThumbnailUrl = &url
						}
					}
				}
				//channelTitle
				if channelTitle, ok := snippet["channelTitle"].(string); ok {
					youtubeDataDtos.ChannelTitle = &channelTitle
				}

			} else {
				fmt.Println("could not parse snippet")
			}

			//video id
			if idMap, ok := videoMap["id"].(map[string]interface{}); ok {
				if videoId, ok := idMap["videoId"].(string); ok {
					youtubeDataDtos.VideoId = &videoId
				}
			}

		}
		// inserting to db
		youtubeDataModel := models.YoutubeData{
			Title:             youtubeDataDtos.Title,
			ChannelId:         youtubeDataDtos.ChannelId,
			Description:       youtubeDataDtos.Description,
			PublishedDateTime: youtubeDataDtos.PublishedDateTime,
			// PublishedDateTime: time.Now(),
			ChannelTitle: youtubeDataDtos.ChannelTitle,
			ThumbnailUrl: youtubeDataDtos.ThumbnailUrl,
			VideoId:      youtubeDataDtos.VideoId,
		}
		err := repository.CreateYoutubeData(nil, &youtubeDataModel)
		if err != nil {
			continue
		}
	}
	return &youtubeDataDtos, nil
}
