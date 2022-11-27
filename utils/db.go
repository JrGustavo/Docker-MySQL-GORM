package utils

import (
	"github.com/fuenrob/server-go/config"
	"github.com/fuenrob/server-go/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func InitConnection() *gorm.DB {
	dsn := config.GetDBConnection()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Panic("Connection error")
	}

	db.AutoMigrate(&model.User{})
	user := &model.User{Name: "Roberto"}

	db.Create(&user)

	return db
}
