package routes

import (
	"blog/controllers"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

func Run() {
	r := mux.NewRouter()
	r.HandleFunc("/articles/{id:[0-9]+}", controllers.ShowArticle).Methods("GET")
	r.HandleFunc("/articles", controllers.IndexArticle).Methods("GET")
	r.HandleFunc("/articles", controllers.StoreArticle).Methods("POST")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	err := http.ListenAndServe(":" + port, r)
	if err != nil {
		fmt.Print(err)
	}
}