package server

import (
	"time"
)

type Opts struct {
	ID            uint64
	MaxConnection uint64
	Tls           bool
	Status        string
	UpSince       time.Time
}

type Server struct {
	Opts
}

func (s *Server) DefaultOpts() Opts {
	return Opts{
		ID:            1,
		MaxConnection: 300,
		Tls:           false,
		Status:        "running",
		UpSince:       time.Now(),
	}
}

func (s *Server) NewServer(opts Opts) *Server {
	return &Server{
		Opts: opts,
	}
}
