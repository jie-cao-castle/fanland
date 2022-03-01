package dao

import "time"

type NftSkuDO struct {
	Id            uint64
	ProductId     uint64
	ProductName   string
	ChainId       uint64
	ChainCode     string
	ChainName     string
	TokenSymbol   string
	TokenName     string
	Price         uint64
	PriceUnit     uint64
	StartTime     uint64
	EndTime       uint64
	EffectiveTime uint64
	Status        int16
	CreateTime    time.Time
	UpdateTime    time.Time
}
