package controllers

import (
	"blog/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func ShowArticle(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.ParseUint(params["id"], 10, 64)
	article := models.ShowArticle(id)

	response := Message(200, "Success")
	response["data"] = article

	JsonResponse(w, 200, response)
}

func IndexArticle(w http.ResponseWriter, r *http.Request) {
	//todo pagination
	articles := models.IndexArticle()

	response := Message(200, "Success")
	response["data"] = articles

	JsonResponse(w, 200, response)
}

func StoreArticle(w http.ResponseWriter, r *http.Request) {
	article := &models.Article{}
	err := json.NewDecoder(r.Body).Decode(article)
	if err != nil {

	}
	resp := models.StoreArticle(article)

	response := Message(200, "Success")
	response["data"] = resp

	JsonResponse(w, 201, response)
}