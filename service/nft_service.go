package service

import (
	"fanland/manager"
	"fanland/server"
)

type NftService struct {
	categoryManager *manager.NftManager
	options         *server.ServerOptions
}
