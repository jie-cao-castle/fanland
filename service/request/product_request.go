package request

import "time"

type ProductByIdRequest struct {
	Id uint64 `json:"id"`
}

type ProductsByCategoryIdRequest struct {
	CategoryId uint64 `json:"categoryId"`
}

type ListRequest struct {
	Offset uint64 `json:"offset"`
	Limit  uint64 `json:"limit"`
}

type LoginRequest struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}

type AddProductRequest struct {
	Name        string `json:"name"`
	ProductDesc string `json:"productDesc"`
	Id          uint64 `json:"id"`
	ImgUrl      string `json:"imgUrl"`
	CreatorId   uint64 `json:"creatorId"`
	CategoryId  uint64 `json:"categoryId"`
	ExternalUrl string `json:"externalUrl"`
}

type UpdateProductRequest struct {
	Name        string `json:"name"`
	ProductDesc string `json:"productDesc"`
	Id          uint64 `json:"id"`
	ImgUrl      string `json:"imgUrl"`
	CategoryId  uint64 `json:"categoryId"`
}

type AddProductTagRequest struct {
	TagName   string `json:"tagName"`
	ProductId uint64 `json:"productId"`
}

type ProductTagRequest struct {
	TagName string `json:"tagName"`
	TagId   uint64 `json:"tagId"`
}

type AddProductSaleRequest struct {
	ProductId       uint64    `json:"productId"`
	ChainId         uint64    `json:"chainId"`
	ChainCode       string    `json:"chainCode"`
	ContractId      uint64    `json:"contractId"`
	Price           uint64    `json:"price"`
	PriceUnit       uint64    `json:"priceUnit"`
	StartTime       time.Time `json:"startTime"`
	EndTime         time.Time `json:"endTime"`
	EffectiveTime   time.Time `json:"effectiveTime"`
	Status          int16     `json:"status"`
	FromUserId      uint64    `json:"fromUserId"`
	TokenId         string    `json:"tokenId"`
	TransactionHash string    `json:"transactionHash"`
}
type UpdateProductSaleRequest struct {
	Id uint64 `json:"id"`
	AddProductSaleRequest
}
