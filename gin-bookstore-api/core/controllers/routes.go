package controllers

func (s *Server) endpoints() {
	bookRoute(s)
	authorRoute(s)
	customerRoute(s)
	reviewRoute(s)
}
func bookRoute(s *Server) {

	s.gin.POST("/book", s.CreateBook)
	s.gin.GET("/", s.ListBooks)
	s.gin.GET("/book/:id", s.GetBookById)
	s.gin.PUT("/book/:id", s.UpdateBook)
	s.gin.DELETE("/book/:id", s.DeleteBook)
}
func authorRoute(s *Server) {

}
func customerRoute(s *Server) {

}
func reviewRoute(s *Server) {

}
