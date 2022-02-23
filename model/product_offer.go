type ProductOffer struct {
	id int64
	product *model.Product
	nft *model.NFT
	nftUnit int64
	createTime time.Time
	updateTime time.Time
	status int32
}