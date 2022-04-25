package app

import (
	"log"

	"github.com/sandriansyafridev/golang/api/book/model/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDatabase() *gorm.DB {
	dsn := "root:@tcp(localhost:3306)/golang-api-book?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Panic(err)
	}

	db.AutoMigrate(&entity.Book{})

	return db

}

func CloseDatabase(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		log.Panic(err)
	}

	sqlDB.Close()
}
