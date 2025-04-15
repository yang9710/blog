package handler

import (
	"blog/internal/model"
	"blog/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ArticleHandler struct {
	articleService *service.ArticleService
}

func NewArticleHandler(articleService *service.ArticleService) *ArticleHandler {
	return &ArticleHandler{articleService: articleService}
}

// 请求结构体
type CreateArticleRequest struct {
	Title   string   `json:"title" binding:"required"`
	Content string   `json:"content" binding:"required"`
	Status  string   `json:"status" binding:"required,oneof=draft published"`
	Tags    []string `json:"tags"`
}

type UpdateArticleRequest struct {
	ID      uint     `json:"id" binding:"required"`
	Title   string   `json:"title" binding:"required"`
	Content string   `json:"content" binding:"required"`
	Status  string   `json:"status" binding:"required,oneof=draft published"`
	Tags    []string `json:"tags"`
}

type ListArticleRequest struct {
	Page     int    `json:"page" binding:"required,min=1"`
	PageSize int    `json:"page_size" binding:"required,min=1,max=100"`
	Status   string `json:"status"`
	AuthorID uint   `json:"author_id"`
	Tag      string `json:"tag"`
}

// 响应结构体
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// CreateArticle 创建文章
func (h *ArticleHandler) CreateArticle(c *gin.Context) {
	var req CreateArticleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Code:    400,
			Message: "参数错误: " + err.Error(),
		})
		return
	}

	// 获取当前用户
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, Response{
			Code:    401,
			Message: "未登录或登录已过期",
		})
		return
	}
	currentUser := user.(*model.User)

	// 构建文章对象
	article := &model.Article{
		Title:    req.Title,
		Content:  req.Content,
		Status:   req.Status,
		AuthorID: currentUser.ID,
	}

	// 处理标签
	for _, tagName := range req.Tags {
		article.Tags = append(article.Tags, model.Tag{Name: tagName})
	}

	// 创建文章
	if err := h.articleService.CreateArticle(article); err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Code:    500,
			Message: "创建文章失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, Response{
		Code:    201,
		Message: "创建成功",
		Data:    article,
	})
}

// UpdateArticle 更新文章
func (h *ArticleHandler) UpdateArticle(c *gin.Context) {
	var req UpdateArticleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Code:    400,
			Message: "参数错误: " + err.Error(),
		})
		return
	}

	// 构建文章对象
	article := &model.Article{
		ID:      req.ID,
		Title:   req.Title,
		Content: req.Content,
		Status:  req.Status,
	}

	// 处理标签
	for _, tagName := range req.Tags {
		article.Tags = append(article.Tags, model.Tag{Name: tagName})
	}

	// 更新文章
	if err := h.articleService.UpdateArticle(article); err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Code:    500,
			Message: "更新文章失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "更新成功",
		Data:    article,
	})
}

// DeleteArticle 删除文章
func (h *ArticleHandler) DeleteArticle(c *gin.Context) {
	var req struct {
		ID uint `json:"id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Code:    400,
			Message: "参数错误: " + err.Error(),
		})
		return
	}

	// 获取当前用户
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, Response{
			Code:    401,
			Message: "未登录或登录已过期",
		})
		return
	}
	currentUser := user.(*model.User)

	// 删除文章
	if err := h.articleService.DeleteArticle(req.ID, currentUser.ID); err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Code:    500,
			Message: "删除文章失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "删除成功",
	})
}

// GetArticle 获取文章详情
func (h *ArticleHandler) GetArticle(c *gin.Context) {
	var req struct {
		ID uint `json:"id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Code:    400,
			Message: "参数错误: " + err.Error(),
		})
		return
	}

	// 获取文章
	article, err := h.articleService.GetArticle(req.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Code:    500,
			Message: "获取文章失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "获取成功",
		Data:    article,
	})
}

// ListArticles 获取文章列表
func (h *ArticleHandler) ListArticles(c *gin.Context) {
	var req ListArticleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Code:    400,
			Message: "参数错误: " + err.Error(),
		})
		return
	}

	// 获取文章列表
	articles, total, err := h.articleService.ListArticles(
		req.Page,
		req.PageSize,
		req.Status,
		req.AuthorID,
		req.Tag,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Code:    500,
			Message: "获取文章列表失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "获取成功",
		Data: gin.H{
			"total":    total,
			"articles": articles,
		},
	})
}

// RegisterRoutes 注册路由
func (h *ArticleHandler) RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api/v1/articles")
	{
		api.POST("/create", h.CreateArticle)
		api.POST("/update", h.UpdateArticle)
		api.POST("/delete", h.DeleteArticle)
		api.POST("/detail", h.GetArticle)
		api.POST("/list", h.ListArticles)
	}
}