package model

import "time"

type NftContract struct {
	Id              uint64
	ProductId       uint64
	ChainId         uint64
	ChainCode       string
	ContractAddress string
	Status          int8

	TokenSymbol string
	TokenName   string

	CreateTime time.Time
	UpdateTime time.Time
}
