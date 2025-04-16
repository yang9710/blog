export const API_BASE_URL = process.env.NEXT_PUBLIC_API_BASE_URL || 'http://localhost:8080';

export const API_URLS = {
  article: {
    create: `${API_BASE_URL}/api/v1/articles/create`,
    update: `${API_BASE_URL}/api/v1/articles/update`,
    delete: `${API_BASE_URL}/api/v1/articles/delete`,
    detail: `${API_BASE_URL}/api/v1/articles/detail`,
    list: `${API_BASE_URL}/api/v1/articles/list`,
  },
  auth: {
    register: `${API_BASE_URL}/api/v1/auth/register`,
    login: `${API_BASE_URL}/api/v1/auth/login`,
  }
};
