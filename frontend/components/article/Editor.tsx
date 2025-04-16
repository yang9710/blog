import React, { useState, useEffect } from 'react';
import { useRouter } from 'next/router';
import { Article, CreateArticleRequest, UpdateArticleRequest } from '../../types/article';
import { articleService } from '../../services';
import dynamic from 'next/dynamic';
import { Button, Input, Select, Space, message } from 'antd';

const MDEditor = dynamic(() => import('@uiw/react-md-editor'), { ssr: false });

interface EditorProps {
  article?: Article;
  isEdit?: boolean;
}

export const ArticleEditor: React.FC<EditorProps> = ({ article, isEdit }) => {
  const router = useRouter();
  const [title, setTitle] = useState(article?.title || '');
  const [content, setContent] = useState(article?.content || '');
  const [status, setStatus] = useState<'draft' | 'published'>(article?.status || 'draft');
  const [tags, setTags] = useState<string[]>(article?.tags.map(t => t.name) || []);
  const [tagInput, setTagInput] = useState('');
  const [loading, setLoading] = useState(false);

  useEffect(() => {
    if (article) {
      setTitle(article.title);
      setContent(article.content);
      setStatus(article.status);
      setTags(article.tags.map(t => t.name));
    }
  }, [article]);

  const handleAddTag = () => {
    if (tagInput && !tags.includes(tagInput)) {
      setTags([...tags, tagInput]);
      setTagInput('');
    }
  };

  const handleRemoveTag = (tag: string) => {
    setTags(tags.filter(t => t !== tag));
  };

  const handleSubmit = async () => {
    try {
      setLoading(true);
      const data: CreateArticleRequest = {
        title,
        content,
        status,
        tags,
      };

      if (isEdit && article) {
        const updateData: UpdateArticleRequest = {
          ...data,
          id: article.id,
        };
        await articleService.update(updateData);
        message.success('文章更新成功');
      } else {
        await articleService.create(data);
        message.success('文章创建成功');
      }
      router.push('/articles');
    } catch (error) {
      message.error('操作失败：' + (error as Error).message);
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="max-w-4xl mx-auto p-4">
      <Space direction="vertical" size="large" style={{ width: '100%' }}>
        <Input
          placeholder="文章标题"
          value={title}
          onChange={e => setTitle(e.target.value)}
          size="large"
        />

        <div data-color-mode="light">
          <MDEditor
            value={content}
            onChange={value => setContent(value || '')}
            height={400}
          />
        </div>

        <Space>
          <Select
            value={status}
            onChange={value => setStatus(value)}
            style={{ width: 120 }}
          >
            <Select.Option value="draft">草稿</Select.Option>
            <Select.Option value="published">发布</Select.Option>
          </Select>

          <Input
            placeholder="添加标签"
            value={tagInput}
            onChange={e => setTagInput(e.target.value)}
            onPressEnter={handleAddTag}
            style={{ width: 120 }}
          />
          <Button onClick={handleAddTag}>添加标签</Button>
        </Space>

        <div>
          {tags.map(tag => (
            <span
              key={tag}
              className="inline-block bg-gray-200 rounded-full px-3 py-1 text-sm font-semibold text-gray-700 mr-2 mb-2"
            >
              {tag}
              <button
                onClick={() => handleRemoveTag(tag)}
                className="ml-2 text-red-500"
              >
                ×
              </button>
            </span>
          ))}
        </div>

        <div className="flex justify-end">
          <Space>
            <Button onClick={() => router.back()}>取消</Button>
            <Button type="primary" onClick={handleSubmit} loading={loading}>
              {isEdit ? '更新' : '创建'}
            </Button>
          </Space>
        </div>
      </Space>
    </div>
  );
};