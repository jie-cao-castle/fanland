package fanland

import (
	"fanland/db"
	"fanland/db/converter"
	"fanland/model"
	"strconv"
	"strings"
)

type ProductManager struct {
	productDB         *dao.ProductDB
	chainNetDB        *dao.ChainNetDB
	chainTokenDB      *dao.ChainTokenDB
	nftDB             *dao.NftDB
	productCategoryDB *dao.ProductCategoryDB
	productOrderDB    *dao.ProductOrderDB
	productTagDB      *dao.ProductTagDB
}

func (manager *ProductManager) GetProduct(productId uint64) (*model.Product, error) {
	manager.productDB.Init()
	defer manager.productDB.Close()
	product, err := manager.productDB.GetById(productId)
	if err != nil {
		return nil, err
	}

	manager.nftDB.Init()
	defer manager.nftDB.Close()
	manager.nftDB.Close()
	nft, err := manager.nftDB.GetById(product.NftId)
	tagIdStrs := strings.Split(product.Tags, ",")

	var tagIds []uint64
	for _, tagId := range tagIdStrs {
		intVar, err := strconv.ParseUint(tagId, 10, 64)
		if err != nil {
			return nil, err
		}

		tagIds = append(tagIds, intVar)
	}

	manager.productTagDB.Init()
	defer manager.nftDB.Close()
	tags, err := manager.productTagDB.GetListByIds(tagIds)
	return converter.ConvertToProduct(product, nft, tags), nil
}
