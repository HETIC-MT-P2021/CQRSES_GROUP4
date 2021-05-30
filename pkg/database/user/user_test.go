package user

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/pkg"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/pkg/database/query"
	"github.com/stretchr/testify/assert"
)

const (
	DEFAULT_USERNAME = "admin"
	DEFAULT_PASSWORD = "admin"
	DEFAULT_EMAIL = "admin@gmail.com"
	DEFAULT_ROLE = 1
	
	STATUS_SUCCESS = "success" 
	STATUS_ERROR = "error" 
)

func TestCreateAccount(t *testing.T) {
	db, mock := pkg.NewSQLMock()
	repo := UserRepositoryImpl{
		DbConn: db,
	}

	defer func() {
		repo.Close()
	}()

	userInput := RequestRegister{
		Username: DEFAULT_USERNAME,
		Password: DEFAULT_PASSWORD,
		Email: DEFAULT_EMAIL,
	}

	var cases = []struct {
		what        		string // What I want to test
		userInput        RequestRegister // Input
		status        	string // status, success | error
	}{
		{"Ok", userInput, STATUS_SUCCESS},
	}

	for _, tt := range cases {
		mock.ExpectBegin()

		prep := mock.ExpectPrepare(regexp.QuoteMeta(query.QUERY_CREATE_ACCOUNT))
		prep.ExpectExec().
			WithArgs(tt.userInput.Email, tt.userInput.Username, tt.userInput.Password).
			WillReturnResult(sqlmock.NewResult(0, 1))

		mock.ExpectCommit()

		err := repo.CreateAccount(tt.userInput)

		if tt.status == "success" {
			assert.NoError(t, err)
		} else {
			assert.Error(t, err)
		}
	}
}

func TestGetUserFromUsername(t *testing.T) {
	db, mock := pkg.NewSQLMock()
	repo := UserRepositoryImpl{
		DbConn: db,
	}

	defer func() {
		repo.Close()
	}()

	user := User {
		Password: DEFAULT_PASSWORD,
		Email: DEFAULT_EMAIL,
		Role: DEFAULT_ROLE,
	}

	var cases = []struct {
		what        		string // What I want to test
		username        string // Input
		status        	string // status, success | error
	}{
		{"Ok", "admin", STATUS_SUCCESS},
		{"User not found", "tt", STATUS_ERROR},
	}

	for _, tt := range cases {
		rows := sqlmock.NewRows([]string{"Email", "Password", "Role"}).
		AddRow(user.Email, user.Password, user.Role)
	
		prep := mock.ExpectPrepare(regexp.QuoteMeta(query.QUERY_FIND_USERS_BY_USERNAME))
		prep.ExpectQuery().
			WithArgs(DEFAULT_USERNAME).
			WillReturnRows(rows)

		_, err := repo.GetUserFromUsername(tt.username)

		if (tt.status == "success") {
			assert.NoError(t, err)
		} else {
			assert.Error(t, err)
		}
	}
}
