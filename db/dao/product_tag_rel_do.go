package dao

import "time"

type ProductTagRelDO struct {
	Id         uint64
	TagId      uint64
	ProductId  uint64
	CreateTime time.Time
	UpdateTime time.Time
}
