package server

import "github.com/Owbird/SVault-Engine/pkg/server"

type ServerFunctions struct {
	server *server.Server
}

func NewServerFunctions(dir string) *ServerFunctions {
	return &ServerFunctions{
		server: server.NewServer(dir),
	}
}

func (sf *ServerFunctions) Start() {
	sf.server.Start()
}
