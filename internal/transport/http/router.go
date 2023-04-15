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
			book.GET("/get-all", s.handler.GetAll)
			book.DELETE("/delete", s.handler.DeleteBook)
		}

		transactions := api.Group("/transactions")
		{
			transactions.GET("/get", s.handler.GetTransactions)
		}
	}
}
