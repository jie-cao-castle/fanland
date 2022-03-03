package main

import (
	"fanland/common"
	"fanland/server"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os/signal"
	"syscall"
)

type App struct {
	options *common.ServerOptions
}

func (app *App) Start(cmd *cobra.Command, args []string) {
	app.options = &common.ServerOptions{}
	cmd.Flags().StringVarP(&app.options.DbName, "DB", "d", "", "db name")
	srv := server.Server{}
	srv.Init(app.options)
	srvCh, err := srv.Start()
	if err != nil {
		return
	}

	signal.Notify(srvCh, syscall.SIGINT, syscall.SIGTERM)
	<-srvCh
	log.Println("Shutting down server...")
	srv.Stop()
	<-srv.DoneCh
	log.Println("server shut down ...")
}
