package manager

import (
	"fanland/common"
	dao "fanland/db"
)

type NftManager struct {
	nftDB   *dao.NftDB
	options *common.ServerOptions
}
