package server

import (
	"github.com/Owbird/SVault-Engine/pkg/models"
	"github.com/Owbird/SVault-Engine/pkg/server"
)

type ServerFunctions struct {
	server *server.Server
	LogCh  chan models.ServerLog
}

func NewServerFunctions(dir string) *ServerFunctions {
	logCh := make(chan models.ServerLog)

	return &ServerFunctions{
		server: server.NewServer(dir, logCh),
		LogCh:  logCh,
	}
}

func (sf *ServerFunctions) Start() {
	sf.server.Start()
}
