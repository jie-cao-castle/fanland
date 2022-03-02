package request

type ProductByIdRequest struct {
	ProductId uint64 `json:"productId"`
}

type ProductsByCategoryIdRequest struct {
	CategoryId uint64 `json:"categoryId"`
}

type ListRequest struct {
	Offset uint64 `json:"offset"`
	Limit  uint64 `json:"limit"`
}
