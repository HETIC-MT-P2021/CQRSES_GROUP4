package user

import (
	"testing"
)

func TestCreateAccount(t *testing.T) {
	/*userInput := RequestRegister{
		Username: "admin",
		Password: "admin",
		Email: "admin@gmail.com",
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)
	sqlStmt := query.QUERY_CREATE_ACCOUNT

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectBegin()
	mock.ExpectRollback()
	mock.ExpectExec(sqlStmt).
		WithArgs(userInput.Email, userInput.Username, hash).
		WillReturnResult(sqlmock.NewResult(1, 1))
	
	mock.ExpectCommit()*/
	/*stmt, es := tx.Prepare(sqlStmt)
	if es != nil {
		return es
	}
	defer stmt.Close()

	if _, err := stmt.Exec(userInput.Email, userInput.Username, hash); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
	
	}*/
}
