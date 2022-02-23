package fanland

type ProductOfferDAO struct {
	id int64
	productId int64
	offerId int64
	nftId int64
	price int64
	nftUnit int64
	createTime time.Time
	updateTime time.Time
}