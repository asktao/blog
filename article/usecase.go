package article

import "blog/models"

type Usecase interface {
	GetArticle(id uint64) (article *models.Article, err error)
	SaveArticle(article *models.Article) (id uint64, err error)
	ListArticle(limit uint64, offset uint64) (articles []*models.Article, err error)
}