package service

import (
	"fanland/common"
	"fanland/service/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"path/filepath"
)

type ProductUploadService struct {
	options *common.ServerOptions
}

func (s *ProductUploadService) InitService(options *common.ServerOptions) {
	s.options = options
}

func (s *ProductUploadService) UploadProduct(c *gin.Context) {
	file, err := c.FormFile("file")

	// The file cannot be received.
	if err != nil {
		res := response.GenericResponse{Success: false, Message: err.Error()}
		c.JSON(http.StatusOK, res)
	}

	// Retrieve file information
	extension := filepath.Ext(file.Filename)
	// Generate random file name for the new uploaded file so it doesn't override the old file with same name
	newFileName := uuid.New().String() + extension

	// The file is received, so let's save it
	if err := c.SaveUploadedFile(file, "/some/path/on/server/"+newFileName); err != nil {
		res := response.GenericResponse{Success: false, Message: err.Error()}
		c.JSON(http.StatusOK, res)
	}

	// File saved successfully. Return proper result
	res := response.GenericResponse{Success: true, Result: newFileName}
	c.JSON(http.StatusOK, res)
}
