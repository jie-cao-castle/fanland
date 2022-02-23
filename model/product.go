package model

import "time"

type Product struct {
	name       string
	desc       string
	id         uint64
	imgUrl     string
	nft        *NFT
	tags       []*ProductTag
	createTime time.Time
	updateTime time.Time
}
