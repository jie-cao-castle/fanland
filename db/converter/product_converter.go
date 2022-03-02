package converter

import (
	"fanland/db/dao"
	"fanland/model"
	"fanland/service/request"
)

func ConvertToProduct(productDO *dao.ProductDO, nftDO *dao.NftDO, tags []*dao.ProductTagDO) *model.Product {
	return nil
}

func ConvertToProductDO(product *model.Product) *dao.ProductDO {
	return nil
}

func ConvertReqToProduct(req *request.AddProductRequest) *model.Product {
	return nil
}

func ConvertReqToProductTag(req *request.AddProductTagRequest) *model.ProductTag {
	return nil
}

func ConvertToProductTagDO(tag *model.ProductTag) *dao.ProductTagDO {
	return nil
}
