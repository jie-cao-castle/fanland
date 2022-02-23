package dao

import "time"

type ProductDO struct {
	name       string
	desc       string
	id         uint64
	imgUrl     string
	nft_id     uint64
	tags       string
	createTime time.Time
	updateTime time.Time
}
