package model

import (
	"time"

	"gorm.io/gorm"
)

// Article 文章模型
type Article struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	Title     string         `gorm:"type:varchar(200);not null" json:"title"`
	Content   string         `gorm:"type:text;not null" json:"content"`
	Status    string         `gorm:"type:varchar(20);default:draft" json:"status"`
	AuthorID  uint          `gorm:"not null" json:"author_id"`
	Author    User          `gorm:"foreignKey:AuthorID" json:"author"`
	Tags      []Tag         `gorm:"many2many:article_tags;" json:"tags"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}

// Tag 标签模型
type Tag struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	Name      string    `gorm:"type:varchar(50);unique;not null" json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

// ArticleRequest 文章请求基础结构
type ArticleRequest struct {
	Title   string   `json:"title" binding:"required,min=1,max=200"`
	Content string   `json:"content" binding:"required"`
	Status  string   `json:"status" binding:"required,oneof=draft published"`
	Tags    []string `json:"tags"`
}

// UpdateArticleRequest 更新文章请求
type UpdateArticleRequest struct {
	ID uint `json:"id" binding:"required"`
	ArticleRequest
}

// ListArticleRequest 文章列表请求
type ListArticleRequest struct {
	Page     int    `json:"page" binding:"required,min=1"`
	PageSize int    `json:"page_size" binding:"required,min=1,max=100"`
	Status   string `json:"status"`
	AuthorID uint   `json:"author_id"`
	Tag      string `json:"tag"`
}

// TableName 指定文章表名
func (Article) TableName() string {
	return "articles"
}

// TableName 指定标签表名
func (Tag) TableName() string {
	return "tags"
}