package user

import (
	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/pkg/database/query"
)

// GetUserFromUsername method for retrieve user from bdd
func (r *UserRepositoryImpl) GetUserFromUsername(username string) (*User, error) {
	var (
		Email    string
		Password string
		RoleInt  uint8
	)

	stmt, err := r.DbConn.Prepare(query.QUERY_FIND_USERS_BY_USERNAME)
	if err != nil {
		return nil, err
	}
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
	tx, err := r.DbConn.Begin()
	if err != nil {
		return err
	}

	stmt, es := tx.Prepare(query.QUERY_CREATE_ACCOUNT)
	if es != nil {
		return es
	}
	defer stmt.Close()

	if _, err := stmt.Exec(userInput.Email, userInput.Username, userInput.Password); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}