package dao

import "time"

type ChainNetDO struct {
	Id         uint64
	ChainId    uint64
	ChainCode  string
	ChainName  string
	Desc       string
	CreateTime time.Time
	UpdateTime time.Time
}
