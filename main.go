package main

import (
	"fanland/server"
	log "github.com/sirupsen/logrus"
	"os/signal"
	"syscall"
)

const (
	userkey = "user"
)

func main() {
	srv := server.Server{}
	srvCh, err := srv.Start()
	if err != nil {
		return
	}

	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be caught, so don't need to add it
	signal.Notify(srvCh, syscall.SIGINT, syscall.SIGTERM)
	<-srvCh
	log.Println("Shutting down server...")
	srv.Stop()
}
