package user

import (
	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/pkg/database/query"
	"golang.org/x/crypto/bcrypt"
)

// GetUserFromUsername method for retrieve user from bdd
func (r *UserRepositoryImpl) GetUserFromUsername(username string) (*User, error) {
	var (
		Email    string
		Password string
		RoleInt  uint8
	)

	sqlStmt := query.QUERY_FIND_USERS_BY_USERNAME
	stmt, err := r.DbConn.Prepare(sqlStmt)
	defer stmt.Close()

	err = stmt.QueryRow(username).Scan(&Email, &Password, &RoleInt)
	if err != nil {
		return nil, err
	}

	user := &User{
		Username: username,
		Email:    Email,
		Password: Password,
		Role:     Role(RoleInt),
	}

	return user, nil
}

// CreateAccount method for create an account with role operator
func (r *UserRepositoryImpl) CreateAccount(userInput RequestRegister) (err error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)
	sqlStmt := query.QUERY_CREATE_ACCOUNT

	tx, err := r.DbConn.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmt, es := tx.Prepare(sqlStmt)
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
}