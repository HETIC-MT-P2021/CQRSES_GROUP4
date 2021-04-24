package pkg

import (
	"database/sql"
	"log"
	"net/http"
	"net/http/httptest"

	"github.com/DATA-DOG/go-sqlmock"
)

func NewSQLMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

func NewHandlerMock(response string) *httptest.Server{
	handler := http.NotFound
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler(w, r)
	}))

	handler = func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(response))
	}

	return server
}