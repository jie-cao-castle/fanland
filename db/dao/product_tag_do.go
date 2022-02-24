package dao

import "time"

type ProductTagDO struct {
	Name       string
	Id         uint64
	CreateTime time.Time
	UpdateTime time.Time
}
