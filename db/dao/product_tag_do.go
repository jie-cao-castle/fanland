package dao

import "time"

type ProductTagDO struct {
	name       string
	id         uint64
	createTime time.Time
	updateTime time.Time
}
