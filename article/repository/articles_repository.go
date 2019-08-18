package repository

import (
	"blog/article"
	"blog/models"
	"github.com/jinzhu/gorm"
)

type ArticleRepositoryInterface interface {
	GetArticle(id uint64) (article *models.Article, err error)
	SaveArticle(article *models.Article) (id uint64, err error)
	ListArticle(limit uint64, offset uint64) (articles []*models.Article, err error)
}

type ArticleRepository struct {
	DB *gorm.DB
}

func NewArticleRepository(DB *gorm.DB) article.Repository {
	return &ArticleRepository{DB}
}

func (A ArticleRepository) GetArticle(id uint64) (article *models.Article, err error) {

	if err := A.DB.Where("id = ? ", id).First(&article).Error; err != nil {
		return article, err
	}

	return article, nil
}

func (A ArticleRepository) SaveArticle(article *models.Article) (id uint64, err error) {

	if err := A.DB.Save(&article).Error; err != nil {
		return id, err
	}

	return article.Id, nil
}

func (A ArticleRepository) ListArticle(limit uint64, offset uint64) (articles []*models.Article, err error) {

	if err := A.DB.Limit(limit).Offset(offset).Find(&articles).Error; err != nil {
		return articles, err
	}

	return articles, nil
}