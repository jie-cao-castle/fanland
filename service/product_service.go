package service

import (
	"fanland/manager"
	"fanland/model"
)

type ProductService struct {
	productManager *manager.ProductManager
}

func (s *ProductService) GetProduct(productId uint64) (*model.Product, error) {
	return s.productManager.GetProduct(productId)
}
