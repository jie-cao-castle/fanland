package dao

import "time"

//Token on the chain net. e.g. USDT
type ChainTokenDO struct {
	Id          uint64
	TokenSymbol string
	TokenName   string
	TokenDesc   string
	CreateTime  time.Time
	UpdateTime  time.Time
}
