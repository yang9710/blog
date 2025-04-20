import axios, { AxiosInstance, AxiosRequestConfig } from 'axios';
// import { message } from 'antd';

// 创建 axios 实例
const request: AxiosInstance = axios.create({
  timeout: 10000,
  headers: {
    'Content-Type': 'application/json',
  },
});

// 请求拦截器
request.interceptors.request.use(
  (config) => {
    // 从 localStorage 获取 token
    const token = localStorage.getItem('token');
    if (token && config.headers) {
      config.headers.Authorization = `Bearer ${token}`;
    }
    return config;
  },
  (error) => {
    return Promise.reject(error);
  }
);

// 响应拦截器
request.interceptors.response.use(
  (response) => {
    const { code, message: msg, data } = response.data;

    // 如果后端返回的不是成功状态码
    if (code !== 200 && code !== 201) {
      // message.error(msg || '请求失败');
      return Promise.reject(new Error(msg || '请求失败'));
    }

    return response;
  },
  (error) => {
    // 处理 401 未授权
    if (error.response?.status === 401) {
      localStorage.removeItem('token');
      window.location.href = '/login';
      return Promise.reject(new Error('请先登录'));
    }

    // 处理其他错误
    // message.error(error.response?.data?.message || '请求失败');
    return Promise.reject(error);
  }
);

export default request;