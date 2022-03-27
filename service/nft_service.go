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
)

type NftService struct {
	nftManager *manager.NftManager
	options    *common.ServerOptions
}

func (s *NftService) InitService(options *common.ServerOptions) {
	s.options = options
	s.nftManager = &manager.NftManager{}
	s.nftManager.InitManager(options)
}

func (s *NftService) AddNFTContract(c *gin.Context) {
	var req request.AddNftContractRequest

	if err := c.BindJSON(&req); err != nil {
		res := response.GenericResponse{Success: false, Message: err.Error()}
		c.JSON(http.StatusOK, res)
	}

	var contract *model.NftContract
	contract = converter.ConvertReqToNftContract(&req)
	if err := s.nftManager.AddNFTContract(contract); err != nil {
		res := response.GenericResponse{Success: false, Message: err.Error()}
		c.JSON(http.StatusOK, res)
	}

	res := response.GenericResponse{Success: true, Result: contract}
	c.JSON(http.StatusOK, res)
}

func (s *NftService) AddNFTOrder(c *gin.Context) {
	var req request.AddNftOrderRequest

	if err := c.BindJSON(&req); err != nil {
		res := response.GenericResponse{Success: false, Message: err.Error()}
		c.JSON(http.StatusOK, res)
	}

	var order *model.NftOrder
	order = converter.ConvertReqToNftOrder(&req)
	if err := s.nftManager.AddNFTOrder(order); err != nil {
		res := response.GenericResponse{Success: false, Message: err.Error()}
		c.JSON(http.StatusOK, res)
	}

	res := response.GenericResponse{Success: true, Result: order}
	c.JSON(http.StatusOK, res)
}

func (s *NftService) GetNFTContractsByProduct(c *gin.Context) {
	var req request.ProductContractRequest

	if err := c.BindJSON(&req); err != nil {
		res := response.GenericResponse{Success: false, Message: err.Error()}
		c.JSON(http.StatusOK, res)
	}

	var contracts []*model.NftContract
	contracts, err := s.nftManager.GetProductContracts(req.ProductId)
	if err != nil {
		res := response.GenericResponse{Success: false, Message: err.Error()}
		c.JSON(http.StatusOK, res)
	}

	res := response.GenericResponse{Success: true, Result: contracts}
	c.JSON(http.StatusOK, res)
}

func (s *NftService) GetNFTOrdersByProduct(c *gin.Context) {
	var req request.ProductContractRequest

	if err := c.BindJSON(&req); err != nil {
		res := response.GenericResponse{Success: false, Message: err.Error()}
		c.JSON(http.StatusOK, res)
	}

	var orders []*model.NftOrder
	orders, err := s.nftManager.GetProductOrders(req.ProductId)
	if err != nil {
		res := response.GenericResponse{Success: false, Message: err.Error()}
		c.JSON(http.StatusOK, res)
	}

	res := response.GenericResponse{Success: true, Result: orders}
	c.JSON(http.StatusOK, res)
}

func (s *NftService) UpdateNFTOrder(c *gin.Context) {
	var req request.UpdateNftContractRequest

	if err := c.BindJSON(&req); err != nil {
		res := response.GenericResponse{Success: false, Message: err.Error()}
		c.JSON(http.StatusOK, res)
	}

	var contract *model.NftContract
	contract = converter.ConvertReqToUpdateNftContract(&req)
	if err := s.nftManager.UpdateNFTContract(contract); err != nil {
		res := response.GenericResponse{Success: false, Message: err.Error()}
		c.JSON(http.StatusOK, res)
	}

	res := response.GenericResponse{Success: true, Result: contract}
	c.JSON(http.StatusOK, res)
}

func (s *NftService) UpdateNFTContract(c *gin.Context) {
	var req request.UpdateNftContractRequest

	if err := c.BindJSON(&req); err != nil {
		res := response.GenericResponse{Success: false, Message: err.Error()}
		c.JSON(http.StatusOK, res)
	}

	var order *model.NftContract
	order = converter.ConvertReqToUpdateNftContract(&req)
	if err := s.nftManager.UpdateNFTContract(order); err != nil {
		res := response.GenericResponse{Success: false, Message: err.Error()}
		c.JSON(http.StatusOK, res)
	}

	res := response.GenericResponse{Success: true, Result: order}
	c.JSON(http.StatusOK, res)
}
