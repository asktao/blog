package models

type Article struct {
	Id      uint64	`json:"id"`
	Title   string	`json:"title"`
	Content string	`json:"content"`
	Author  string	`json:"author"`
}

func ShowArticle(id uint64) *Article {

	article := &Article{}

	err := DB().Where("id = ?", id).First(article).Error

	if err != nil {
		//todo
	}

	return article
}

func StoreArticle(article *Article) (id uint64) {
	err := DB().Save(&article)

	if err != nil {
		//todo
	}

	return article.Id
}

func IndexArticle() []*Article {
	articles := make([]*Article, 0)

	err := DB().Find(&articles).Error

	if err != nil {
		//todo
	}

	return articles
}