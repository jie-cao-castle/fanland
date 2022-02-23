package db

type ProductDO struct {
	name string
	desc string
	id int64
	imgUrl string
	nft_id int64
	tags string
	createTime time.Time
	updateTime time.Time
}