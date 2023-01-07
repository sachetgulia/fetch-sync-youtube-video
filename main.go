package main

import (
	database "youtubesync/utils/database"
	fetch "youtubesync/utils/youtubefetch"

	handlers "youtubesync/handlers"

	"github.com/gin-gonic/gin"
)

func main() {

	database.Init()
	database.CreateTables()
	go fetch.FetchInit()
	r := gin.Default()
	//health check
	r.GET("/healthCheck", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "working fine...",
		})
	})
	r.GET("/search", handlers.SearchHandler)
	r.GET("/data", handlers.GetDataPaginated)
	r.Run()
	// listen and serve on 0.0.0.0:8080
}

//AIzaSyB12amtMkVda3obVW1O-U39P-8t2NMH3v4
