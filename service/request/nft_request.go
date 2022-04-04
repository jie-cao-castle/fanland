package request

type AddNftRequest struct {
	Name        string `json:"name"`
	ProductDesc string `json:"productDesc"`
	Id          uint64 `json:"id"`
	ImgUrl      string `json:"imgUrl"`
	AuthorId    uint64 `json:"authorId"`
	CategoryId  uint64 `json:"categoryId"`
}

type AddNftContractRequest struct {
	ProductId       uint64 `json:"productId"`
	ChainId         uint64 `json:"chainId"`
	ChainCode       string `json:"chainCode"`
	ContractAddress string `json:"contractAddress"`
	Status          int8   `json:"status"`
	TokenSymbol     string `json:"tokenSymbol"`
	TokenName       string `json:"tokenName"`
	TokenAmount     uint64 `json:"tokenAmount"`
	NextTokenId     uint64 `json:"nextTokenId"`
}
type UpdateNftContractRequest struct {
	Id     uint64 `json:"id"`
	Status int8   `json:"status"`
}
type ProductContractRequest struct {
	ProductId uint64 `json:"productId"`
}

type AddNftOrderRequest struct {
	ProductId       uint64 `json:"productId"`
	NftKey          string `json:"nftKey"`
	Price           uint64 `json:"price"`
	PriceUnit       uint64 `json:"priceUnit"`
	Amount          uint64 `json:"amount"`
	Status          int8   `json:"status"`
	ChainId         uint64 `json:"chainId"`
	ChainCode       string `json:"chainCode"`
	TransactionHash string `json:"transactionHash"`
	ToUserId        uint64 `json:"toUserId"`
	SaleId          uint64 `json:"saleId"`
}
type UpdateNftOrderRequest struct {
	Id uint64 `json:"id"`
	AddNftOrderRequest
}
