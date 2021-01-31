package event

import (
	"github.com/gin-gonic/gin"
)

func ApplyRoutes(r *gin.RouterGroup) {
	/*auth, err := NewAuth()

	if err != nil {
		log.Print(err)
	}*/

	// r.Use(jwt.ErrorHandler)
	r.GET("/events", GetEvents)
	r.POST("/events", CreateEvent)

	/*authRouter := r.Group("/auth")
	{
		authRouter.POST("/register", Register)
	}*/
}
