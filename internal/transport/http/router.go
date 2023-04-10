package http

func (s *Server) SetupRoutes() {

	auth := s.App.Group("/auth")
	{
		auth.POST("/sign-up", s.handler.SignUp)
		auth.POST("/sign-in", s.handler.SignIn)
	}

	api := s.App.Group("/api")
	{
		list := api.Group("/lists")
		{
			list.POST("/", s.handler.UserIdentity)
		}
	}

}
