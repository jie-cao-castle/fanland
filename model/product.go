package model

import "time"

type Product struct {
	name       string
	desc       string
	id         uint64
	imgUrl     string
	authorId   uint64
	nft        *NftSku
	tags       []*ProductTag
	createTime time.Time
	updateTime time.Time
}
