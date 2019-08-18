package mocks

import (
	"blog/models"
	"github.com/stretchr/testify/mock"
)

type Repository struct {
	mock.Mock
}

func (_m *Repository) GetArticle(id uint64) (article *models.Article, err error) {
	ret := _m.Called(id)

	var r0 *models.Article
	if rf, ok := ret.Get(0).(func(uint64) *models.Article); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Article)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint64) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *Repository) SaveArticle(article *models.Article) (id uint64, err error) {
	ret := _m.Called(article)

	var r0 uint64
	if rf, ok := ret.Get(0).(func(*models.Article) uint64); ok {
		r0 = rf(article)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(uint64)
		}
	}

	var r1 error
	if rf, ok := ret.Get(0).(func(*models.Article) error); ok {
		r1 = rf(article)
	} else {
		r1 = ret.Error(0)
	}

	return r0, r1
}

func (_m *Repository) ListArticle(limit uint64, offset uint64) ([]*models.Article, error) {
	ret := _m.Called(limit, offset)

	var r0 []*models.Article
	if rf, ok := ret.Get(0).(func(uint64, uint64) []*models.Article); ok {
		r0 = rf(limit, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Article)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint64, uint64) error); ok {
		r1 = rf(limit, offset)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
