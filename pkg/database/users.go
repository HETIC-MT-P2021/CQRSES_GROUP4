package database

import (
	"encoding/base64"
	"github.com/jibe0123/survey/pkg/database/query"
	"golang.org/x/crypto/bcrypt"
)

type Role uint8

type RequestRegister struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

const (
	OPERATOR     Role = 0x1
	ADMIN        Role = 0x1 << 1
	SYSTEM_ADMIN Role = 0x1 << 2
)

func (r Role) IsOperator() bool {
	return r&OPERATOR != 0
}

func (r Role) IsAdmin() bool {
	return r&ADMIN != 0
}

func (r Role) IsSystemAdmin() bool {
	return r&SYSTEM_ADMIN != 0
}

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"-"` // here is just for example
	Role     Role   `json:"role"`
}

// GetUserFromUsername method for retrieve user from bdd
func GetUserFromUsername(username string) (user *User, err error) {
	var (
		Email    string
		Password string
		RoleInt  uint8
	)

	sqlStmt := query.QUERY_FIND_USERS_BY_USERNAME
	stmt, err := DbConn.Prepare(sqlStmt)
	defer stmt.Close()

	err = stmt.QueryRow(username).Scan(&Email, &Password, &RoleInt)
	if err != nil {
		return nil, err
	}

	user = &User{
		Username: username,
		Email:    Email,
		Password: Password,
		Role:     Role(RoleInt),
	}

	return user, nil
}

// CreateAccount method for create an account with role operator
func CreateAccount(userInput RequestRegister) (err error) {
	data, err := base64.StdEncoding.DecodeString(userInput.Password)
	hash, err := bcrypt.GenerateFromPassword(data, bcrypt.DefaultCost)
	sqlStmt := query.QUERY_CREATE_ACCOUNT

	tx, err := DbConn.Begin()
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
