'use client';

import LoginForm from '@/components/auth/LoginForm';
import { motion } from 'framer-motion';

export default function LoginPage() {
  return (
    <div className="min-h-screen gradient-bg flex flex-col justify-center py-12 sm:px-6 lg:px-8">
      <div className="sm:mx-auto sm:w-full sm:max-w-md">
        <motion.div
          initial={{ opacity: 0, y: -20 }}
          animate={{ opacity: 1, y: 0 }}
          transition={{ duration: 0.5 }}
          className="text-center"
        >
          <h2 className="mt-6 text-4xl font-extrabold text-white">
            欢迎回来
          </h2>
          <p className="mt-2 text-lg text-white text-opacity-80">
            登录您的账号继续使用
          </p>
        </motion.div>
      </div>
      <motion.div
        initial={{ opacity: 0, scale: 0.9 }}
        animate={{ opacity: 1, scale: 1 }}
        transition={{ duration: 0.5, delay: 0.2 }}
        className="mt-8 sm:mx-auto sm:w-full sm:max-w-md"
      >
        <div className="glass py-8 px-4 shadow-xl sm:rounded-lg sm:px-10">
          <LoginForm />
        </div>
      </motion.div>
      <div className="absolute bottom-4 right-4 text-white text-opacity-60 text-sm">
        © {new Date().getFullYear()} 博客系统
      </div>
    </div>
  );
}