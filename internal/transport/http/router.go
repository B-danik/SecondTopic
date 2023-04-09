package http

func (s *Server) SetupRoutes() {

	auth := s.App.Group("/auth")
	{
		auth.POST("/sign-up", s.handler.CreateUser)
		auth.POST("/sign-in", s.handler.CreateUser1)
	}

}
