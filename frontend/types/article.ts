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
  author: Author;
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

export interface UpdateArticleRequest extends CreateArticleRequest {
  id: number;
}

export interface ListArticleRequest {
  page: number;
  page_size: number;
  status?: string;
  author_id?: number;
  tag?: string;
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