package service

import (
	"fanland/common"
)

type Service interface {
	InitService(options *common.ServerOptions)
}
