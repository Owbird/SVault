package server

import (
	"fmt"
	"log"

	"github.com/Owbird/SVault-Engine/pkg/models"
	"github.com/Owbird/SVault-Engine/pkg/server"
)

type ServerFunctions struct {
	server *server.Server
	LogCh  chan models.ServerLog
}

// ShareCallBacks defines a set of callback functions for handling file sharing events.
type ShareCallBacks struct {
	// OnFileSent is called when a file has been successfully sent.
	OnFileSent func()

	// OnSendErr is called when an error occurs during the file sending process.
	OnSendErr func(err error)

	// OnProgressChange is called to provide updates on the progress of the file sharing operation.
	OnProgressChange func(progress models.FileShareProgress)

	// OnCodeReceive is called when the code to initiate the file sharing process has been received.
	OnCodeReceive func(code string)
}

func NewServerFunctions() *ServerFunctions {
	logCh := make(chan models.ServerLog)

	return &ServerFunctions{
		LogCh: logCh,
	}
}

func (sf *ServerFunctions) Host(dir string) {
	sf.server = server.NewServer(dir, sf.LogCh)
	sf.server.Start()
}

func (sf *ServerFunctions) Receive(code string) error {
	return sf.server.Receive(code)
}

func (sf *ServerFunctions) Share(file string, callbacks ShareCallBacks) {
	progressCh := make(chan models.FileShareProgress, 1)

	code, st, err := sf.server.Share(file, progressCh)

	if err != nil && callbacks.OnSendErr != nil {
		callbacks.OnSendErr(err)

		return
	}

	if callbacks.OnCodeReceive != nil {
		callbacks.OnCodeReceive(code)
	}

	go func() {
		for {
			select {
			case status := <-st:
				log.Println(status)
				if status.Error != nil && callbacks.OnSendErr != nil {
					callbacks.OnSendErr(status.Error)

					return
				}

				if !status.OK && status.Error != nil && callbacks.OnSendErr != nil {
					callbacks.OnSendErr(fmt.Errorf("unknown error occurred"))
					return

				} else {
					if callbacks.OnFileSent != nil {
						callbacks.OnFileSent()
					}

					return
				}

			case progress := <-progressCh:
				if callbacks.OnProgressChange != nil {
					callbacks.OnProgressChange(progress)
				}
			}
		}
	}()
}
