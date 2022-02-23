package fanland
type NftDAO struct {
	id int64
	productId int64
	productName string
	chainId int64
	chainCode string
	chainName string
	tokenSymbol string
	tokenName string
	price int64
	priceUnit int64
	createTime time.Time
	updateTime time.Time
}