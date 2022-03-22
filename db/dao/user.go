package dao

import "time"

type UserDO struct {
	Id         uint64
	AvatarUrl  string
	UserName   string
	UserDesc   string
	CreateTime time.Time
	UpdateTime time.Time
}
