package model

import "time"

type ProductTag struct {
	name       string
	id         uint64
	createTime time.Time
	updateTime time.Time
}
