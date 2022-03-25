package manager

import (
	"fanland/common"
	"fanland/db"
	"fanland/db/converter"
	"fanland/model"
	"strconv"
	"strings"
)

type ProductManager struct {
	productDB            *dao.ProductDB
	productSaleDB        *dao.ProductSaleDB
	userDB               *dao.UserDB
	chainNetDB           *dao.ChainNetDB
	chainTokenDB         *dao.ChainTokenDB
	nftDB                *dao.NftDB
	productCategoryDB    *dao.ProductCategoryDB
	productTagDB         *dao.ProductTagDB
	productTagRelDB      *dao.ProductTagRelDB
	productCategoryRelDB *dao.ProductCategoryRelDB
	options              *common.ServerOptions
}

func (manager *ProductManager) InitManager(options *common.ServerOptions) {
	manager.options = options
	manager.productDB = &dao.ProductDB{}
	manager.nftDB = &dao.NftDB{}
	manager.productSaleDB = &dao.ProductSaleDB{}
	manager.userDB = &dao.UserDB{}

	manager.productDB.InitDB(options.DbName)
	manager.nftDB.InitDB(options.DbName)
	manager.productSaleDB.InitDB(options.DbName)
	manager.userDB.InitDB(options.DbName)
}

func (manager *ProductManager) GetTitleProduct() (*model.Product, error) {
	manager.productDB.Open()
	defer manager.productDB.Close()
	productDO, err := manager.productDB.GetTitleProduct()
	if err != nil {
		return nil, err
	}
	if productDO == nil {
		return nil, nil
	}

	manager.userDB.Open()
	defer manager.userDB.Close()
	userDO, err := manager.userDB.GetById(productDO.Id)

	var product = converter.ConvertToProduct(productDO, userDO, nil)
	return product, nil
}

func (manager *ProductManager) GetProductsByUserId(userId uint64) ([]*model.Product, error) {
	manager.productDB.Open()
	defer manager.productDB.Close()
	productDOs, err := manager.productDB.GetListByUserId(userId)
	if err != nil {
		return nil, err
	}

	var products []*model.Product
	for i, productDO := range productDOs {
		userDO, err := manager.userDB.GetById(productDO.Id)
		if err != nil {
			return nil, err
		}
		products[i] = converter.ConvertToProduct(productDO, userDO, nil)
	}

	return products, nil
}

func (manager *ProductManager) GetProductDetails(productId uint64) (*model.Product, []*model.ProductSale, error) {
	manager.productDB.Open()
	defer manager.productDB.Close()
	productDO, err := manager.productDB.GetById(productId)
	if err != nil {
		return nil, nil, err
	}

	manager.nftDB.Open()
	defer manager.nftDB.Close()
	manager.nftDB.Close()

	var product *model.Product
	if len(productDO.Tags) > 0 {
		tagIdStrs := strings.Split(productDO.Tags, ",")

		var tagIds []uint64
		for _, tagId := range tagIdStrs {
			intVar, err := strconv.ParseUint(tagId, 10, 64)
			if err != nil {
				return nil, nil, err
			}

			tagIds = append(tagIds, intVar)
		}

		manager.productTagDB.Init()
		defer manager.nftDB.Close()
		tags, err := manager.productTagDB.GetListByIds(tagIds)
		if err != nil {
			return nil, nil, err
		}
		product = converter.ConvertToProduct(productDO, nil, tags)
	}

	salesDO, err := manager.productSaleDB.GetListByProductId(product.Id)
	var sales []*model.ProductSale
	for _, saleDO := range salesDO {
		sale := converter.ConvertToProductSale(saleDO)
		sales = append(sales, sale)
	}

	return product, sales, nil
}

func (manager *ProductManager) GetProductsByCategory(categoryId uint64) ([]*model.Product, error) {
	manager.productCategoryRelDB.Open()
	defer manager.productCategoryRelDB.Close()
	relationships, err := manager.productCategoryRelDB.GetByRelationships(categoryId)
	if err != nil {
		return nil, err
	}

	var productIds []uint64
	for i, rel := range relationships {
		productIds[i] = rel.ProductId
	}
	productDOs, err := manager.productDB.GetListByIds(productIds)

	if err != nil {
		return nil, err
	}

	var products []*model.Product
	for i, productDO := range productDOs {
		products[i] = converter.ConvertToProduct(productDO, nil, nil)
	}

	return products, nil
}

func (manager *ProductManager) GetProductsByTagId(tagId uint64) ([]*model.Product, error) {
	manager.productTagRelDB.Open()
	defer manager.productTagRelDB.Close()
	relationships, err := manager.productTagRelDB.GetListByTagId(tagId)
	if err != nil {
		return nil, err
	}

	var productIds []uint64
	for i, rel := range relationships {
		productIds[i] = rel.ProductId
	}
	productDOs, err := manager.productDB.GetListByIds(productIds)

	if err != nil {
		return nil, err
	}

	var products []*model.Product
	for i, productDO := range productDOs {
		products[i] = converter.ConvertToProduct(productDO, nil, nil)
	}

	return products, nil
}

func (manager *ProductManager) GetProductTagsByProductId(productId uint64) ([]*model.ProductTag, error) {
	manager.productTagRelDB.Open()
	defer manager.productTagRelDB.Close()
	productTagRels, err := manager.productTagRelDB.GetListByProductId(productId)
	if err != nil {
		return nil, err
	}

	var productIds []uint64
	for _, productTagRel := range productTagRels {
		productIds = append(productIds, productTagRel.ProductId)
	}

	productTagDOs, err := manager.productTagDB.GetListByIds(productIds)
	var productTags []*model.ProductTag
	for _, productTagDO := range productTagDOs {
		productTag := converter.ConvertToProductTag(productTagDO)
		productTags = append(productTags, productTag)
	}

	return productTags, nil
}

func (manager *ProductManager) AddProductTag(productTag *model.ProductTag) error {
	manager.productTagDB.Open()
	defer manager.productTagDB.Close()
	tag := converter.ConvertToProductTagDO(productTag)
	if err := manager.productTagDB.Insert(tag); err != nil {
		return err
	}

	return nil
}

func (manager *ProductManager) AddProduct(product *model.Product) error {
	manager.productDB.Open()
	defer manager.productDB.Close()
	productDO := converter.ConvertToProductDO(product)
	if err := manager.productDB.Insert(productDO); err != nil {
		return err
	}

	return nil
}

func (manager *ProductManager) AddProductSale(product *model.ProductSale) error {
	manager.productDB.Open()
	defer manager.productDB.Close()
	productDO := converter.ConvertToProductSaleDO(product)
	if err := manager.productSaleDB.Insert(productDO); err != nil {
		return err
	}

	return nil
}

func (manager *ProductManager) UpdateProduct(product *model.Product) error {
	manager.productDB.Open()
	defer manager.productDB.Close()
	productDO := converter.ConvertToProductDO(product)
	if err := manager.productDB.Update(productDO); err != nil {
		return err
	}

	return nil
}
