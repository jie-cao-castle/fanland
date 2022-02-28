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
	auth    *service.Auth
	srv     *http.Server
}

func (s *Server) Init() *gin.Engine {
	r := gin.New()
	r.Use(sessions.Sessions("mysession", sessions.NewCookieStore([]byte("secret"))))

	s.initService()
	r.POST("/login", s.auth.Login)
	r.GET("/logout", s.auth.Logout)

	private := r.Group("/private")
	private.Use(s.auth.AuthRequired)
	{
		private.GET("/me", s.auth.Me)
		private.GET("/status", s.auth.Status)
	}
	s.engine = r
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

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
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

	return nil
}

func (s *Server) initService() {

}
