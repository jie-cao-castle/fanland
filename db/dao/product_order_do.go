package dao

import "time"

type ProductOrderDO struct {
	id         uint64
	productId  uint64
	offerId    uint64
	nftId      uint64
	price      uint64
	nftUnit    uint64
	createTime time.Time
	updateTime time.Time
}
