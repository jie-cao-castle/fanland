package service

import (
	"fanland/manager"
	"fanland/model"
	"fanland/server"
	"fanland/service/request"
	"fanland/service/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CategoryService struct {
	categoryManager *manager.CategoryManager
	options         *server.ServerOptions
}

func (s *CategoryService) GetProductCategories(c *gin.Context) {
	var req request.ListRequest

	if err := c.BindJSON(&req); err != nil {
		res := response.GenericResponse{Success: false, Message: err.Error()}
		c.JSON(http.StatusOK, res)
	}

	var categories []*model.ProductCategory
	var err error
	if categories, err = s.categoryManager.GetCategories(req.Offset, req.Limit); err != nil {
		res := response.GenericResponse{Success: false, Message: err.Error()}
		c.JSON(http.StatusOK, res)
	}

	res := response.GenericResponse{Success: true, Result: categories}
	c.JSON(http.StatusOK, res)
}
