package handler

import (
	"blog/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IAuthService interface {
	Register(req *model.RegisterRequest) (*model.User, error)
	Login(req *model.LoginRequest) (*model.LoginResponse, error)
}

type AuthHandler struct {
	authService IAuthService
}

func NewAuthHandler(authService IAuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req model.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Code:    http.StatusBadRequest,
			Message: "参数错误: " + err.Error(),
			Data:    nil,
		})
		return
	}

	user, err := h.authService.Register(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Code:    http.StatusBadRequest,
			Message: "注册失败: " + err.Error(),
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusCreated, Response{
		Code:    http.StatusCreated,
		Message: "注册成功",
		Data:    user,
	})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req model.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Code:    http.StatusBadRequest,
			Message: "参数错误: " + err.Error(),
			Data:    nil,
		})
		return
	}

	response, err := h.authService.Login(&req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, Response{
			Code:    http.StatusUnauthorized,
			Message: "登录失败: " + err.Error(),
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Code:    http.StatusOK,
		Message: "登录成功",
		Data:    response,
	})
}

// RegisterRoutes 注册路由
func (h *AuthHandler) RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api/v1/auth")
	{
		api.POST("/register", h.Register)
		api.POST("/login", h.Login)
	}
}
