type Product struct {
	name string;
	desc string;
	id int64;
	imgUrl string;
	nft *NFT;
	tags []*ProductTag;
}