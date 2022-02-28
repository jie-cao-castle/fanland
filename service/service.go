package service

import "fanland/server"

type Service interface {
	InitService(options *server.ServerOptions)
}
