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
	ShowArticle(w http.ResponseWriter, r *http.Request)
	IndexArticle(w http.ResponseWriter, r *http.Request)
	StoreArticle(w http.ResponseWriter, r *http.Request)
}

func (ac *ArticleController) ShowArticle(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, _ := strconv.ParseUint(params["id"], 10, 64)

	article, err := article.GetArticle(id)

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
	limit, offset := ac.Pagination(r)

	articles, err := article.ListArticle(limit, offset)
	if err != nil {
		http.Error(w, "Query failed", http.StatusInternalServerError)
	}

	response := Response{http.StatusOK, "Success", articles}
	ac.JsonResponse(w, http.StatusOK, response)
}

func (ac *ArticleController) StoreArticle(w http.ResponseWriter, r *http.Request) {

	Article := models.Article{}
	err := json.NewDecoder(r.Body).Decode(&Article)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	type ValidateArticle struct {
		Title   string	`valid:"required" json"title"`
		Content string	`valid:"required" json"content"`
		Author  string	`valid:"required" json:"author"`
	}

	validateArticle := &ValidateArticle{Article.Title, Article.Content, Article.Author}

	err = ac.Validator(validateArticle)
	if err != nil {
		response := Response{http.StatusUnprocessableEntity, err.Error(), nil}
		ac.JsonResponse(w, http.StatusUnprocessableEntity, response)
		return
	}

	resp, err := article.SaveArticle(Article)
	if err != nil {
		http.Error(w, "Failed to write data", http.StatusInternalServerError)
	}

	response := Response{http.StatusOK,"Success", resp}
	ac.JsonResponse(w, http.StatusCreated, response)
}