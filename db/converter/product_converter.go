package converter

import (
	"fanland/db/dao"
	"fanland/model"
)

func ConvertToProduct(productDO *dao.ProductDO, nftDO *dao.NftDO, tags []*dao.ProductTagDO) *model.Product {
	return nil
}

func ConvertToProductDO(product *model.Product) *dao.ProductDO {
	return nil
}
