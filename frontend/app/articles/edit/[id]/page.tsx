'use client';

import { useEffect, useState } from 'react';
import { useParams, useRouter } from 'next/navigation';
import { Article } from '@/types/article';
import { articleService } from '@/services/article';
import { useAuthStore } from '@/stores/auth';
import Layout from '@/components/common/Layout';
import Editor from '@/components/article/Editor';

export default function EditArticlePage() {
  const params = useParams();
  const router = useRouter();
  const { user } = useAuthStore();
  const [article, setArticle] = useState<Article | null>(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchArticle = async () => {
      try {
        const data = await articleService.getDetail(Number(params.id));
        if (data.author.id !== user?.id) {
          router.push('/articles');
          return;
        }
        setArticle(data);
      } catch (error) {
        console.error('获取文章失败:', error);
        router.push('/articles');
      } finally {
        setLoading(false);
      }
    };

    fetchArticle();
  }, [params.id, router, user?.id]);

  if (loading) {
    return (
      <Layout>
        <div className="text-center py-8">加载中...</div>
      </Layout>
    );
  }

  if (!article) {
    return (
      <Layout>
        <div className="text-center py-8">文章不存在</div>
      </Layout>
    );
  }

  return (
    <Layout>
      <div className="max-w-4xl mx-auto">
        <div className="flex items-center mb-6">
          <button
            onClick={() => router.push('/articles')}
            className="flex items-center text-gray-600 hover:text-gray-900 transition-colors"
          >
            <svg className="w-5 h-5 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="M10 19l-7-7m0 0l7-7m-7 7h18" />
            </svg>
            返回列表
          </button>
        </div>
        <Editor article={article} isEdit />
      </div>
    </Layout>
  );
}