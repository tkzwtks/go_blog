package models

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

func Migrate() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error")
	}

	db, _ := gorm.Open("mysql",
		os.Getenv("DB_USER")+":"+os.Getenv("DB_PASS")+"@/"+os.Getenv("DB_NAME")+"?charset=utf8&parseTime=True&loc=Local")

	db.CreateTable(&Article{})
	db.CreateTable(&Photo{})
	db.CreateTable(&Album{})
	db.CreateTable(&Game{})

	db.AutoMigrate(&Article{}, &Photo{}, &Album{}, &Game{})

}
