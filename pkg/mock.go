package pkg

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/DATA-DOG/go-sqlmock"
	elastic "github.com/olivere/elastic/v7"
)

// NewSQLMock Creates fake sql connection
func NewSQLMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

// NewSQLMock Creates fake http handler
func NewHandlerMock() http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("success"))
	}
}

// NewSQLMock Creates fake elastic client
func NewElasticClientMock(url string) (*elastic.Client, error) {
	client, err := elastic.NewSimpleClient(elastic.SetURL(url))
	if err != nil {
		return nil, err
	}
	return client, nil
}