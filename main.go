package main

import (
	database "youtubesync/utils/database"
	fetch "youtubesync/utils/youtubefetch"

	search "youtubesync/handlers"

	"github.com/gin-gonic/gin"
)

func main() {

	database.Init()
	database.CreateTables()
	go fetch.FetchInit()
	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "nine",
		})
	})
	r.GET("/search", search.SearchHandler)
	r.Run()
	// listen and serve on 0.0.0.0:8080
}

//AIzaSyB12amtMkVda3obVW1O-U39P-8t2NMH3v4
