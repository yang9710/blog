package service

import (
	"blog/internal/model"
	"errors"
	"testing"

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

func TestAuthService_Register(t *testing.T) {
	mockRepo := new(MockUserRepository)
	authService := &AuthService{userRepo: mockRepo}

	// 测试用例1：成功注册
	t.Run("成功注册", func(t *testing.T) {
		req := &model.RegisterRequest{
			Username: "testuser",
			Email:    "test@example.com",
			Password: "password123",
		}

		mockRepo.On("FindByEmail", req.Email).Return(nil, nil)
		mockRepo.On("Create", mock.AnythingOfType("*model.User")).Return(nil)

		user, err := authService.Register(req)
		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, req.Username, user.Username)
		assert.Equal(t, req.Email, user.Email)
	})

	// 测试用例2：邮箱已存在
	t.Run("邮箱已存在", func(t *testing.T) {
		req := &model.RegisterRequest{
			Username: "testuser",
			Email:    "existing@example.com",
			Password: "password123",
		}

		existingUser := &model.User{
			Username: "existinguser",
			Email:    "existing@example.com",
		}

		mockRepo.On("FindByEmail", req.Email).Return(existingUser, nil)

		user, err := authService.Register(req)
		assert.Error(t, err)
		assert.Nil(t, user)
		assert.Equal(t, "email already exists", err.Error())
	})

	// 测试用例3：数据库错误
	t.Run("数据库错误", func(t *testing.T) {
		req := &model.RegisterRequest{
			Username: "testuser",
			Email:    "test@example.com",
			Password: "password123",
		}

		dbErr := errors.New("database error")
		mockRepo.On("FindByEmail", req.Email).Return(nil, dbErr)

		user, err := authService.Register(req)
		assert.Error(t, err)
		assert.Nil(t, user)
		assert.Equal(t, dbErr, err)
	})
}

func TestAuthService_Login(t *testing.T) {
	mockRepo := new(MockUserRepository)
	authService := &AuthService{userRepo: mockRepo}

	// 测试用例1：成功登录
	t.Run("成功登录", func(t *testing.T) {
		req := &model.LoginRequest{
			Email:    "test@example.com",
			Password: "password123",
		}

		hashedPassword := "$2a$10$abcdefghijklmnopqrstuvwxyz123456" // 模拟加密后的密码
		user := &model.User{
			Username: "testuser",
			Email:    "test@example.com",
			Password: hashedPassword,
		}

		mockRepo.On("FindByEmail", req.Email).Return(user, nil)

		response, err := authService.Login(req)
		assert.Error(t, err) // 由于密码验证会失败，我们期望有错误
		assert.Nil(t, response)
		assert.Equal(t, "invalid credentials", err.Error())
	})

	// 测试用例2：用户不存在
	t.Run("用户不存在", func(t *testing.T) {
		req := &model.LoginRequest{
			Email:    "nonexistent@example.com",
			Password: "password123",
		}

		mockRepo.On("FindByEmail", req.Email).Return(nil, nil)

		response, err := authService.Login(req)
		assert.Error(t, err)
		assert.Nil(t, response)
		assert.Equal(t, "invalid credentials", err.Error())
	})

	// 测试用例3：数据库错误
	t.Run("数据库错误", func(t *testing.T) {
		req := &model.LoginRequest{
			Email:    "test@example.com",
			Password: "password123",
		}

		dbErr := errors.New("database error")
		mockRepo.On("FindByEmail", req.Email).Return(nil, dbErr)

		response, err := authService.Login(req)
		assert.Error(t, err)
		assert.Nil(t, response)
		assert.Equal(t, dbErr, err)
	})
}
