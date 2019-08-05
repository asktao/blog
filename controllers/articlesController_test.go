package controllers
//
//import (
//	"bytes"
//	"encoding/json"
//	"github.com/gorilla/mux"
//	"net/http"
//	"net/http/httptest"
//	"testing"
//)
//
//type TestArticle struct {
//	Id		uint64	`json:id`
//	Title   string	`json:"title"`
//	Content string	`json:"content"`
//	Author  string	`json:"author"`
//}
//
//var testArticle = TestArticle{
//	Id:			1,
//	Title:   	"Blog",
//	Content:	"Hello world",
//	Author:		"Bruce",
//}
//
//func Router() *mux.Router {
//	a := ArticleController{}
//
//	r := mux.NewRouter()
//	r.HandleFunc("/articles", a.ShowArticle).Methods("POST")
//	r.HandleFunc("/articles", a.IndexArticle).Methods("GET")
//	r.HandleFunc("/articles/1", a.ShowArticle).Methods("GET")
//	return r
//}
//
//func TestArticleController_StoreArticle(t *testing.T) {
//
//	payload, _ := json.Marshal(testArticle)
//
//	request, _ := http.NewRequest("POST", "/article", bytes.NewReader(payload))
//	response := httptest.NewRecorder()
//	Router().ServeHTTP(response, request)
//	if response.Code != http.StatusCreated {
//		t.Errorf("HTTP Response code expecting: %d Got %d", http.StatusCreated, response.Code)
//	}
//
//	//expected := `{"id": 1, "title": "Blog", "content": "Hello world", "author": "Bruce"}`
//	//if response.Body.String() != expected {
//	//	t.Errorf("HTTP Response body expecting %v Got %v", expected, response.Body.String())
//	//}
//}
//
//func TestArticleController_ShowArticle(t *testing.T) {
//	request, _ := http.NewRequest("GET", "/articles/1", nil)
//	response := httptest.NewRecorder()
//	Router().ServeHTTP(response, request)
//	if response.Code != http.StatusOK {
//		t.Errorf("HTTP Response code expecting: %d Got %d", http.StatusOK, response.Code)
//	}
//}
//
//func TestArticleController_IndexArticle(t *testing.T) {
//	request, _ := http.NewRequest("GET", "/articles", nil)
//	response := httptest.NewRecorder()
//	Router().ServeHTTP(response, request)
//	if response.Code != http.StatusOK {
//		t.Errorf("HTTP Response code expecting: %d Got %d", http.StatusOK, response.Code)
//	}
//}