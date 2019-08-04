package models

type Article struct {
	Id      uint64	`json:"id"`
	Title   string	`json:"title"`
	Content string	`json:"content"`
	Author  string	`json:"author"`
}

type ArticleInterface interface {
	SHowArticle(id uint64) (Article, err error)
	StoreArticle(article *Article) (id uint64, err error)
	IndexArticle() (articles []Article, err error)
}

func (A Article) ShowArticle(id uint64) (article Article, err error) {

	article = Article{}

	if err := DB().Where("id = ? ", id).First(&article).Error; err != nil {
		return article, err
	}

	return article, nil
}

func (A Article) StoreArticle(article Article) (id uint64, err error) {

	if err := DB().Save(&article).Error; err != nil {
		return id, err
	}

	return article.Id, nil
}

func (A Article) IndexArticle() (articles []Article, err error) {

	articles = []Article{}

	if err := DB().Find(&articles).Error; err != nil {
		return articles, err
	}

	return articles, nil
}