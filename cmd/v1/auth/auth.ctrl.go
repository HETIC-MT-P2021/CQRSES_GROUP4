package auth

import (
	"net/http"

	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/pkg/database/user"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Register route for creating account
// @Summary Create new account
// @Description Using JWT auth
// @Tags auth
// @Accept  json
// @Produce  json
// @Param body body user.requestRegister true "Add account"
// @Success 200 {object} pkg.HTTPStatus "Status"
// @Failure 500 {object} pkg.HTTPError "Error"
// @Router /register [post]
func Register(c *gin.Context) {
	var req user.RequestRegister

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	req.Password = string(hash)
	
	err = user.UserImpl.CreateAccount(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": "Account_created",
	})
}

// Login An account
// @Summary Connect user to app
// @Description Using JWT auth (look headers for token)
// @Tags auth
// @Accept  json
// @Produce  json
// @Param body body user.requestLogin true "Account to login"
// @Success 200 {string} string "Empty"
// @Failure 500 {object} pkg.HTTPError "Error"
// @Router /login [post]
func Login(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{
		"message": "Go in ./auth.go to see which jwt function is used",
	})
}
