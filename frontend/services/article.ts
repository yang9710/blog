import request from '@/utils/api';
import { API_URLS } from './url';
import type {
  Article,
  CreateArticleRequest,
  UpdateArticleRequest,
  ArticleListRequest,
  ArticleListResponse,
  ApiResponse,
  ListResponse,
} from '@/types/article';

export const articleService = {
  // 创建文章
  async create(data: CreateArticleRequest): Promise<Article> {
    const response = await request.post<ApiResponse<Article>>(
      API_URLS.article.create,
      data
    );
    return response.data.data;
  },

  // 更新文章
  async update(data: UpdateArticleRequest): Promise<Article> {
    const response = await request.post<ApiResponse<Article>>(
      API_URLS.article.update,
      data
    );
    return response.data.data;
  },

  // 删除文章
  async delete(id: number): Promise<void> {
    const response = await request.post<ApiResponse<void>>(
      API_URLS.article.delete,
      { id }
    );
    return response.data.data;
  },

  // 获取文章详情
  async getDetail(id: number): Promise<Article> {
    const response = await request.post<ApiResponse<Article>>(
      API_URLS.article.detail,
      { id }
    );
    return response.data.data;
  },

  // 获取文章列表
  async getList(params: ArticleListRequest): Promise<ArticleListResponse> {
    const response = await request.post<ApiResponse<ArticleListResponse>>(
      API_URLS.article.list,
      params
    );
    return response.data.data;
  },
};