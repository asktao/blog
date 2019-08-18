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

func (au *articleUsecase) GetArticle(id uint64) (*models.Article, error) {

	res, err := au.articleRepository.GetArticle(id)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (au *articleUsecase) ListArticle(limit uint64, offset uint64) ([]*models.Article, error) {

	res, err := au.articleRepository.ListArticle(limit, offset)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (au *articleUsecase) SaveArticle(m *models.Article) (id uint64, err error) {

	res, err := au.articleRepository.SaveArticle(m)

	if err != nil {
		return res, err
	}

	return res, err
}