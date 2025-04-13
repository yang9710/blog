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

	r.Use(middleware.CORSMiddleware())

	// 创建处理器
	authHandler := handler.NewAuthHandler(service.NewAuthService())

	// 注册路由
	api := r.Group("/api/v1")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/register", authHandler.Register)
			auth.POST("/login", authHandler.Login)
		}

		// 需要认证的路由
		authenticated := api.Group("")
		authenticated.Use(middleware.AuthMiddleware())
		{
			// 在这里添加需要认证的路由
		}
	}

	// 启动服务器
	log.Printf("Server starting on %s", config.AppConfig.Server.Port)
	if err := r.Run(config.AppConfig.Server.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
