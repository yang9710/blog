import request from './request';
import { API_URLS } from './url';
import {
  Article,
  CreateArticleRequest,
  UpdateArticleRequest,
  ArticleListRequest,
  ApiResponse,
  ListResponse
} from '../types/article';
import { RegisterRequest, LoginRequest, LoginResponse } from '../types/auth';

// 文章相关服务
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
    await request.post(API_URLS.article.delete, { id });
  },

  // 获取文章详情
  async getById(id: number): Promise<Article> {
    const response = await request.post<ApiResponse<Article>>(
      API_URLS.article.detail,
      { id }
    );
    return response.data.data;
  },

  // 获取文章列表
  async list(params: ArticleListRequest): Promise<ListResponse<Article>> {
    const response = await request.post<ApiResponse<ListResponse<Article>>>(
      API_URLS.article.list,
      params
    );
    return response.data.data;
  }
};

// 认证相关服务
export const authService = {
  // 注册
  async register(data: RegisterRequest): Promise<void> {
    const response = await request.post<ApiResponse<void>>(
      API_URLS.auth.register,
      data
    );
    return response.data.data;
  },

  // 登录
  async login(data: LoginRequest): Promise<LoginResponse> {
    const response = await request.post<ApiResponse<LoginResponse>>(
      API_URLS.auth.login,
      data
    );
    const { token, user } = response.data.data;
    localStorage.setItem('token', token);
    return { token, user };
  },

  // 登出
  logout(): void {
    localStorage.removeItem('token');
  },
};
