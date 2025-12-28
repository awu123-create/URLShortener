package Config

import (
	"URLShortener/Models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	dsn := "url_shortener:Jhqlzy060103@@tcp(127.0.0.1:3306)/URL_shortener?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	db.AutoMigrate(&Models.URL{})

	return db
}
