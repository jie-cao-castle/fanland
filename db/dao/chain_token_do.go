package dao

import "time"

type ChainTokenDO struct {
	Id          uint64
	TokenSymbol string
	TokenName   string
	TokenDesc   string
	CreateTime  time.Time
	UpdateTime  time.Time
}
