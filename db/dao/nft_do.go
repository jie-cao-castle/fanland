package dao

import "time"

type NftDO struct {
	Id            uint64
	ProductId     uint64
	ProductName   string
	ChainId       uint64
	ContractId    uint64
	ChainCode     string
	ChainName     string
	TokenSymbol   string
	TokenName     string
	LastSoldPrice uint64
	PriceUnit     uint64
	CreateTime    time.Time
	UpdateTime    time.Time
	ownerId       uint64
}
