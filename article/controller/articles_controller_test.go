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
	"strconv"
	"strings"
	"testing"
)

func TestSaveArticle(t *testing.T) {
	mockArticle := models.Article{
		Id:			0,
		Title:		"Title",
		Content:	"Content",
		Author:		"Bruce",
	}
	tempMockArticle := mockArticle
	mockUCase := new(mocks.Usecase)

	j, err := json.Marshal(tempMockArticle)
	assert.NoError(t, err)

	mockUCase.On("SaveArticle", mock.AnythingOfType("*models.Article")).Return(nil)

	req, err := http.NewRequest("POST", "/articles", strings.NewReader(string(j)))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json;charset=UTF=8")

	rec := httptest.NewRecorder()

	r := mux.NewRouter()

	handler := controller.ArticleController{
		AUsecase: mockUCase,
	}
	r.HandleFunc("/articles", handler.SaveArticle).Methods("POST")

	r.ServeHTTP(rec, req)
	require.NoError(t, err)

	assert.Equal(t, http.StatusCreated, rec.Code)
	mockUCase.AssertExpectations(t)

}

func TestGetArticle(t *testing.T) {

	mockArticle := models.Article{
		Id:			0,
		Title:		"Title",
		Content:	"Content",
		Author:		"Bruce",
	}
	mockUCase := new(mocks.Usecase)

	id := int(mockArticle.Id)

	mockUCase.On("GetArticle", uint64(id)).Return(&mockArticle, nil)

	req, err := http.NewRequest("GET", "/articles/" + strconv.Itoa(id), strings.NewReader(string("")))

	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	r := mux.NewRouter()
	handler := controller.ArticleController{
		AUsecase: mockUCase,
	}
	r.HandleFunc("/articles/" + strconv.Itoa(id), handler.GetArticle).Methods("GET")
	r.ServeHTTP(rec, req)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockUCase.AssertExpectations(t)
}

func TestListArticle(t *testing.T) {
	mockArticle := models.Article{
		Id:			0,
		Title:		"Title",
		Content:	"Content",
		Author:		"Bruce",
	}
	mockUCase := new(mocks.Usecase)
	mockListArticle := make([]*models.Article, 0)
	mockListArticle = append(mockListArticle, &mockArticle)
	var limit uint64 = 10
	var offset uint64 = 0
	mockUCase.On("ListArticle", limit, offset).Return(mockListArticle, nil)

	req, err := http.NewRequest("GET", "/articles?limit=10&offset=0", strings.NewReader(""))
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	r := mux.NewRouter()
	handler := controller.ArticleController{
		AUsecase: mockUCase,
	}
	r.HandleFunc("/articles", handler.ListArticle).Methods("GET")
	r.ServeHTTP(rec, req)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rec.Code)
	mockUCase.AssertExpectations(t)
}