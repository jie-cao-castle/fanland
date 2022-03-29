package service

import (
	"fanland/common"
	"fanland/db/converter"
	"fanland/manager"
	"fanland/model"
	"fanland/service/request"
	"fanland/service/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ProductService struct {
	productManager *manager.ProductManager
	nftManager     *manager.NftManager
	options        *common.ServerOptions
}

func (s *ProductService) InitService(options *common.ServerOptions) {
	s.options = options
	s.productManager = &manager.ProductManager{}
	s.productManager.InitManager(options)
	s.nftManager = &manager.NftManager{}
	s.nftManager.InitManager(options)
}

func (s *ProductService) GetTitleProduct(c *gin.Context) {
	var product *model.Product
	var err error
	if product, err = s.productManager.GetTitleProduct(); err != nil {
		res := response.GenericResponse{Success: false, Message: err.Error()}
		c.JSON(http.StatusOK, res)
		return
	}

	res := response.GenericResponse{Success: true, Result: product}
	c.JSON(http.StatusOK, res)
}

func (s *ProductService) GetUserProducts(c *gin.Context) {
	uidStr := c.Param("uid")
	uid, err := strconv.ParseUint(uidStr, 10, 64)
	var products []*model.Product
	if products, err = s.productManager.GetProductsByUserId(uid); err != nil {
		res := response.GenericResponse{Success: false, Message: err.Error()}
		c.JSON(http.StatusOK, res)
	}

	res := response.GenericResponse{Success: true, Result: products}
	c.JSON(http.StatusOK, res)
}

func (s *ProductService) GetProductById(c *gin.Context) {
	var req request.AddProductRequest

	if err := c.BindJSON(&req); err != nil {
		res := response.GenericResponse{Success: false, Message: err.Error()}
		c.JSON(http.StatusOK, res)
	}

	var product *model.Product
	var sales []*model.ProductSale
	var err error
	if product, sales, err = s.productManager.GetProductDetails(req.Id); err != nil {
		res := response.GenericResponse{Success: false, Message: err.Error()}
		c.JSON(http.StatusOK, res)
	}

	res := response.GenericResponse{Success: true, Result: &response.ProductDetails{Product: product, Sales: sales}}
	c.JSON(http.StatusOK, res)
}

func (s *ProductService) AddProduct(c *gin.Context) {
	var req request.AddProductRequest

	if err := c.BindJSON(&req); err != nil {
		res := response.GenericResponse{Success: false, Message: err.Error()}
		c.JSON(http.StatusOK, res)
	}

	var product *model.Product
	product = converter.ConvertAddReqToProduct(&req)
	if err := s.productManager.AddProduct(product); err != nil {
		res := response.GenericResponse{Success: false, Message: err.Error()}
		c.JSON(http.StatusOK, res)
	}

	res := response.GenericResponse{Success: true, Result: product}
	c.JSON(http.StatusOK, res)
}

func (s *ProductService) AddProductSale(c *gin.Context) {
	var req request.AddProductSaleRequest

	if err := c.BindJSON(&req); err != nil {
		res := response.GenericResponse{Success: false, Message: err.Error()}
		c.JSON(http.StatusOK, res)
		return
	}

	contract, err := s.nftManager.GetProductContractInChainId(req.ProductId, req.ChainId)
	if err != nil {
		res := response.GenericResponse{Success: false, Message: err.Error()}
		c.JSON(http.StatusOK, res)
		return
	}

	var productSale *model.ProductSale
	productSale = converter.ConvertReqToProductSale(&req)
	productSale.TokenId = strconv.FormatUint(contract.NextTokenId, 10)
	if err := s.productManager.AddProductSale(productSale); err != nil {
		res := response.GenericResponse{Success: false, Message: err.Error()}
		c.JSON(http.StatusOK, res)
		return
	}
	contract.NextTokenId++
	err = s.nftManager.UpdateNFTContractTokenId(contract)
	if err != nil {
		res := response.GenericResponse{Success: false, Message: err.Error()}
		c.JSON(http.StatusOK, res)
		return
	}

	res := response.GenericResponse{Success: true, Result: productSale}
	c.JSON(http.StatusOK, res)
}

func (s *ProductService) UpdateProduct(c *gin.Context) {
	var req request.UpdateProductRequest

	if err := c.BindJSON(&req); err != nil {
		res := response.GenericResponse{Success: false, Message: err.Error()}
		c.JSON(http.StatusOK, res)
	}

	var product *model.Product
	product = converter.ConvertUpdateReqToProduct(&req)
	if err := s.productManager.UpdateProduct(product); err != nil {
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

func (s *ProductService) GetTrendingProducts(c *gin.Context) {
	var products []*model.Product
	var err error
	if products, err = s.productManager.GetTrendingProducts(); err != nil {
		res := response.GenericResponse{Success: false, Message: err.Error()}
		c.JSON(http.StatusOK, res)
		return
	}

	res := response.GenericResponse{Success: true, Result: products}
	c.JSON(http.StatusOK, res)
}

func (s *ProductService) GetProductTags(c *gin.Context) {
	var req request.ProductByIdRequest

	if err := c.BindJSON(&req); err != nil {
		res := response.GenericResponse{Success: false, Message: err.Error()}
		c.JSON(http.StatusOK, res)
	}

	var tags []*model.ProductTag
	var err error

	if tags, err = s.productManager.GetProductTagsByProductId(req.Id); err != nil {
		res := response.GenericResponse{Success: false, Message: err.Error()}
		c.JSON(http.StatusOK, res)
	}

	res := response.GenericResponse{Success: true, Result: tags}
	c.JSON(http.StatusOK, res)
}

func (s *ProductService) GetProductsByTag(c *gin.Context) {
	var req request.ProductTagRequest

	if err := c.BindJSON(&req); err != nil {
		res := response.GenericResponse{Success: false, Message: err.Error()}
		c.JSON(http.StatusOK, res)
	}

	var products []*model.Product
	var err error

	if products, err = s.productManager.GetProductsByTagId(req.TagId); err != nil {
		res := response.GenericResponse{Success: false, Message: err.Error()}
		c.JSON(http.StatusOK, res)
	}

	res := response.GenericResponse{Success: true, Result: products}
	c.JSON(http.StatusOK, res)
}

func (s *ProductService) AddProductTags(c *gin.Context) {
	var req request.AddProductTagRequest

	if err := c.BindJSON(&req); err != nil {
		res := response.GenericResponse{Success: false, Message: err.Error()}
		c.JSON(http.StatusOK, res)
	}

	tag := converter.ConvertReqToProductTag(&req)

	if err := s.productManager.AddProductTag(tag); err != nil {
		res := response.GenericResponse{Success: false, Message: err.Error()}
		c.JSON(http.StatusOK, res)
	}

	res := response.GenericResponse{Success: true}
	c.JSON(http.StatusOK, res)
}
