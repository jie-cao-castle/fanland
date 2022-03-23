package service

import (
	"fanland/common"
	"fanland/manager"
	"fanland/service/response"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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
}

// AuthRequired is a simple middleware to check the session
func (s *AuthService) AuthRequired(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(userId)
	if user == nil {
		// Abort the request with the appropriate error code
		//c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		res := response.GenericResponse{Success: true, Message: "unauthorized"}
		c.AbortWithStatusJSON(http.StatusUnauthorized, res)
		return
	}
	// Continue down the chain to handler etc
	c.Next()
}

// login is a handler that parses a form and checks for specific data
func (s *AuthService) Login(c *gin.Context) {
	session := sessions.Default(c)
	username := c.PostForm("username")
	password := c.PostForm("password")
	uidStr := c.PostForm("userId")
	uid, err := strconv.ParseUint(uidStr, 10, 64)
	if err != nil {
		res := response.GenericResponse{Success: false, Message: "Authentication failed"}
		c.JSON(http.StatusOK, res)
		return
	}
	// Validate form input
	if strings.Trim(username, " ") == "" || strings.Trim(username, " ") == "" || strings.Trim(password, " ") == "" {
		res := response.GenericResponse{Success: false, Message: "Parameters can't be empty"}
		c.JSON(http.StatusOK, res)
		return
	}

	// Check for username and password match, usually from a database
	userInDB, err := s.userManager.GetUser(uid)
	if err != nil {
		res := response.GenericResponse{Success: false, Message: err.Error()}
		c.JSON(http.StatusOK, res)
		return
	}

	if userInDB == nil || userInDB.UserName != username || userInDB.UserHash != password {
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
	c.JSON(http.StatusOK, gin.H{"message": "Successfully authenticated user"})
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
	c.JSON(http.StatusOK, gin.H{"user": user})
}

func (s *AuthService) Status(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(userId)
	if user == nil {
		res := response.GenericResponse{Success: true, Message: "Already logged in"}
		c.JSON(http.StatusOK, res)
	}
	c.JSON(http.StatusOK, gin.H{"status": "You are logged in"})
}
