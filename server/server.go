package server

import (
	"context"
	"errors"
	"fanland/common"
	"fanland/service"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"time"
)

const (
	userkey = "user"
)

type Server struct {
	options *common.ServerOptions
	engine  *gin.Engine
	srv     *http.Server
	DoneCh  chan bool

	authService     *service.AuthService
	productService  *service.ProductService
	categoryService *service.CategoryService

	productUploadService *service.ProductUploadService
	nftService           *service.NftService
}

func (s *Server) Init(options *common.ServerOptions) *gin.Engine {
	s.options = options
	//TODO:: replace with config
	s.options.DbName = "fanland"
	r := gin.New()
	r.Use(sessions.Sessions("mysession", sessions.NewCookieStore([]byte("secret"))))

	// Initialize all services
	s.initService()
	r.Static("/upload", "./upload")

	v1 := r.Group("/api/v1")
	//v1.Use(s.authService.CORSMiddleware)

	v1.POST("/products/details", s.productService.GetProductById)
	v1.POST("/products/category", s.productService.GetProductsByCategoryId)
	v1.POST("/products/add", s.productService.AddProduct)
	v1.POST("/products/update", s.productService.UpdateProduct)
	v1.POST("/products/tags", s.productService.GetProductsByTag)

	v1.POST("/products/addSale", s.productService.AddProductSale)
	v1.GET("/products/top", s.productService.GetTitleProduct)
	v1.GET("/products/user/:uid", s.productService.GetUserProducts)

	v1.POST("/asset/addContract", s.nftService.AddNFTContract)
	v1.POST("/asset/addOrder", s.nftService.AddNFTOrder)
	v1.POST("/asset/updateContract", s.nftService.UpdateNFTContract)
	v1.POST("/asset/updateOrder", s.nftService.UpdateNFTOrder)
	v1.POST("/asset/contracts", s.nftService.GetNFTContractsByProduct)
	v1.POST("/asset/orders", s.nftService.GetNFTOrdersByProduct)

	v1.POST("/productsUpload/postContent", s.productUploadService.UploadProductImg)

	v1.POST("/tags/products", s.productService.GetProductTags)

	v1.POST("/category/list", s.categoryService.GetProductCategories)

	v1.GET("/user/logout", s.authService.Logout)
	v1.POST("/user/login", s.authService.Login)
	private := v1.Group("/user")
	private.Use(s.authService.AuthRequired)
	{

		private.GET("/current", s.authService.Me)
		private.GET("/status", s.authService.Status)
	}
	s.engine = r
	s.DoneCh = make(chan bool, 1)
	return r
}

func (s *Server) Start() (chan os.Signal, error) {
	s.srv = &http.Server{
		Addr:    ":8080",
		Handler: s.engine,
	}
	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := s.srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Printf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)

	return quit, nil
}

func (s *Server) Stop() error {
	log.Println("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
		return err
	}

	log.Println("Server exiting")
	s.DoneCh <- true
	return nil
}

func (s *Server) initService() {
	s.authService = &service.AuthService{}
	s.authService.InitService(s.options)

	s.productService = &service.ProductService{}
	s.productService.InitService(s.options)

	s.productUploadService = &service.ProductUploadService{}
	s.productUploadService.InitService(s.options)

	s.nftService = &service.NftService{}
	s.nftService.InitService(s.options)
}
