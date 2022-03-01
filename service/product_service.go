package service

import (
	"fanland/manager"
	"fanland/model"
	"fanland/server"
)

type ProductService struct {
	productManager *manager.ProductManager
	options        *server.ServerOptions
}

func (s *ProductService) InitService(options *server.ServerOptions) {
	s.options = options
	s.productManager = &manager.ProductManager{}
}

func (s *ProductService) AddProduct(product *model.Product) error {
	return s.productManager.AddProduct(product)
}

func (s *ProductService) UpdateProduct(product *model.Product) error {
	return s.productManager.UpdateProduct(product)
}

func (s *ProductService) GetProduct(productId uint64) (*model.Product, error) {
	return s.productManager.GetProductDetails(productId)
}

func (s *ProductService) GetProductList(categoryId uint64) ([]*model.Product, error) {
	return s.productManager.GetProductsByCategory(categoryId)
}
