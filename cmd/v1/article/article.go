package article

func ApplyRoutes(r *gin.RouterGroup) {
	/*auth, err := NewAuth()

	if err != nil {
		log.Print(err)
	}*/

	// r.Use(jwt.ErrorHandler)
	r.GET("/articles", GetArticles)

	/*authRouter := r.Group("/auth")
	{
		authRouter.POST("/register", Register)
	}*/
}