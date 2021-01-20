package article

import (
	"github.com/gin-gonic/gin"
)

func ApplyRoutes(r *gin.RouterGroup) {
	/*auth, err := NewAuth()

	if err != nil {
		log.Print(err)
	}*/

	// r.Use(jwt.ErrorHandler)
	r.GET("/articles", GetArticle)
	r.POST("/articles", CreateArticle)

	/*authRouter := r.Group("/auth")
	{
		authRouter.POST("/register", Register)
	}*/
}
