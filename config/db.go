package config

import (
	"aldinoh8/example-gin/entity"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDb(dsn string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	db.AutoMigrate(&entity.Todo{})
	return db
}
