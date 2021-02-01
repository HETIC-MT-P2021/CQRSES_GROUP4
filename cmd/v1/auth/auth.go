package auth

import (
	"log"

	"github.com/gin-gonic/gin"
	jwt "github.com/kyfk/gin-jwt"
)

var Auth jwt.Auth

func ApplyRoutes(r *gin.RouterGroup) {
	Auth, err := NewAuth()
	if err != nil {
		log.Println(err)
	}

	r.Use(jwt.ErrorHandler)
	r.POST("/login", Auth.Authenticate)
	r.POST("/auth/refresh_token", Auth.RefreshToken)

	authRouter := r.Group("/auth")
	{
		authRouter.POST("/register", Register)
	}
}
