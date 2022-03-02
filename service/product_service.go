package service

import (
	"fanland/db/converter"
	"fanland/manager"
	"fanland/model"
	"fanland/server"
	"fanland/service/request"
	"fanland/service/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ProductService struct {
	productManager *manager.ProductManager
	options        *server.ServerOptions
}

func (s *ProductService) InitService(options *server.ServerOptions) {
	s.options = options
	s.productManager = &manager.ProductManager{}
}

func (s *ProductService) UpdateProduct(product *model.Product) error {
	return s.productManager.UpdateProduct(product)
}

func (s *ProductService) GetProductList(categoryId uint64) ([]*model.Product, error) {
	return s.productManager.GetProductsByCategory(categoryId)
}

func (s *ProductService) GetProductById(c *gin.Context) {
	var req request.ProductByIdRequest

	if err := c.BindJSON(&req); err != nil {
		res := response.GenericResponse{Success: false, Message: err.Error()}
		c.JSON(http.StatusOK, res)
	}

	var product *model.Product
	var err error
	if product, err = s.productManager.GetProductDetails(req.ProductId); err != nil {
		res := response.GenericResponse{Success: false, Message: err.Error()}
		c.JSON(http.StatusOK, res)
	}

	res := response.GenericResponse{Success: true, Result: product}
	c.JSON(http.StatusOK, res)
}

func (s *ProductService) AddProduct(c *gin.Context) {
	var req request.AddProductRequest

	if err := c.BindJSON(&req); err != nil {
		res := response.GenericResponse{Success: false, Message: err.Error()}
		c.JSON(http.StatusOK, res)
	}

	var product *model.Product
	product = converter.ConvertReqToProduct(&req)
	if err := s.productManager.AddProduct(product); err != nil {
		res := response.GenericResponse{Success: false, Message: err.Error()}
		c.JSON(http.StatusOK, res)
	}

	res := response.GenericResponse{Success: true, Result: product}
	c.JSON(http.StatusOK, res)
}

func (s *ProductService) GetProductsByCategoryId(c *gin.Context) {
	var req request.ProductsByCategoryIdRequest

	if err := c.BindJSON(&req); err != nil {
		res := response.GenericResponse{Success: false, Message: err.Error()}
		c.JSON(http.StatusOK, res)
	}

	var products []*model.Product
	var err error
	if products, err = s.productManager.GetProductsByCategory(req.CategoryId); err != nil {
		res := response.GenericResponse{Success: false, Message: err.Error()}
		c.JSON(http.StatusOK, res)
	}

	res := response.GenericResponse{Success: true, Result: products}
	c.JSON(http.StatusOK, res)
}
