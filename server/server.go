package server

import "github.com/20zinnm/ecs"

type Server struct {
	ecs.World
}

func New(addr string) {
	s := Server{ecs.World{}}
	s.World.AddSystem()
}

func (s *Server) Start() {

}
