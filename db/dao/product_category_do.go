package dao

import "time"

type ProductCategoryDO struct {
	Id         uint64
	Name       string
	Desc       string
	CreateTime time.Time
	UpdateTime time.Time
}
