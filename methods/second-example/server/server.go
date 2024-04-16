package server

import (
	"time"
)

type Opts struct {
	MaxConnection uint
	Id            uint
	Tls           bool
	Status        string
	UpSince       time.Time
}

type Server struct {
	Opts
}

func (s *Server) DefaultOpts() Opts {
	return Opts{
		MaxConnection: 100,
		Id:            1,
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
