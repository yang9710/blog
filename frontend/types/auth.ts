export interface User {
  id: number;
  username: string;
  email: string;
  role?: string;
  created_at?: string;
  updated_at?: string;
}

export interface RegisterRequest {
  username: string;
  email: string;
  password: string;
}

export interface LoginRequest {
  email: string;
  password: string;
}

export interface LoginResponse {
  token: string;
  user: User;
}

export interface AuthState {
  isAuthenticated: boolean;
  user: User | null;
  token: string | null;
  isLoading: boolean;
  error: string | null;
}

export interface ErrorResponse {
  error: string;
  message: string;
  statusCode: number;
}