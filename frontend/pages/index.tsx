import { useRouter } from 'next/router';
import { useEffect } from 'react';
import { useAuth } from '../contexts/AuthContext';
import { motion } from 'framer-motion';

export default function Home() {
  const router = useRouter();
  const { isAuthenticated } = useAuth();

  useEffect(() => {
    if (!isAuthenticated) {
      router.push('/auth/login');
    }
  }, [isAuthenticated, router]);

  return (
    <div className="min-h-screen bg-gradient-to-br from-gray-50 to-gray-100">
      <main className="container mx-auto px-4 py-8">
        <motion.div
          initial={{ opacity: 0, y: 20 }}
          animate={{ opacity: 1, y: 0 }}
          transition={{ duration: 0.5 }}
          className="max-w-4xl mx-auto"
        >
          <h1 className="text-4xl font-bold text-gray-900 mb-6 text-center">
            欢迎来到博客系统
          </h1>
          <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
            <motion.div
              whileHover={{ scale: 1.02 }}
              className="card p-6"
            >
              <h2 className="text-2xl font-semibold text-gray-800 mb-4">
                开始写作
              </h2>
              <p className="text-gray-600">
                创建您的第一篇博客文章，分享您的想法和见解。
              </p>
            </motion.div>
            <motion.div
              whileHover={{ scale: 1.02 }}
              className="card p-6"
            >
              <h2 className="text-2xl font-semibold text-gray-800 mb-4">
                探索社区
              </h2>
              <p className="text-gray-600">
                发现其他作者的精彩内容，参与讨论和交流。
              </p>
            </motion.div>
          </div>
          <motion.div
            initial={{ opacity: 0 }}
            animate={{ opacity: 1 }}
            transition={{ delay: 0.3 }}
            className="mt-8 text-center"
          >
            <button className="btn btn-primary">
              开始创作
            </button>
          </motion.div>
        </motion.div>
      </main>
    </div>
  );
}