package auth

import (
	"net/http"

	"github.com/HETIC-MT-P2021/CQRSES_GROUP4/pkg/database"
	"github.com/gin-gonic/gin"
)

type requestRegister struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
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
