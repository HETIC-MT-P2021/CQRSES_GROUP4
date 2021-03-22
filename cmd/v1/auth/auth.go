package auth

import (
	"github.com/gin-gonic/gin"
	jwt "github.com/kyfk/gin-jwt"
)

// ApplyRoutes All routes from authentification
func ApplyRoutes(r *gin.RouterGroup, jwtAuth jwt.Auth) {
	r.Use(jwt.ErrorHandler)
	r.POST("/login", jwtAuth.Authenticate)
	r.POST("/auth/refresh_token", jwtAuth.RefreshToken)

	authRouter := r.Group("/auth")
	{
		authRouter.POST("/register", Register)
	}
}
