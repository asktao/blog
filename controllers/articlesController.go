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
	ShowArticle(w http.ResponseWriter, r *http.Request)
	IndexArticle(w http.ResponseWriter, r *http.Request)
	StoreArticle(w http.ResponseWriter, r *http.Request)
}

func (ac *ArticleController) ShowArticle(w http.ResponseWriter, r *http.Request) {

	var err error

	defer func(){
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}()

	params := mux.Vars(r)

	id, _ := strconv.ParseUint(params["id"], 10, 64)

	article, err := article.GetArticle(id)

	if (models.Article{}) == article {
		http.NotFound(w, r)
		return
	}
	if err != nil {
		return
	}

	response := Response{http.StatusOK,"Success", article}
	ac.JsonResponse(w, response)
}

func (ac *ArticleController) IndexArticle(w http.ResponseWriter, r *http.Request) {
	limit, offset := ac.Pagination(r)

	articles, err := article.ListArticle(limit, offset)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := Response{http.StatusOK, "Success", articles}
	ac.JsonResponse(w, response)
}

func (ac *ArticleController) StoreArticle(w http.ResponseWriter, r *http.Request) {
	var err error

	defer func(){
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}()

	Article := models.Article{}
	err = json.NewDecoder(r.Body).Decode(&Article)
	if err != nil {
		return
	}

	validaErr := ac.Validator(&Article)
	if validaErr != nil {
		response := Response{http.StatusUnprocessableEntity, validaErr.Error(), nil}
		ac.JsonResponse(w, response)
		return
	}

	resp, err := article.SaveArticle(Article)
	if err != nil {
		return
	}

	response := Response{http.StatusCreated,"Success", resp}
	ac.JsonResponse(w, response)
}