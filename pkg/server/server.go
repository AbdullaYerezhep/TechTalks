package server

import "net/http"

type Server struct {
	server *http.Server
}

func (s *Server) Run(port string, handler *http.ServeMux) error {
	s.server = &http.Server{
		Addr:    ":" + port,
		Handler: handler,
	}
	return s.server.ListenAndServe()
}
