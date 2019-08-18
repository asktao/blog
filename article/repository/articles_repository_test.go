package repository

import (
	"blog/article/repository"
	"blog/models"
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"regexp"
	"testing"
)

type Suite struct {
	suite.Suite
	DB	*gorm.DB
	mock sqlmock.Sqlmock
	repository repository.ArticleRepositoryInterface
	article *models.Article
}

func (s *Suite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)

	s.DB, err = gorm.Open("mysql", db)
	require.NoError(s.T(), err)

	s.DB.LogMode(true)

	s.repository = repository.NewArticleRepository(s.DB)
}

//func (s *Suite) AfterTest(_, _ string) {
//	require.NoError(s.T(), s.mock.ExpectationsWereMet())
//}

func (s *Suite) TestSaveArticle(t *testing.T) {

	mockArticle := &models.Article{
		Id:			1,
		Title:		"Title",
		Content:	"Content",
		Author:		"Bruce",
	}

	s.mock.ExpectQuery(regexp.QuoteMeta(
		`INSERT INTO "articles" ("id", "title", "content", "author") 
			VALUES ($1,$2) RETURNING "article"."id"`)).
		WithArgs(mockArticle).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "title", "content", "author"}).
				AddRow(mockArticle.Id, mockArticle.Title, mockArticle.Content, mockArticle.Author))

	id, err := s.repository.SaveArticle(mockArticle)

	assert.Equal(t, id, mockArticle.Id)
	require.NoError(s.T(), err)
}

func (s *Suite) TestGetArticle(t *testing.T) {

	mockArticle := &models.Article{
		Id:			1,
		Title:		"Title",
		Content:	"Content",
		Author:		"Bruce",
	}

	s.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "articles" WHERE (id = $1)`)).
		WithArgs(mockArticle).
		WillReturnRows(sqlmock.NewRows([]string{"id", "title", "content", "author"}).
			AddRow(mockArticle.Id))

	anArticle, err := s.repository.GetArticle(mockArticle.Id)
	assert.NoError(t, err)
	assert.NotNil(t, anArticle)
	assert.Equal(t, mockArticle, anArticle)
}

func (s *Suite) TestListArticle(t *testing.T) {
	mockArticles := []models.Article{
		{
			Id:      1,
			Title:   "Title",
			Content: "Content",
			Author:  "Bruce",
		},
		{
			Id:      2,
			Title:   "Title",
			Content: "Content",
			Author:  "Bruce",
		},
	}

	s.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "articles"`)).
		WithArgs(mockArticles).
		WillReturnRows(sqlmock.NewRows([]string{"id", "title", "content", "author"}).
			AddRow(mockArticles))

	anArticle, err := s.repository.ListArticle(10, 0)
	assert.NoError(t, err)
	assert.NotNil(t, anArticle)
	assert.Equal(t, mockArticles, anArticle)
}