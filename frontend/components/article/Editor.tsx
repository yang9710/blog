'use client';

import { useState, useEffect } from 'react';
import { useRouter } from 'next/navigation';
import dynamic from 'next/dynamic';
import { Article, CreateArticleRequest, UpdateArticleRequest } from '@/types/article';
import { articleService } from '@/services/article';

const MDEditor = dynamic(() => import('@uiw/react-md-editor'), { ssr: false });

interface EditorProps {
  article?: Article;
  isEdit?: boolean;
}

export default function Editor({ article, isEdit }: EditorProps) {
  const router = useRouter();
  const [title, setTitle] = useState(article?.title || '');
  const [content, setContent] = useState(article?.content || '');
  const [tags, setTags] = useState<string[]>(article?.tags.map(t => t.name) || []);
  const [status, setStatus] = useState<'draft' | 'published'>(article?.status || 'draft');
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const handleSubmit = async () => {
    if (!title.trim() || !content.trim()) {
      setError('标题和内容不能为空');
      return;
    }

    setLoading(true);
    setError(null);

    try {
      if (isEdit && article) {
        const updateData: UpdateArticleRequest = {
          id: article.id,
          title: title.trim(),
          content,
          status,
          tags: tags.filter(Boolean),
        };
        await articleService.update(updateData);
      } else {
        const createData: CreateArticleRequest = {
          title: title.trim(),
          content,
          status,
          tags: tags.filter(Boolean),
        };
        await articleService.create(createData);
      }
      router.push('/articles');
    } catch (err) {
      setError('保存失败，请重试');
      console.error('保存文章失败:', err);
    } finally {
      setLoading(false);
    }
  };

  const handleTagsChange = (value: string) => {
    const tagList = value.split(',').map(tag => tag.trim()).filter(Boolean);
    setTags(tagList);
  };

  return (
    <div className="max-w-5xl mx-auto py-8 px-4">
      <div className="space-y-8">
        {error && (
          <div className="bg-red-50 border-l-4 border-red-500 p-4 rounded-r-lg">
            <div className="flex">
              <div className="flex-shrink-0">
                <svg className="h-5 w-5 text-red-500" viewBox="0 0 20 20" fill="currentColor">
                  <path fillRule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zM8.707 7.293a1 1 0 00-1.414 1.414L8.586 10l-1.293 1.293a1 1 0 101.414 1.414L10 11.414l1.293 1.293a1 1 0 001.414-1.414L11.414 10l1.293-1.293a1 1 0 00-1.414-1.414L10 8.586 8.707 7.293z" clipRule="evenodd" />
                </svg>
              </div>
              <div className="ml-3">
                <p className="text-sm text-red-700">{error}</p>
              </div>
            </div>
          </div>
        )}

        <div className="bg-white rounded-2xl shadow-lg p-8 border border-gray-100">
          <input
            type="text"
            value={title}
            onChange={(e) => setTitle(e.target.value)}
            placeholder="输入文章标题..."
            className="w-full px-4 py-3 text-2xl font-semibold border-b-2 border-gray-100 focus:border-blue-500 focus:outline-none bg-transparent mb-8 transition-colors"
          />

          <div className="rounded-xl overflow-hidden border border-gray-100">
            <MDEditor
              value={content}
              onChange={(value) => setContent(value || '')}
              height={500}
              preview="edit"
              className="!border-0"
            />
          </div>
        </div>

        <div className="bg-white rounded-2xl shadow-lg p-8 border border-gray-100 space-y-6">
          <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div>
              <label className="block text-sm font-medium text-gray-700 mb-2">
                标签
              </label>
              <input
                type="text"
                value={tags.join(', ')}
                onChange={(e) => handleTagsChange(e.target.value)}
                placeholder="输入标签，用逗号分隔..."
                className="w-full px-4 py-3 border border-gray-200 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent transition-all"
              />
            </div>

            <div>
              <label className="block text-sm font-medium text-gray-700 mb-2">
                状态
              </label>
              <select
                value={status}
                onChange={(e) => setStatus(e.target.value as 'draft' | 'published')}
                className="w-full px-4 py-3 border border-gray-200 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent bg-white transition-all"
              >
                <option value="draft">草稿</option>
                <option value="published">发布</option>
              </select>
            </div>
          </div>

          <div className="flex justify-end space-x-4 pt-4">
            <button
              type="button"
              onClick={() => router.back()}
              className="px-6 py-3 border border-gray-200 rounded-lg text-gray-600 hover:bg-gray-50 transition-colors"
            >
              取消
            </button>
            <button
              type="button"
              onClick={handleSubmit}
              disabled={loading}
              className="px-6 py-3 bg-gradient-to-r from-blue-600 to-blue-400 text-white rounded-lg font-medium hover:opacity-90 disabled:opacity-50 transform hover:scale-105 transition-all duration-200 shadow-md disabled:cursor-not-allowed"
            >
              {loading ? (
                <span className="flex items-center">
                  <svg className="animate-spin -ml-1 mr-2 h-4 w-4 text-white" fill="none" viewBox="0 0 24 24">
                    <circle className="opacity-25" cx="12" cy="12" r="10" stroke="currentColor" strokeWidth="4" />
                    <path className="opacity-75" fill="currentColor" d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z" />
                  </svg>
                  保存中...
                </span>
              ) : (
                isEdit ? '更新文章' : '发布文章'
              )}
            </button>
          </div>
        </div>
      </div>
    </div>
  );
}