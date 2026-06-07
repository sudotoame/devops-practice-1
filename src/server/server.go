package server

import "net/http"

type Server struct {
	Handlers *Handlers
	Server   *http.Server
}

func NewServer(handlers *Handlers) *Server {
	router := http.NewServeMux()
	router.HandleFunc("POST /number", handlers.HandleIncrease)
	router.HandleFunc("GET /number", handlers.HandleOutput)
	router.HandleFunc("GET /healthcheck", handlers.HanldeHealthcheck)

	return &Server{
		Handlers: handlers,
		Server: &http.Server{
			Addr:    ":9091",
			Handler: router,
		},
	}
}

func (s *Server) Run() {
	s.Server.ListenAndServe()
}
