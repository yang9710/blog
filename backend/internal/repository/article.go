package repository

import (
	"blog/internal/model"

	"gorm.io/gorm"
)

// ArticleRepository 实现 IArticleRepository 接口
type ArticleRepository struct {
	db *gorm.DB
}

// Create 创建文章
func (r *ArticleRepository) Create(article *model.Article) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		// 处理标签
		var tags []model.Tag
		for _, tagName := range article.Tags {
			var tag model.Tag
			// 先尝试查找标签
			result := tx.Where("name = ?", tagName.Name).First(&tag)
			if result.Error == gorm.ErrRecordNotFound {
				// 如果标签不存在，创建新标签
				tag = model.Tag{Name: tagName.Name}
				if err := tx.Create(&tag).Error; err != nil {
					return err
				}
			} else if result.Error != nil {
				return result.Error
			}
			tags = append(tags, tag)
		}

		// 清空原有标签列表
		article.Tags = nil

		// 创建文章
		if err := tx.Create(article).Error; err != nil {
			return err
		}

		// 建立标签关联
		if len(tags) > 0 {
			if err := tx.Model(article).Association("Tags").Replace(tags); err != nil {
				return err
			}
		}

		// 重新加载文章信息（包括标签）
		if err := tx.Preload("Tags").First(article, article.ID).Error; err != nil {
			return err
		}

		return nil
	})
}

// Update 更新文章
func (r *ArticleRepository) Update(article *model.Article) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		// 更新文章基本信息
		if err := tx.Model(article).Updates(map[string]interface{}{
			"title":   article.Title,
			"content": article.Content,
			"status":  article.Status,
		}).Error; err != nil {
			return err
		}

		// 如果有标签需要更新
		if len(article.Tags) > 0 {
			// 先清除原有标签关联
			if err := tx.Model(article).Association("Tags").Clear(); err != nil {
				return err
			}

			// 处理每个标签
			var tags []model.Tag
			for _, tagInfo := range article.Tags {
				var tag model.Tag
				// 先尝试查找标签
				result := tx.Where("name = ?", tagInfo.Name).First(&tag)
				if result.Error == gorm.ErrRecordNotFound {
					// 如果标签不存在，创建新标签
					tag = model.Tag{Name: tagInfo.Name}
					if err := tx.Create(&tag).Error; err != nil {
						return err
					}
				} else if result.Error != nil {
					return result.Error
				}
				tags = append(tags, tag)
			}

			// 建立新的标签关联
			if err := tx.Model(article).Association("Tags").Replace(tags); err != nil {
				return err
			}
		}

		// 重新加载文章信息（包括标签）
		if err := tx.Preload("Tags").First(article, article.ID).Error; err != nil {
			return err
		}

		return nil
	})
}

// Delete 删除文章（软删除）
func (r *ArticleRepository) Delete(id uint, authorID uint) error {
	return r.db.Where("id = ? AND author_id = ?", id, authorID).Delete(&model.Article{}).Error
}

// FindByID 通过ID查找文章
func (r *ArticleRepository) FindByID(id uint) (*model.Article, error) {
	var article model.Article
	err := r.db.Preload("Author").Preload("Tags").First(&article, id).Error
	if err != nil {
		return nil, err
	}
	return &article, nil
}

// List 获取文章列表
func (r *ArticleRepository) List(page, pageSize int, status string, authorID uint, tag string) ([]model.Article, int64, error) {
	var articles []model.Article
	var total int64

	query := r.db.Model(&model.Article{})

	// 添加查询条件
	if status != "" {
		query = query.Where("status = ?", status)
	}
	if authorID != 0 {
		query = query.Where("author_id = ?", authorID)
	}
	if tag != "" {
		query = query.Joins("JOIN article_tags ON articles.id = article_tags.article_id").
			Joins("JOIN tags ON article_tags.tag_id = tags.id").
			Where("tags.name = ?", tag)
	}

	// 获取总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 获取分页数据
	err := query.Preload("Author").Preload("Tags").
		Offset((page - 1) * pageSize).
		Limit(pageSize).
		Order("created_at DESC").
		Find(&articles).Error

	if err != nil {
		return nil, 0, err
	}

	return articles, total, nil
}

// UpdateTags 更新文章和标签
func (r *ArticleRepository) UpdateTags(article *model.Article, tags []string) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		// 更新文章基本信息
		if err := tx.Model(article).Updates(map[string]interface{}{
			"title":   article.Title,
			"content": article.Content,
			"status":  article.Status,
		}).Error; err != nil {
			return err
		}

		// 清除原有标签关联
		if err := tx.Model(article).Association("Tags").Clear(); err != nil {
			return err
		}

		// 处理新标签
		for _, tagName := range tags {
			var tag model.Tag
			// 先尝试查找标签，如果不存在则创建
			if err := tx.Where("name = ?", tagName).FirstOrCreate(&tag, model.Tag{Name: tagName}).Error; err != nil {
				return err
			}
			// 建立标签关联
			if err := tx.Model(article).Association("Tags").Append(&tag); err != nil {
				return err
			}
		}

		// 重新加载文章信息（包括标签）
		if err := tx.Preload("Tags").First(article, article.ID).Error; err != nil {
			return err
		}

		return nil
	})
}