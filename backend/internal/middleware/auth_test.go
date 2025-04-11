package middleware

import (
	"blog/config"
	"blog/internal/model"
	"blog/internal/repository"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockUserRepository 模拟用户仓库
type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) Create(user *model.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepository) FindByEmail(email string) (*model.User, error) {
	args := m.Called(email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.User), args.Error(1)
}

func (m *MockUserRepository) FindByID(id uint) (*model.User, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.User), args.Error(1)
}

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	return r
}

func generateTestToken(userID uint, secret string) string {
	claims := jwt.MapClaims{
		"user_id":  userID,
		"username": "testuser",
		"email":    "test@example.com",
		"role":     "user",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString([]byte(secret))
	return tokenString
}

func TestAuthMiddleware(t *testing.T) {
	// 初始化配置
	config.AppConfig.JWT.Secret = "test-secret"

	// 测试用例1：缺少认证头
	t.Run("缺少认证头", func(t *testing.T) {
		router := setupRouter()
		router.Use(AuthMiddleware())
		router.GET("/test", func(c *gin.Context) {
			c.Status(http.StatusOK)
		})

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/test", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})

	// 测试用例2：无效的认证头格式
	t.Run("无效的认证头格式", func(t *testing.T) {
		router := setupRouter()
		router.Use(AuthMiddleware())
		router.GET("/test", func(c *gin.Context) {
			c.Status(http.StatusOK)
		})

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/test", nil)
		req.Header.Set("Authorization", "InvalidFormat")
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})

	// 测试用例3：无效的token
	t.Run("无效的token", func(t *testing.T) {
		router := setupRouter()
		router.Use(AuthMiddleware())
		router.GET("/test", func(c *gin.Context) {
			c.Status(http.StatusOK)
		})

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/test", nil)
		req.Header.Set("Authorization", "Bearer invalid-token")
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})

	// 测试用例4：用户不存在
	t.Run("用户不存在", func(t *testing.T) {
		mockRepo := new(MockUserRepository)
		originalNewUserRepository := repository.NewUserRepository
		repository.NewUserRepository = func() repository.IUserRepository {
			return mockRepo
		}
		defer func() {
			repository.NewUserRepository = originalNewUserRepository
		}()

		tokenString := generateTestToken(1, config.AppConfig.JWT.Secret)

		router := setupRouter()
		router.Use(AuthMiddleware())
		router.GET("/test", func(c *gin.Context) {
			c.Status(http.StatusOK)
		})

		mockRepo.On("FindByID", uint(1)).Return(nil, nil)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/test", nil)
		req.Header.Set("Authorization", "Bearer "+tokenString)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})

	// 测试用例5：成功认证
	t.Run("成功认证", func(t *testing.T) {
		mockRepo := new(MockUserRepository)
		originalNewUserRepository := repository.NewUserRepository
		repository.NewUserRepository = func() repository.IUserRepository {
			return mockRepo
		}
		defer func() {
			repository.NewUserRepository = originalNewUserRepository
		}()

		tokenString := generateTestToken(1, config.AppConfig.JWT.Secret)

		user := &model.User{
			ID:       1,
			Username: "testuser",
			Email:    "test@example.com",
		}

		router := setupRouter()
		router.Use(AuthMiddleware())
		router.GET("/test", func(c *gin.Context) {
			userFromContext, exists := c.Get("user")
			assert.True(t, exists)
			assert.Equal(t, user, userFromContext)
			c.Status(http.StatusOK)
		})

		mockRepo.On("FindByID", uint(1)).Return(user, nil)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/test", nil)
		req.Header.Set("Authorization", "Bearer "+tokenString)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})
}
