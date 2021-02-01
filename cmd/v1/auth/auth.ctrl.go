package auth

import (
	"net/http"

	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/pkg/database"
	"github.com/gin-gonic/gin"
	jwt "github.com/kyfk/gin-jwt"
	"golang.org/x/crypto/bcrypt"
)

type User = database.User
type Role = database.Role

type requestRegister struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

// NewAuth Create new auth context
func NewAuth() (jwt.Auth, error) {
	return jwt.New(jwt.Auth{
		SecretKey: []byte("S3CR3TK3Y733T"),
		Authenticator: func(c *gin.Context) (jwt.MapClaims, error) {
			var req struct {
				Username string `json:"username"`
				Password string `json:"password"`
			}

			if err := c.ShouldBind(&req); err != nil {
				return nil, jwt.ErrorAuthenticationFailed
			}

			u, err := database.GetUserFromUsername(req.Username)

			if err != nil {
				return nil, jwt.ErrorUserNotFound
			}

			if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(req.Password)); err != nil {
				return nil, jwt.ErrorAuthenticationFailed
			}

			return jwt.MapClaims{
				"username": u.Username,
				"role":     u.Role,
			}, nil
		},
		UserFetcher: func(c *gin.Context, claims jwt.MapClaims) (interface{}, error) {
			username, ok := claims["username"].(string)
			if !ok {
				return nil, nil
			}
			u, err := database.GetUserFromUsername(username)
			if err != nil {
				return nil, nil
			}
			return u, nil
		},
	})
}

// Operator is user is admin
func Operator(m jwt.Auth) gin.HandlerFunc {
	return m.VerifyPerm(func(claims jwt.MapClaims) bool {
		return role(claims).IsOperator()
	})
}

// Admin is user is admin
func Admin(m jwt.Auth) gin.HandlerFunc {
	return m.VerifyPerm(func(claims jwt.MapClaims) bool {
		return role(claims).IsAdmin()
	})
}

// SystemAdmin is user is system admin
func SystemAdmin(m jwt.Auth) gin.HandlerFunc {
	return m.VerifyPerm(func(claims jwt.MapClaims) bool {
		return role(claims).IsSystemAdmin()
	})
}

// role get the role of the user
func role(claims jwt.MapClaims) Role {
	return Role(claims["role"].(float64))
}

// Register routes for creating account
func Register(c *gin.Context) {
	var req database.RequestRegister

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	err := database.CreateAccount(req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Account_created",
	})
}
