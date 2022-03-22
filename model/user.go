package model

import "time"

type User struct {
	Id         uint64
	AvatarUrl  string
	UserName   string
	UserDesc   string
	CreateTime time.Time
	UpdateTime time.Time
}
