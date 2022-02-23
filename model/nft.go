package model

type NFT struct {
	id          uint64
	productId   uint64
	productName string
	chainId     uint64
	chainCode   string
	chainName   string
	tokenSymbol string
	tokenName   string
	price       float64
}
