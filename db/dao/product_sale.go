package dao

import "time"

type ProductSaleDO struct {
	Id            uint64
	ProductId     uint64
	ProductName   string
	ChainId       uint64
	ChainCode     string
	ChainName     string
	ContractId    uint64
	Price         uint64
	PriceUnit     uint64
	StartTime     time.Time
	EndTime       time.Time
	EffectiveTime time.Time
	Status        int16
	CreateTime    time.Time
	UpdateTime    time.Time
	FromUserId    uint64
	FromUserName  string
	TokenId       string
}
