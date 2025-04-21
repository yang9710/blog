'use client';

import { useEffect, useState } from 'react';
import { useParams, useRouter } from 'next/navigation';
import { Article } from '@/types/article';
import { articleService } from '@/services/article';
import { useAuthStore } from '@/stores/auth';
import dynamic from 'next/dynamic';
import Layout from '@/components/common/Layout';

const MDPreview = dynamic(() => import('@uiw/react-md-editor'), {
  ssr: false,
  loading: () => <div className="text-center py-8">加载中...</div>,
});

export default function ArticleDetailPage() {
  const params = useParams();
  const router = useRouter();
  const { user } = useAuthStore();
  const [article, setArticle] = useState<Article | null>(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchArticle = async () => {
      try {
        const data = await articleService.getDetail(Number(params.id));
        setArticle(data);
      } catch (error) {
        console.error('获取文章失败:', error);
      } finally {
        setLoading(false);
      }
    };

    fetchArticle();
  }, [params.id]);

  const handleDelete = async () => {
    if (!article || !confirm('确定要删除这篇文章吗？')) return;
    try {
      await articleService.delete(article.id);
      router.push('/articles');
    } catch (error) {
      console.error('删除文章失败:', error);
    }
  };

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

  const isAuthor = user?.id === article.author.id;

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

        <div className="mb-8">
          <h1 className="text-3xl font-bold mb-4">{article.title}</h1>
          <div className="flex items-center space-x-4 text-gray-500 mb-4">
            <span>作者：{article.author.username}</span>
            <span>
              发布时间：{new Date(article.created_at).toLocaleDateString()}
            </span>
            <span>
              状态：{article.status === 'published' ? '已发布' : '草稿'}
            </span>
          </div>
          {article.tags.length > 0 && (
            <div className="flex space-x-2 mb-4">
              {article.tags.map((tag) => (
                <span
                  key={tag.id}
                  className="px-3 py-1 bg-gray-100 rounded-full text-sm"
                >
                  {tag.name}
                </span>
              ))}
            </div>
          )}
          {isAuthor && (
            <div className="flex space-x-4 mb-8">
              <button
                onClick={() => router.push(`/articles/edit/${article.id}`)}
                className="px-4 py-2 bg-blue-500 text-white rounded hover:bg-blue-600"
              >
                编辑
              </button>
              <button
                onClick={handleDelete}
                className="px-4 py-2 bg-red-500 text-white rounded hover:bg-red-600"
              >
                删除
              </button>
            </div>
          )}
        </div>

        <div className="prose max-w-none">
          <MDPreview value={article.content} preview="preview" />
        </div>
      </div>
    </Layout>
  );
}