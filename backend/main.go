package main

import (
	"blog/config"
	"blog/internal/handler"
	"blog/internal/middleware"
	"blog/internal/repository"
	"blog/internal/service"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// 加载环境变量
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// 加载配置
	config.LoadConfig()

	// 初始化数据库
	repository.InitDB()

	// 创建 Gin 引擎
	r := gin.Default()

	// 添加全局中间件
	r.Use(middleware.CORSMiddleware())

	// 初始化服务和处理器
	authService := service.NewAuthService()
	articleService := service.NewArticleService()

	authHandler := handler.NewAuthHandler(authService)
	articleHandler := handler.NewArticleHandler(articleService)

	// 注册路由
	api := r.Group("/api/v1")
	{
		// 认证相关路由（无需认证）
		auth := api.Group("/auth")
		{
			auth.POST("/register", authHandler.Register)
			auth.POST("/login", authHandler.Login)
		}

		// 需要认证的路由
		authenticated := api.Group("")
		authenticated.Use(middleware.AuthMiddleware())
		{
			// 文章相关路由
			articles := authenticated.Group("/articles")
			{
				articles.POST("/create", articleHandler.CreateArticle)
				articles.POST("/update", articleHandler.UpdateArticle)
				articles.POST("/delete", articleHandler.DeleteArticle)
				articles.POST("/detail", articleHandler.GetArticle)
				articles.POST("/list", articleHandler.ListArticles)
			}
		}
	}

	// 启动服务器
	log.Printf("Server starting on %s", config.AppConfig.Server.Port)
	if err := r.Run(config.AppConfig.Server.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}