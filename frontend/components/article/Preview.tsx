import React from 'react';
import { Article } from '../../types/article';
import dynamic from 'next/dynamic';
import { Tag, Space } from 'antd';
import { format } from 'date-fns';

const MDPreview = dynamic(() => import('@uiw/react-md-editor').then(mod => mod.default.Markdown), { ssr: false });

interface PreviewProps {
  article: Article;
}

export const ArticlePreview: React.FC<PreviewProps> = ({ article }) => {
  return (
    <div className="max-w-4xl mx-auto p-4">
      <article>
        <header className="mb-8">
          <h1 className="text-3xl font-bold mb-4">{article.title}</h1>
          <div className="flex justify-between items-center text-gray-600">
            <Space size="middle">
              <span>作者：{article.author.username}</span>
              <span>
                发布时间：{format(new Date(article.created_at), 'yyyy-MM-dd HH:mm:ss')}
              </span>
              <Tag color={article.status === 'published' ? 'green' : 'orange'}>
                {article.status === 'published' ? '已发布' : '草稿'}
              </Tag>
            </Space>
            <div>
              {article.tags.map(tag => (
                <Tag key={tag.name} className="mr-2">
                  {tag.name}
                </Tag>
              ))}
            </div>
          </div>
        </header>

        <div className="prose max-w-none" data-color-mode="light">
          <MDPreview source={article.content} />
        </div>

        <footer className="mt-8 pt-4 border-t text-gray-600">
          <p>最后更新：{format(new Date(article.updated_at), 'yyyy-MM-dd HH:mm:ss')}</p>
        </footer>
      </article>
    </div>
  );
};