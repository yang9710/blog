import axios from 'axios';
import {
  Article,
  CreateArticleRequest,
  UpdateArticleRequest,
  ListArticleRequest,
  ApiResponse,
  ListResponse
} from '../types/article';

const API_BASE_URL = process.env.NEXT_PUBLIC_API_BASE_URL || 'http://localhost:8080';

export const articleService = {
  // 创建文章
  async create(data: CreateArticleRequest): Promise<Article> {
    const response = await axios.post<ApiResponse<Article>>(
      `${API_BASE_URL}/api/v1/articles/create`,
      data
    );
    return response.data.data;
  },

  // 更新文章
  async update(data: UpdateArticleRequest): Promise<Article> {
    const response = await axios.post<ApiResponse<Article>>(
      `${API_BASE_URL}/api/v1/articles/update`,
      data
    );
    return response.data.data;
  },

  // 删除文章
  async delete(id: number): Promise<void> {
    await axios.post(`${API_BASE_URL}/api/v1/articles/delete`, { id });
  },

  // 获取文章详情
  async getById(id: number): Promise<Article> {
    const response = await axios.post<ApiResponse<Article>>(
      `${API_BASE_URL}/api/v1/articles/detail`,
      { id }
    );
    return response.data.data;
  },

  // 获取文章列表
  async list(params: ListArticleRequest): Promise<ListResponse<Article>> {
    const response = await axios.post<ApiResponse<ListResponse<Article>>>(
      `${API_BASE_URL}/api/v1/articles/list`,
      params
    );
    return response.data.data;
  }
};