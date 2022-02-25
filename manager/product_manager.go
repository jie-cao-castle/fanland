package fanland

import (
	"fanland/db"
	"fanland/db/converter"
	"fanland/model"
)

type ProductManager struct {
	productDb         *dao.ProductDB
	chainNetDB        *dao.ChainNetDB
	chainTokenDB      *dao.ChainTokenDB
	nftDB             *dao.NftDB
	productCategoryDB *dao.ProductCategoryDB
	productOrderDB    *dao.ProductOrderDB
}

func (manager *ProductManager) GetProduct(productId uint64) (*model.Product, error) {
	productDB := &dao.ProductDB{}
	productDB.Init()
	product, err := productDB.GetById(productId)
	if err != nil {
		return nil, err
	}

	return converter.ConvertToProduct(product), nil
}
