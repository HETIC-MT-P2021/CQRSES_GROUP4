package auth

import (
	"github.com/gin-gonic/gin"
	jwt "github.com/kyfk/gin-jwt"
	"log"
)

func ApplyRoutes(r *gin.RouterGroup) {
	auth, err := NewAuth()

	if err != nil {
		log.Print(err)
	}

	r.Use(jwt.ErrorHandler)
	r.POST("/login", auth.Authenticate)
	r.POST("/auth/refresh_token", auth.RefreshToken)

	authRouter := r.Group("/auth")
	{
		authRouter.POST("/register", Admin(auth), Register)
	}
}
