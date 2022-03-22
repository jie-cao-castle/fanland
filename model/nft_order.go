package model

import "time"

type NftOrder struct {
	Id         uint64
	ProductId  uint64
	OfferId    uint64
	NftId      uint64
	Price      uint64
	PriceUnit  uint64
	Amount     uint64
	Status     int8
	CreateTime time.Time
	UpdateTime time.Time
}
