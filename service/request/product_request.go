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

type AddProductRequest struct {
	Name        string `json:"name"`
	ProductDesc string `json:"productDesc"`
	Id          uint64 `json:"id"`
	ImgUrl      string `json:"imgUrl"`
	AuthorId    uint64 `json:"authorId"`
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