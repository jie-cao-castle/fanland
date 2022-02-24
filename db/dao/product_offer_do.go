package dao

import "time"

type ProductOfferDO struct {
	Id         int64
	ProductId  int64
	OfferId    int64
	NftId      int64
	Price      int64
	NftUnit    int64
	CreateTime time.Time
	UpdateTime time.Time
}
