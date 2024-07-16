package handlers

import "github.com/gin-gonic/gin"

type AuthHandler interface {
	Login(c *gin.Context)
	Register(c *gin.Context)
	GetUserInfo(c *gin.Context)
	UpdateUserInfo(c *gin.Context)
	DeleteUser(c *gin.Context)
	GetUsers(c *gin.Context)
	GetUserByUsernameOrEmail(c *gin.Context)
	ChangeUserType(c *gin.Context)
}

type authHandler struct {
}

func NewAuthHandler() AuthHandler {
	return &authHandler{}
}

func (h *authHandler) Login(c *gin.Context) {
	// Implement login logic
}

func (h *authHandler) Register(c *gin.Context) {
	// Implement register logic
}

func (h *authHandler) GetProfileId(c *gin.Context) {
	// Implement get profile by id logic
}

func (h *authHandler) GetUserInfo(c *gin.Context) {
	// Implement get user info logic
}

func (h *authHandler) UpdateUserInfo(c *gin.Context) {
	// Implement update user info logic
}

func (h *authHandler) DeleteUser(c *gin.Context) {
	// Implement delete user logic
}

func (h *authHandler) GetUsers(c *gin.Context) {
	// Implement get users logic
}

func (h *authHandler) GetUserByUsernameOrEmail(c *gin.Context) {
	// Implement get user by username or email logic
}

func (h *authHandler) ChangeUserType(c *gin.Context) {
	// Implement change user type logic
}
