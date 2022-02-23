package dao

import "time"

type ProductCategoryDO struct {
	id         uint64
	name       string
	desc       string
	createTime time.Time
	updateTime time.Time
}
