package dao

import "time"

type NftOrderDO struct {
	Id              uint64
	ProductId       uint64
	NftKey          string
	Price           uint64
	PriceUnit       uint64
	Amount          uint64
	Status          int8
	ChainId         uint64
	ChainCode       string
	TransactionHash string
	ToUserId        uint64
	ToUserName      string
	SaleId          uint64
	CreateTime      time.Time
	UpdateTime      time.Time
}
