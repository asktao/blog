package routes

import (
	"blog/logging"
	"blog/controllers"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

var articleController = controllers.ArticleController{}

func Run() {
	r := mux.NewRouter()
	r.HandleFunc("/articles/{id:[0-9]+}", articleController.ShowArticle).Methods("GET")
	r.HandleFunc("/articles", articleController.IndexArticle).Methods("GET")
	r.HandleFunc("/articles", articleController.StoreArticle).Methods("POST")

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	err := http.ListenAndServe(":" + port, r)
	if err != nil {
		logging.Error(err.Error())
		fmt.Print(err)
	}
}