import { create } from 'zustand';
import { User } from '@/types/auth';
import { authService } from '@/services';

interface AuthState {
  isAuthenticated: boolean;
  user: User | null;
  loading: boolean;
  login: (email: string, password: string) => Promise<void>;
  logout: () => Promise<void>;
  initialize: () => Promise<void>;
}

export const useAuthStore = create<AuthState>((set) => ({
  isAuthenticated: false,
  user: null,
  loading: true,

  initialize: async () => {
    try {
      const token = localStorage.getItem('token');
      const user = localStorage.getItem('user');
      if (!token || !user) {
        set({ loading: false });
        return;
      }

      set({
        user: user ? JSON.parse(user) : null,
        isAuthenticated: true,
        loading: false,
      });
    } catch (error) {
      console.error('Failed to get user info:', error);
      localStorage.removeItem('token');
      set({
        user: null,
        isAuthenticated: false,
        loading: false,
      });
    }
  },

  login: async (email: string, password: string) => {
    try {
      const { user, token } = await authService.login({ email, password });
      localStorage.setItem('token', token);
      localStorage.setItem('user', JSON.stringify(user));
      set({
        user,
        isAuthenticated: true,
      });
    } catch (error) {
      throw error;
    }
  },

  logout: async () => {
    localStorage.removeItem('token');
    localStorage.removeItem('user');
    set({
      user: null,
      isAuthenticated: false,
    });
  },
}));