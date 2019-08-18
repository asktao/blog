package controller

import (
	"blog/article"
	"blog/logging"
	"blog/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type ArticleController struct {
	Controller
	Validation
	AUsecase article.Usecase
}

func NewArticleController(r *mux.Router, us article.Usecase){
	articleController := &ArticleController{
		AUsecase: us,
	}
	r.HandleFunc("/articles", articleController.IndexArticle).Methods("GET")
	r.HandleFunc("/articles", articleController.StoreArticle).Methods("POST")
	r.HandleFunc("/articles/{id:[0-9]+}", articleController.ShowArticle).Methods("GET")
}

func (ac *ArticleController) ShowArticle(w http.ResponseWriter, r *http.Request) {

	var err error

	defer func(){
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			logging.Error(err)
		}
	}()

	params := mux.Vars(r)

	id, _ := strconv.ParseUint(params["id"], 10, 64)

	ret, err := ac.AUsecase.GetArticle(id)

	if (models.Article{}) == ret {
		http.NotFound(w, r)
		return
	}
	if err != nil {
		return
	}

	response := Response{http.StatusOK,"Success", ret}
	ac.JsonResponse(w, response)
}

func (ac *ArticleController) IndexArticle(w http.ResponseWriter, r *http.Request) {
	limit, offset := ac.Pagination(r)

	articles, err := ac.AUsecase.ListArticle(limit, offset)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		logging.Error(err)
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
			logging.Error(err.Error())
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

	resp, err := ac.AUsecase.SaveArticle(Article)
	if err != nil {
		return
	}

	response := Response{http.StatusCreated,"Success", resp}
	ac.JsonResponse(w, response)
}