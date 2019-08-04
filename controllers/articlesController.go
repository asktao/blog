package controllers

import (
	"blog/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

var article models.Article = models.Article{}

type ArticleController struct {
	Controller
	Validation
}

type ArticleInterface interface{
	ControllerInterface
	ValidationInterface
	models.ArticleInterface
}

func (ac *ArticleController) ShowArticle(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, _ := strconv.ParseUint(params["id"], 10, 64)

	article, err := article.ShowArticle(id)

	if article.Id == 0 {
		http.NotFound(w, r)
	}
	if err != nil {
		http.Error(w, "Query failed", http.StatusInternalServerError)
	}

	response := Response{http.StatusOK,"Success", article}
	ac.JsonResponse(w, http.StatusOK, response)
}

func (ac *ArticleController) IndexArticle(w http.ResponseWriter, r *http.Request) {
	//todo pagination
	articles, err := article.IndexArticle()
	if err != nil {
		http.Error(w, "Query failed", http.StatusInternalServerError)
	}

	response := Response{http.StatusOK, "Success", articles}
	ac.JsonResponse(w, http.StatusOK, response)
}

func (ac *ArticleController) StoreArticle(w http.ResponseWriter, r *http.Request) {
	//todo validate
	//body, _ := ioutil.ReadAll(r.Body)

	//var request map[string]interface{}
	//
	//err := json.Unmarshal(body, &request)
	//
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//}

	//ret, message := Validator(request,"title", "string")
	//
	//if (ret != 0) || (message != "") {
	//	response := Response{http.StatusUnprocessableEntity,"Unprocessable Entity" ,message}
	//	JsonResponse(w, http.StatusUnprocessableEntity, response)
	//	return
	//}

	if err := json.NewDecoder(r.Body).Decode(&article); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	resp, err := article.StoreArticle(article)
	if err != nil {
		http.Error(w, "Failed to write data", http.StatusInternalServerError)
	}

	response := Response{http.StatusOK,"Success", resp}
	ac.JsonResponse(w, http.StatusCreated, response)
}