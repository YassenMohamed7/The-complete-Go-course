package server

import "time"

type Server struct {
	Addr           string
	Timeout        time.Duration
	MaxConnections int
}

func New(addr string, options ...option) *Server {
	s := &Server{
		Addr: addr,
	}

	for _, option := range options {
		option(s)
	}
	return s
}

type option func(*Server)

func WithTimeout(timeout time.Duration) option {
	return func(s *Server) {
		s.Timeout = timeout
	}
}

func WithMaxConnections(maxConnections int) option {
	return func(s *Server) {
		s.MaxConnections = maxConnections
	}
}
