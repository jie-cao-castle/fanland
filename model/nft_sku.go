package model

type NftSku struct {
	id            uint64
	name          uint64
	chainId       uint64
	chainCode     string
	chainName     string
	tokenSymbol   string
	tokenName     string
	unit          uint64
	price         float64
	startTime     uint64
	endTime       uint64
	effectiveTime uint64
	status        int16
}
