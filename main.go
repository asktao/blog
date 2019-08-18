package main

import (
	"blog/article/controller"
	"blog/article/repository"
	"blog/article/usecase"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {

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
		log.Fatal("Failed to establish a MySQL link")
	}

	if os.Getenv("APP_DEBUG") == "True" {
		conn.LogMode(true)
	}

	defer func() {
		err := conn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	r := mux.NewRouter()
	ar := repository.NewArticleRepository(conn)
	au := usecase.NewArticleUsecase(ar)

	controller.NewArticleController(r, au)

	port = os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	err = http.ListenAndServe(":"+port, r)
	if err != nil {
		log.Fatal(err)
	}
}
