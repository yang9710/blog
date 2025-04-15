package service

import (
	"blog/internal/model"
	"blog/internal/repository"
	"errors"
)

type ArticleService struct {
	articleRepo repository.IArticleRepository
}

func NewArticleService() *ArticleService {
	return &ArticleService{
		articleRepo: repository.NewArticleRepository(),
	}
}

// CreateArticle 创建文章
func (s *ArticleService) CreateArticle(article *model.Article) error {
	return s.articleRepo.Create(article)
}

// UpdateArticle 更新文章
func (s *ArticleService) UpdateArticle(article *model.Article) error {
	// 检查文章是否存在
	existingArticle, err := s.articleRepo.FindByID(article.ID)
	if err != nil {
		return err
	}
	if existingArticle == nil {
		return errors.New("文章不存在")
	}

	// 确保作者ID不变
	article.AuthorID = existingArticle.AuthorID
	return s.articleRepo.Update(article)
}

// DeleteArticle 删除文章（软删除）
func (s *ArticleService) DeleteArticle(id uint, authorID uint) error {
	// 检查文章是否存在且属于该作者
	article, err := s.articleRepo.FindByID(id)
	if err != nil {
		return err
	}
	if article == nil || article.AuthorID != authorID {
		return errors.New("文章不存在或无权限删除")
	}

	return s.articleRepo.Delete(id, authorID)
}

// GetArticle 获取文章详情
func (s *ArticleService) GetArticle(id uint) (*model.Article, error) {
	article, err := s.articleRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if article == nil {
		return nil, errors.New("文章不存在")
	}
	return article, nil
}

// ListArticles 获取文章列表
func (s *ArticleService) ListArticles(page, pageSize int, status string, authorID uint, tag string) ([]model.Article, int64, error) {
	return s.articleRepo.List(page, pageSize, status, authorID, tag)
}