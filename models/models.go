package models

import (
	"blog/logging"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var db *gorm.DB

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	host		:=	os.Getenv("DB_HOST")
	port		:=	os.Getenv("DB_PORT")
	database	:=	os.Getenv("DB_DATABASE")
	username	:=	os.Getenv("DB_USERNAME")
	password	:=	os.Getenv("DB_PASSWORD")

	dbEnv := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, database)
	conn, err := gorm.Open("mysql", dbEnv)
	if err != nil {
		logging.Error(err.Error())
		log.Fatal("Failed to establish a MySQL link")
	}

	if os.Getenv("APP_DEBUG") == "True" {
		conn.LogMode(true)
	}
	db = conn
}

func DB() *gorm.DB {
	return db
}