package response

import "fanland/model"

type ProductDetails struct {
	Product *model.Product       `json:"product"`
	Sales   []*model.ProductSale `json:"sales"`
}
