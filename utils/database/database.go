package database

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var GDB *gorm.DB

func Init() {
	dsn := "sachet:1234@tcp(127.0.0.1:3306)/pro_db?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	GDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprint("Error while connectiing to db,err", err))
	}
	fmt.Println("connected to database...")
}
