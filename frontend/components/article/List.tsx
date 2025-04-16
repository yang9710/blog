import React, { useState, useEffect } from 'react';
import { useRouter } from 'next/router';
import { Article } from '../../types/article';
import { articleService } from '../../services';
import { Table, Space, Button, Tag, Modal, message } from 'antd';
import type { ColumnsType } from 'antd/es/table';
import { format } from 'date-fns';

export const ArticleList: React.FC = () => {
  const router = useRouter();
  const [articles, setArticles] = useState<Article[]>([]);
  const [total, setTotal] = useState(0);
  const [loading, setLoading] = useState(false);
  const [currentPage, setCurrentPage] = useState(1);
  const pageSize = 10;

  const loadArticles = async (page: number) => {
    try {
      setLoading(true);
      const response = await articleService.list({
        page,
        page_size: pageSize,
      });
      setArticles(response.items);
      setTotal(response.total);
    } catch (error) {
      message.error('加载文章列表失败：' + (error as Error).message);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    loadArticles(currentPage);
  }, [currentPage]);

  const handleDelete = (id: number) => {
    Modal.confirm({
      title: '确认删除',
      content: '确定要删除这篇文章吗？',
      onOk: async () => {
        try {
          await articleService.delete(id);
          message.success('删除成功');
          loadArticles(currentPage);
        } catch (error) {
          message.error('删除失败：' + (error as Error).message);
        }
      },
    });
  };

  const columns: ColumnsType<Article> = [
    {
      title: '标题',
      dataIndex: 'title',
      key: 'title',
      render: (text, record) => (
        <a onClick={() => router.push(`/articles/${record.id}`)}>{text}</a>
      ),
    },
    {
      title: '作者',
      dataIndex: ['author', 'username'],
      key: 'author',
    },
    {
      title: '状态',
      dataIndex: 'status',
      key: 'status',
      render: (status: string) => (
        <Tag color={status === 'published' ? 'green' : 'orange'}>
          {status === 'published' ? '已发布' : '草稿'}
        </Tag>
      ),
    },
    {
      title: '标签',
      key: 'tags',
      dataIndex: 'tags',
      render: (tags: { name: string }[]) => (
        <>
          {tags.map(tag => (
            <Tag key={tag.name}>{tag.name}</Tag>
          ))}
        </>
      ),
    },
    {
      title: '创建时间',
      dataIndex: 'created_at',
      key: 'created_at',
      render: (date: string) => format(new Date(date), 'yyyy-MM-dd HH:mm:ss'),
    },
    {
      title: '操作',
      key: 'action',
      render: (_, record) => (
        <Space size="middle">
          <Button type="link" onClick={() => router.push(`/articles/edit/${record.id}`)}>
            编辑
          </Button>
          <Button type="link" danger onClick={() => handleDelete(record.id)}>
            删除
          </Button>
        </Space>
      ),
    },
  ];

  return (
    <div className="p-4">
      <div className="mb-4 flex justify-between items-center">
        <h1 className="text-2xl font-bold">文章列表</h1>
        <Button type="primary" onClick={() => router.push('/articles/new')}>
          新建文章
        </Button>
      </div>
      <Table
        columns={columns}
        dataSource={articles}
        rowKey="id"
        loading={loading}
        pagination={{
          total,
          current: currentPage,
          pageSize,
          onChange: (page) => setCurrentPage(page),
        }}
      />
    </div>
  );
};