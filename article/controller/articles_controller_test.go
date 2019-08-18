package controller

import (
	"blog/article/controller"
	"blog/article/mocks"
	"blog/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestStoreArticle(t *testing.T) {
	mockArticle := models.Article{
		Title:		"Title",
		Content:	"Content",
		Author:		"Bruce",
	}
	tempMockArticle := mockArticle
	tempMockArticle.Id = 0
	mockUCase := new(mocks.Usecase)

	j, err := json.Marshal(tempMockArticle)
	assert.NoError(t, err)

	mockUCase.On("StoreArticle", mock.Anything, mock.AnythingOfType("*models.Article")).Return(nil)

	req, err := http.NewRequest("POST", "/articles", strings.NewReader(string(j)))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json;charset=UTF=8")

	rec := httptest.NewRecorder()

	r := mux.NewRouter()

	handler := controller.ArticleController{
		AUsecase: mockUCase,
	}
	r.HandleFunc("/articles", handler.StoreArticle).Methods("POST")

	//c := e.NewContext(req, rec)
	//c.SetPath("/article")
	r.ServeHTTP(rec, req)
	require.NoError(t, err)

	assert.Equal(t, http.StatusCreated, rec.Code)
	mockUCase.AssertExpectations(t)

}

//func TestShowArticle(t *testing.T) {
//
//	request, _ := http.NewRequest("GET", "/articles/1", nil)
//	response := httptest.NewRecorder()
//	Router().ServeHTTP(response, request)
//	if response.Code != http.StatusOK {
//		t.Errorf("HTTP Response code expecting: %d Got %d", http.StatusOK, response.Code)
//	}
//}
//
//func TestIndexArticle(t *testing.T) {
//	request, _ := http.NewRequest("GET", "/articles", nil)
//	response := httptest.NewRecorder()
//	Router().ServeHTTP(response, request)
//	if response.Code != http.StatusOK {
//		t.Errorf("HTTP Response code expecting: %d Got %d", http.StatusOK, response.Code)
//	}
//}