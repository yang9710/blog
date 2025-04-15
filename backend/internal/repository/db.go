package repository

import (
	"blog/config"
	"blog/internal/model"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db           *gorm.DB
	userRepo     *UserRepository
	articleRepo  *ArticleRepository
)

// InitDB 初始化数据库连接
func InitDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.AppConfig.Database.User,
		config.AppConfig.Database.Password,
		config.AppConfig.Database.Host,
		config.AppConfig.Database.Port,
		config.AppConfig.Database.DBName,
	)

	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// 设置连接池
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get database instance: %v", err)
	}

	// 设置最大连接数
	sqlDB.SetMaxOpenConns(100)
	// 设置最大空闲连接数
	sqlDB.SetMaxIdleConns(10)

	// 自动迁移数据库表
	err = db.AutoMigrate(
		&model.User{},
		&model.Article{},
		&model.Tag{},
	)
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// 初始化仓储实例
	userRepo = &UserRepository{db: db}
	articleRepo = &ArticleRepository{db: db}
}

// IUserRepository 用户仓库接口
type IUserRepository interface {
	Create(user *model.User) error
	FindByEmail(email string) (*model.User, error)
	FindByID(id uint) (*model.User, error)
}

// IArticleRepository 文章仓库接口
type IArticleRepository interface {
	Create(article *model.Article) error
	Update(article *model.Article) error
	Delete(id uint, authorID uint) error
	FindByID(id uint) (*model.Article, error)
	List(page, pageSize int, status string, authorID uint, tag string) ([]model.Article, int64, error)
	UpdateTags(article *model.Article, tags []string) error
}

// NewUserRepository 创建用户仓库的函数类型
type NewUserRepositoryFunc func() IUserRepository

// NewArticleRepository 创建文章仓库的函数类型
type NewArticleRepositoryFunc func() IArticleRepository

// NewUserRepository 创建用户仓库的默认实现
var NewUserRepository NewUserRepositoryFunc = func() IUserRepository {
	if userRepo == nil {
		userRepo = &UserRepository{db: db}
	}
	return userRepo
}

// NewArticleRepository 创建文章仓库的默认实现
var NewArticleRepository NewArticleRepositoryFunc = func() IArticleRepository {
	if articleRepo == nil {
		articleRepo = &ArticleRepository{db: db}
	}
	return articleRepo
}

// GetDB 获取数据库连接
func GetDB() *gorm.DB {
	return db
}

// GetUserRepository 获取用户仓库
func GetUserRepository() *UserRepository {
	if userRepo == nil {
		userRepo = &UserRepository{db: db}
	}
	return userRepo
}

// GetArticleRepository 获取文章仓库
func GetArticleRepository() *ArticleRepository {
	if articleRepo == nil {
		articleRepo = &ArticleRepository{db: db}
	}
	return articleRepo
}