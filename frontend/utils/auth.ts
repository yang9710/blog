import api from './api';
import { RegisterRequest, LoginRequest, LoginResponse } from '../types/auth';

export const authService = {
  async register(data: RegisterRequest) {
    const response = await api.post('/auth/register', data);
    return response.data;
  },

  async login(data: LoginRequest): Promise<LoginResponse> {
    const response = await api.post('/auth/login', data);
    const { token, user } = response.data;
    localStorage.setItem('token', token);
    return { token, user };
  },

  logout() {
    localStorage.removeItem('token');
  },

  async getCurrentUser() {
    const response = await api.get('/auth/me');
    return response.data;
  },
};