package models

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("mini_order.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect database:", err)
	}

	// 自動建立資料表
	DB.AutoMigrate(&User{})
}
