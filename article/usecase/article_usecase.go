package usecase

import (
	"blog/article"
	"blog/models"
)

type articleUsecase struct {
	articleRepository    article.Repository
}

func NewArticleUsecase (ar article.Repository) article.Usecase {
	return &articleUsecase{
		articleRepository: ar,
	}
}

func (au *articleUsecase) ShowArticle(id int64) (*models.Article, error) {

	res, err := au.articleRepository.GetArticle(id)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (au *articleUsecase) IndexArticle(limit int64, offset int64) ([]*models.Article, error) {

	res, err := au.articleRepository.ListArticle(limit, offset)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (au *articleUsecase) StoreArticle(m *models.Article) (*models.Article, error) {

	res, err := au.articleRepository.SaveArticle(m)

	if err != nil {
		return nil, err
	}

	return res, err
}