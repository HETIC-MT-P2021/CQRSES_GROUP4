package auth

import (
	"net/http"

	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/pkg/database/user"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type requestRegister struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

// Register routes for creating account
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
		"message": "Account_created",
	})
}
