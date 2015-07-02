package db

import (
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

type connection struct {
	connection gorm.DB
}

var conn *connection

func GetConnection() gorm.DB {
	if conn == nil {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}

		db, _ := gorm.Open("mysql",
			os.Getenv("DB_USER")+":"+os.Getenv("DB_PASS")+"@/"+os.Getenv("DB_NAME")+"?charset=utf8&parseTime=True&loc=Local")
		conn = &connection{connection: db}
	}
	return conn.connection
}

func CloseConnection() {
	conn.connection.Close()
}
