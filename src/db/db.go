package db

import (
	"gorm.io/gorm"
    "gorm.io/driver/postgres"
	"log"
)

type Messages struct {
	gorm.Model
	Id    int `gorm:"primaryKey"`
	Text string 
	Processed bool `gorm:"default:false"`
}

func Engine() **gorm.DB {
	dsn := "user=postgres password=secret database=postgres host=postgres port=5432 sslmode=disable TimeZone=Europe/Moscow"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
        log.Fatal(err)
    }
	return &db
}