package user

import "database/sql"

type Role uint8

type RequestRegister struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"-"` // here is just for example
	Role     Role   `json:"role"`
}

// UserRepository functions for user repository
type UserRepository interface {
	GetUserFromUsername(string) (*User, error)
	CreateAccount(*User) error
}

// UserRepositoryImpl struct for db connection
type UserRepositoryImpl struct {
	DbConn *sql.DB
}

var UserImpl *UserRepositoryImpl

// Close the DB connection
func (r *UserRepositoryImpl) Close() {
	r.DbConn.Close()
}

// NewUserRepositoryImpl creates new UserImpl
func NewUserRepositoryImpl(dbConn *sql.DB) {
	UserImpl = &UserRepositoryImpl{
		DbConn: dbConn,
	}
}
