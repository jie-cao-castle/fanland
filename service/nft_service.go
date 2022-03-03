package service

import (
	"fanland/common"
	"fanland/manager"
)

type NftService struct {
	categoryManager *manager.NftManager
	options         *common.ServerOptions
}
