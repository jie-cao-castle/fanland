package dao

import "time"

type ProductCategoryRelDO struct {
	Id         uint64
	ProductId  uint64
	CategoryId uint64
	CreateTime time.Time
	UpdateTime time.Time
}
