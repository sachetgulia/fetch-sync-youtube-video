package database

import (
	"fmt"
	youtubeDataModel "youtubesync/models"
)

func CreateTables() {
	databaseName := GDB.Migrator().CurrentDatabase()
	fmt.Println("current database", databaseName)

	err := GDB.Table("youtube_data").AutoMigrate(youtubeDataModel.YoutubeData{})
	if err != nil {
		fmt.Println("error while creating table , err", err)
	}
	fmt.Println("table created successfully..")
}
