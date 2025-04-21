import { User } from './auth';

export interface Tag {
  id: number;
  name: string;
}

export interface Author {
  id: number;
  username: string;
}

export interface Article {
  id: number;
  title: string;
  content: string;
  status: 'draft' | 'published';
  author: {
    id: number;
    username: string;
  };
  tags: Tag[];
  created_at: string;
  updated_at: string;
}

export interface CreateArticleRequest {
  title: string;
  content: string;
  status: 'draft' | 'published';
  tags: string[];
}

export interface UpdateArticleRequest {
  id: number;
  title: string;
  content: string;
  status: 'draft' | 'published';
  tags: string[];
}

export interface ArticleListRequest {
  page: number;
  page_size: number;
  status?: string;
  author_id?: number;
  tag?: string;
}

export interface ArticleListResponse {
  total: number;
  articles: Article[];
}

export interface ApiResponse<T> {
  code: number;
  message: string;
  data: T;
}

export interface ListResponse<T> {
  total: number;
  items: T[];
}