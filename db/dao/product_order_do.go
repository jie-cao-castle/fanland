package dao

import "time"

type ProductOrderDO struct {
	Id         uint64
	ProductId  uint64
	OfferId    uint64
	NftId      uint64
	Price      uint64
	NftUnit    uint64
	CreateTime time.Time
	UpdateTime time.Time
}
