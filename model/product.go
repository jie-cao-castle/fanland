package model

import "time"

type Product struct {
	Name         string
	Desc         string
	Id           uint64
	ImgUrl       string
	ExternalUrl  string
	Creator      *User
	ProductSales []*ProductSale
	Tags         []*ProductTag
	CreateTime   time.Time
	UpdateTime   time.Time
	Popularity   uint64
}
