package manager

import (
	"fanland/db"
	"fanland/db/converter"
	"fanland/model"
	"fanland/server"
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
	options           *server.ServerOptions
}

func (manager *ProductManager) InitManager(options *server.ServerOptions) {
	manager.options = options
	manager.productDB.InitDB(options.DbName)
	manager.nftDB.InitDB(options.DbName)
}

func (manager *ProductManager) GetProduct(productId uint64) (*model.Product, error) {
	manager.productDB.Open()
	defer manager.productDB.Close()
	product, err := manager.productDB.GetById(productId)
	if err != nil {
		return nil, err
	}

	manager.nftDB.Open()
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
