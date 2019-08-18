package repository

import (
	"blog/article/repository"
	"blog/models"
	"context"
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
	//repository repository.ArticleRepository
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

func (s *Suite) AfterTest(_, _ string) {
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}
//
//func TestInit(t *testing.T) {
//	suite.Run(t, new(Suite))
//}

func (s *Suite) TestGetArticle(t *testing.T) {
	s.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "articles" WHERE (id = $1)`)).
		WithArgs(id.String()).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).
			AddRow(id.String(), name))


	anArticle, err := r.GetArticle(context.TODO())
	assert.NoError(t, err)
	assert.NotNil(t, anArticle)
}

func TestSaveArticle(t *testing.T) {

}