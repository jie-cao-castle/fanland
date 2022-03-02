package server

import (
	"context"
	"errors"
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
	options *ServerOptions
	engine  *gin.Engine
	srv     *http.Server
	DoneCh  chan bool

	authService     *service.AuthService
	productService  *service.ProductService
	categoryService *service.CategoryService
}

func (s *Server) Init(options *ServerOptions) *gin.Engine {
	s.options = options
	r := gin.New()
	r.Use(sessions.Sessions("mysession", sessions.NewCookieStore([]byte("secret"))))

	// Initialize all services
	s.initService()

	// router
	r.POST("/products/details", s.productService.GetProductById)
	r.POST("/products/category", s.productService.GetProductsByCategoryId)
	r.POST("/category/list", s.categoryService.GetProductCategories)
	
	r.POST("/login", s.authService.Login)
	r.GET("/logout", s.authService.Logout)

	private := r.Group("/private")
	private.Use(s.authService.AuthRequired)
	{
		private.GET("/me", s.authService.Me)
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
}
