package model

import "time"

type User struct {
	name       string
	avatar     string
	status     string
	desc       string
	id         uint64
	createTime time.Time
	updateTime time.Time
}
