import { NextPage } from 'next';
import { useRouter } from 'next/router';
import { useEffect, useState } from 'react';
import { ArticlePreview } from '../../components/article/Preview';
import { Article } from '../../types/article';
import { articleService } from '../../services';
import { Layout, Spin, Button, Space, message } from 'antd';

const { Content } = Layout;

const ArticleDetailPage: NextPage = () => {
  const router = useRouter();
  const { id } = router.query;
  const [article, setArticle] = useState<Article | null>(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchArticle = async () => {
      if (id) {
        try {
          const data = await articleService.getById(Number(id));
          setArticle(data);
        } catch (error) {
          message.error('加载文章失败：' + (error as Error).message);
          router.push('/articles');
        } finally {
          setLoading(false);
        }
      }
    };

    fetchArticle();
  }, [id, router]);

  const handleDelete = async () => {
    if (!article) return;

    try {
      await articleService.delete(article.id);
      message.success('删除成功');
      router.push('/articles');
    } catch (error) {
      message.error('删除失败：' + (error as Error).message);
    }
  };

  if (loading) {
    return (
      <Layout>
        <Content className="flex justify-center items-center min-h-screen">
          <Spin size="large" />
        </Content>
      </Layout>
    );
  }

  if (!article) {
    return null;
  }

  return (
    <Layout>
      <Content>
        <div className="max-w-4xl mx-auto">
          <div className="mb-4 flex justify-end">
            <Space>
              <Button onClick={() => router.push('/articles')}>返回列表</Button>
              <Button type="primary" onClick={() => router.push(`/articles/edit/${article.id}`)}>
                编辑文章
              </Button>
              <Button danger onClick={handleDelete}>
                删除文章
              </Button>
            </Space>
          </div>
          <ArticlePreview article={article} />
        </div>
      </Content>
    </Layout>
  );
};

export default ArticleDetailPage;