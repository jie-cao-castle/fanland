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

func (s *ProductService) GetProduct(productId uint64) (*model.Product, error) {
	return s.productManager.GetProduct(productId)
}
