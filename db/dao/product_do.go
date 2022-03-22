package dao

import "time"

type ProductDO struct {
	Name        string
	Desc        string
	Id          uint64
	ImgUrl      string
	Tags        string
	ExternalUrl string
	CreatorId   uint64
	CreateTime  time.Time
	UpdateTime  time.Time
}
