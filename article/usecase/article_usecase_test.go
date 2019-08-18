package usecase

import (
	"blog/article/mocks"
	"blog/article/usecase"
	"blog/models"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestListArticle(t *testing.T) {
	mockArticleRepo := new(mocks.Repository)
	mockArticle := &models.Article{
		Title:		"Hello",
		Content:	"Content",
		Author:		"Bruce",
	}

	mockListArticle := make([]*models.Article, 0)
	mockListArticle = append(mockListArticle, mockArticle)

	t.Run("success", func(t *testing.T) {
		mockArticleRepo.On("ListArticle", mock.AnythingOfType("uint64"),
			mock.AnythingOfType("uint64")).Return(mockListArticle, nil).Once()

		u := usecase.NewArticleUsecase(mockArticleRepo)
		limit := uint64(10)
		offset := uint64(0)
		list, err := u.ListArticle(limit, offset)
		assert.NoError(t, err)
		assert.Equal(t, list, mockListArticle)

		mockArticleRepo.AssertExpectations(t)
	})

	t.Run("error-failed", func(t *testing.T) {
		mockArticleRepo.On("ListArticle", mock.AnythingOfType("uint64"),
			mock.AnythingOfType("uint64")).Return(nil, errors.New("Unexpexted Error")).Once()

		u := usecase.NewArticleUsecase(mockArticleRepo)
		limit := uint64(10)
		offset := uint64(0)
		list, err := u.ListArticle(limit, offset)

		assert.Error(t, err)
		assert.Nil(t, list)
		mockArticleRepo.AssertExpectations(t)
	})
}

func TestGetArticle(t *testing.T) {
	mockArticleRepo := new(mocks.Repository)
	mockArticle := &models.Article{
		Id:			0,
		Title:		"Hello",
		Content:	"Content",
		Author:		"Bruce",
	}

	t.Run("success", func(t *testing.T) {
		mockArticleRepo.On("GetArticle", mock.AnythingOfType("uint64")).Return(mockArticle, nil).Once()
		u := usecase.NewArticleUsecase(mockArticleRepo)

		a, err := u.GetArticle(mockArticle.Id)

		assert.NoError(t, err)
		assert.NotNil(t, a)

		mockArticleRepo.AssertExpectations(t)
	})
	t.Run("error-failed", func(t *testing.T) {
		mockArticleRepo.On("GetArticle", mock.AnythingOfType("uint64")).Return(nil, errors.New("Unexpected")).Once()

		u := usecase.NewArticleUsecase(mockArticleRepo)

		a, err := u.GetArticle(mockArticle.Id)

		assert.Error(t, err)
		assert.Nil(t, a)

		mockArticleRepo.AssertExpectations(t)
	})
}