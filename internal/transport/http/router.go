package http

func (s *Server) SetupRoutes() {

	auth := s.App.Group("/auth")
	{
		auth.POST("/sign-up", s.handler.SignUp)
		auth.POST("/sign-in", s.handler.SignIn)
	}

	api := s.App.Group("/api", s.handler.UserIdentity)
	{
		book := api.Group("/books")
		{
			book.POST("/create", s.handler.CreateBook)
			book.GET("/get-book", s.handler.GetBook)
			book.GET("/get-list", s.handler.GetList)
			book.DELETE("/delete", s.handler.DeleteBook)
		}

		rent := api.Group("/rent")
		{
			rent.GET("/get", s.handler.GetRent)
			rent.POST("/create", s.handler.CreateRent)
		}
	}
}
