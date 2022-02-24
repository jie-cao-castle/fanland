package dao

import "time"

type ProductDO struct {
	Name       string
	Desc       string
	Id         uint64
	ImgUrl     string
	NftId      uint64
	Tags       string
	CreateTime time.Time
	UpdateTime time.Time
}
