package manager

import (
	dao "fanland/db"
	"fanland/server"
)

type NftManager struct {
	nftDB   *dao.NftDB
	options *server.ServerOptions
}
