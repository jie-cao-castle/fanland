package converter

import (
	"fanland/db/dao"
	"fanland/model"
	"fanland/service/request"
)

func ConvertToProduct(productDO *dao.ProductDO, userDO *dao.UserDO, tags []*dao.ProductTagDO) *model.Product {
	product := &model.Product{
		Id:          productDO.Id,
		Name:        productDO.Name,
		Desc:        productDO.Desc,
		ExternalUrl: productDO.ExternalUrl,
		ImgUrl:      productDO.ImgUrl,
		CreateTime:  productDO.CreateTime,
		UpdateTime:  productDO.UpdateTime,
	}
	product.Creator = ConvertToUser(userDO)
	return product
}

func ConvertToProductSale(saleDO *dao.ProductSaleDO) *model.ProductSale {
	sale := &model.ProductSale{
		Id:           saleDO.Id,
		ProductId:    saleDO.ProductId,
		ProductName:  saleDO.ProductName,
		ChainId:      saleDO.ChainId,
		ChainCode:    saleDO.ChainCode,
		ChainName:    saleDO.ChainName,
		ContractId:   saleDO.ContractId,
		Price:        saleDO.Price,
		PriceUnit:    saleDO.PriceUnit,
		StartTime:    saleDO.StartTime,
		EndTime:      saleDO.EndTime,
		Status:       saleDO.Status,
		TokenId:      saleDO.TokenId,
		CreateTime:   saleDO.CreateTime,
		UpdateTime:   saleDO.UpdateTime,
		FromUserId:   saleDO.FromUserId,
		FromUserName: saleDO.FromUserName,
	}
	return sale
}

func ConvertAddReqToProduct(req *request.AddProductRequest) *model.Product {
	creater := &model.User{Id: req.CreatorId}

	product := &model.Product{
		Name:        req.Name,
		Desc:        req.ProductDesc,
		Id:          req.Id,
		ImgUrl:      req.ImgUrl,
		ExternalUrl: req.ExternalUrl,
		Creator:     creater,
	}

	return product
}

func ConvertToProductDO(product *model.Product) *dao.ProductDO {
	productDO := &dao.ProductDO{
		Name:        product.Name,
		Desc:        product.Desc,
		Id:          product.Id,
		ImgUrl:      product.ImgUrl,
		ExternalUrl: product.ExternalUrl,
		CreatorId:   product.Creator.Id,
	}

	return productDO
}

func ConvertToProductSaleDO(product *model.ProductSale) *dao.ProductSaleDO {
	productDo := &dao.ProductSaleDO{
		ProductId:     product.ProductId,
		ChainId:       product.ChainId,
		ChainCode:     product.ChainCode,
		ContractId:    product.ContractId,
		Price:         product.Price,
		PriceUnit:     product.PriceUnit,
		StartTime:     product.StartTime,
		EndTime:       product.EndTime,
		EffectiveTime: product.EffectiveTime,
		Status:        product.Status,
		FromUserId:    product.FromUserId,
		FromUserName:  product.FromUserName,
		TokenId:       product.TokenId,
	}
	return productDo
}

func ConvertReqToProductSale(req *request.AddProductSaleRequest) *model.ProductSale {
	productSale := &model.ProductSale{
		ProductId:     req.ProductId,
		ChainId:       req.ChainId,
		ChainCode:     req.ChainCode,
		ContractId:    req.ContractId,
		Price:         req.Price,
		PriceUnit:     req.PriceUnit,
		StartTime:     req.StartTime,
		EndTime:       req.EndTime,
		EffectiveTime: req.EffectiveTime,
		Status:        req.Status,
		FromUserId:    req.FromUserId,
		TokenId:       req.TokenId,
	}
	return productSale
}

func ConvertAddReqToNft(req *request.AddNftRequest) *model.NFT {
	return nil
}

func ConvertReqToProductTag(req *request.AddProductTagRequest) *model.ProductTag {
	return nil
}

func ConvertToProductTagDO(tag *model.ProductTag) *dao.ProductTagDO {
	return nil
}

func ConvertToProductTag(tagDO *dao.ProductTagDO) *model.ProductTag {
	return nil
}

func ConvertToProductCategory(categoryDO *dao.ProductCategoryDO) *model.ProductCategory {
	return nil
}

func ConvertUpdateReqToProduct(updateRequest *request.UpdateProductRequest) *model.Product {
	return nil
}
