package model

import "time"

type ProductOffer struct {
	id         uint64
	product    *Product
	nft        *NFT
	nftUnit    uint64
	createTime time.Time
	updateTime time.Time
	status     int32
}
