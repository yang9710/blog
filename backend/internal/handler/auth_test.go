package handler

import (
	"blog/internal/model"
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockAuthService 模拟认证服务
type MockAuthService struct {
	mock.Mock
}

func (m *MockAuthService) Register(req *model.RegisterRequest) (*model.User, error) {
	args := m.Called(req)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.User), args.Error(1)
}

func (m *MockAuthService) Login(req *model.LoginRequest) (*model.LoginResponse, error) {
	args := m.Called(req)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.LoginResponse), args.Error(1)
}

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	return r
}

func TestAuthHandler_Register(t *testing.T) {
	mockService := new(MockAuthService)
	handler := NewAuthHandler(mockService)

	// 测试用例1：成功注册
	t.Run("成功注册", func(t *testing.T) {
		req := model.RegisterRequest{
			Username: "testuser",
			Email:    "test@example.com",
			Password: "password123",
		}

		user := &model.User{
			Username: req.Username,
			Email:    req.Email,
		}

		mockService.On("Register", &req).Return(user, nil)

		router := setupRouter()
		router.POST("/register", handler.Register)

		jsonData, _ := json.Marshal(req)
		reqBody := bytes.NewBuffer(jsonData)
		w := httptest.NewRecorder()
		request, _ := http.NewRequest("POST", "/register", reqBody)
		request.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(w, request)

		assert.Equal(t, http.StatusCreated, w.Code)
		var response model.User
		json.Unmarshal(w.Body.Bytes(), &response)
		assert.Equal(t, req.Username, response.Username)
		assert.Equal(t, req.Email, response.Email)
	})

	// 测试用例2：无效的请求数据
	t.Run("无效的请求数据", func(t *testing.T) {
		invalidReq := map[string]string{
			"username": "te", // 用户名太短
			"email":    "invalid-email",
			"password": "123", // 密码太短
		}

		router := setupRouter()
		router.POST("/register", handler.Register)

		jsonData, _ := json.Marshal(invalidReq)
		reqBody := bytes.NewBuffer(jsonData)
		w := httptest.NewRecorder()
		request, _ := http.NewRequest("POST", "/register", reqBody)
		request.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(w, request)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	// 测试用例3：服务错误
	// t.Run("服务错误", func(t *testing.T) {
	// 	req := model.RegisterRequest{
	// 		Username: "testuser",
	// 		Email:    "test@example.com",
	// 		Password: "password123",
	// 	}

	// 	mockService.On("Register", &req).Return(nil, errors.New("service error"))

	// 	router := setupRouter()
	// 	router.POST("/register", handler.Register)

	// 	jsonData, _ := json.Marshal(req)
	// 	reqBody := bytes.NewBuffer(jsonData)
	// 	w := httptest.NewRecorder()
	// 	request, _ := http.NewRequest("POST", "/register", reqBody)
	// 	request.Header.Set("Content-Type", "application/json")

	// 	router.ServeHTTP(w, request)

	// 	assert.Equal(t, http.StatusBadRequest, w.Code)
	// })
}

func TestAuthHandler_Login(t *testing.T) {
	mockService := new(MockAuthService)
	handler := NewAuthHandler(mockService)

	// 测试用例1：成功登录
	t.Run("成功登录", func(t *testing.T) {
		req := model.LoginRequest{
			Email:    "test@example.com",
			Password: "password123",
		}

		response := &model.LoginResponse{
			Token: "test-token",
			User: model.User{
				Username: "testuser",
				Email:    req.Email,
			},
		}

		mockService.On("Login", &req).Return(response, nil)

		router := setupRouter()
		router.POST("/login", handler.Login)

		jsonData, _ := json.Marshal(req)
		reqBody := bytes.NewBuffer(jsonData)
		w := httptest.NewRecorder()
		request, _ := http.NewRequest("POST", "/login", reqBody)
		request.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(w, request)

		assert.Equal(t, http.StatusOK, w.Code)
		var loginResponse model.LoginResponse
		json.Unmarshal(w.Body.Bytes(), &loginResponse)
		assert.Equal(t, response.Token, loginResponse.Token)
		assert.Equal(t, response.User.Username, loginResponse.User.Username)
		assert.Equal(t, response.User.Email, loginResponse.User.Email)
	})

	// 测试用例2：无效的请求数据
	t.Run("无效的请求数据", func(t *testing.T) {
		invalidReq := map[string]string{
			"email":    "invalid-email",
			"password": "",
		}

		router := setupRouter()
		router.POST("/login", handler.Login)

		jsonData, _ := json.Marshal(invalidReq)
		reqBody := bytes.NewBuffer(jsonData)
		w := httptest.NewRecorder()
		request, _ := http.NewRequest("POST", "/login", reqBody)
		request.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(w, request)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	// 测试用例3：认证失败
	t.Run("认证失败", func(t *testing.T) {
		req := model.LoginRequest{
			Email:    "test@example.com",
			Password: "wrongpassword",
		}

		mockService.On("Login", &req).Return(nil, errors.New("invalid credentials"))

		router := setupRouter()
		router.POST("/login", handler.Login)

		jsonData, _ := json.Marshal(req)
		reqBody := bytes.NewBuffer(jsonData)
		w := httptest.NewRecorder()
		request, _ := http.NewRequest("POST", "/login", reqBody)
		request.Header.Set("Content-Type", "application/json")

		router.ServeHTTP(w, request)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})
}
