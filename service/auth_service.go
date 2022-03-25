package service

import (
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
	"fanland/common"
	"fanland/manager"
	"fanland/model"
	"fanland/service/request"
	"fanland/service/response"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type AuthService struct {
	options     *common.ServerOptions
	userManager *manager.UserManager
}

const (
	userId = "uid"
)

func (s *AuthService) InitService(options *common.ServerOptions) {
	s.options = options
	s.userManager = &manager.UserManager{}
	s.userManager.InitManager(options)
	gob.Register(model.User{})
}

// AuthRequired is a simple middleware to check the session
func (s *AuthService) AuthRequired(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(userId)
	if user == nil {
		// Abort the request with the appropriate error code
		//c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		res := response.GenericResponse{Success: true, Message: "unauthorized"}
		c.AbortWithStatusJSON(http.StatusOK, res)
		return
	}
	// Continue down the chain to handler etc
	c.Next()
}

func (s *AuthService) CORSMiddleware(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:8000")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(204)
		return
	}
	c.Next()
}

// login is a handler that parses a form and checks for specific data
func (s *AuthService) Login(c *gin.Context) {
	var req request.LoginRequest

	if err := c.BindJSON(&req); err != nil {
		res := response.GenericResponse{Success: false, Message: err.Error()}
		c.JSON(http.StatusOK, res)
	}

	session := sessions.Default(c)
	// Validate form input
	if strings.Trim(req.UserName, " ") == "" || strings.Trim(req.UserName, " ") == "" || strings.Trim(req.Password, " ") == "" {
		res := response.GenericResponse{Success: false, Message: "Parameters can't be empty"}
		c.JSON(http.StatusOK, res)
		return
	}

	// Check for username and password match, usually from a database
	userInDB, err := s.userManager.GetUserByName(req.UserName)
	if err != nil {
		res := response.GenericResponse{Success: false, Message: err.Error()}
		c.JSON(http.StatusOK, res)
		return
	}

	userHash := sha256.Sum256([]byte(req.Password))
	userHashStr := hex.EncodeToString(userHash[:])

	if userInDB == nil || userInDB.UserName != req.UserName || userInDB.UserHash != userHashStr {
		res := response.GenericResponse{Success: false, Message: "Authentication failed"}
		c.JSON(http.StatusOK, res)
		return
	}

	// Save the username in the session
	session.Set(userId, userInDB) // In real world usage you'd set this to the users ID
	if err := session.Save(); err != nil {
		res := response.GenericResponse{Success: false, Message: "Failed to save session"}
		c.JSON(http.StatusOK, res)
		return
	}
	res := response.GenericResponse{Success: true}
	c.JSON(http.StatusOK, res)
}

func (s *AuthService) Logout(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(userId)
	if user == nil {
		res := response.GenericResponse{Success: false, Message: "Invalid session token"}
		c.JSON(http.StatusOK, res)
		return
	}
	session.Delete(userId)
	if err := session.Save(); err != nil {
		res := response.GenericResponse{Success: false, Message: "Failed to save session"}
		c.JSON(http.StatusOK, res)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
}

func (s *AuthService) Me(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(userId)
	res := response.GenericResponse{Success: true, Result: user}
	c.JSON(http.StatusOK, res)
}

func (s *AuthService) Status(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(userId)
	if user == nil {
		res := response.GenericResponse{Success: true, Message: "Not logged in"}
		c.JSON(http.StatusOK, res)
	}
	res := response.GenericResponse{Success: true, Message: "Already logged in", Result: user}
	c.JSON(http.StatusOK, res)
}
