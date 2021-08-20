package server

import (
	"github.com/takama/daemon"
	config "github.com/xiusin/pinecms/src/server"
	"os"
	"os/signal"
	"syscall"
)

type Service struct{ daemon.Daemon }

var serv *Service

func (service *Service) Manage(args []string, usage string) (string, error) {
	if len(args) >= 1 && args[0] != "run" {
		switch args[0] {
		case "install":
			return service.Install()
		case "remove":
			return service.Remove()
		case "start":
			return service.Start()
		case "stop":
			return service.Stop()
		case "status":
			return service.Status()
		default:
			return usage, nil
		}
	}

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, os.Kill, syscall.SIGTERM)

	go func() {
		config.InitDB()
		config.Server()
	}()

	for {
		select {
		case killSignal := <-interrupt:
			if killSignal == os.Interrupt {
				return "Daemon was interrupted by system signal", nil
			}
			return "Daemon was killed", nil
		}
	}
}