import { NextPage } from 'next';
import { useRouter } from 'next/router';
import { useEffect, useState } from 'react';
import { ArticleEditor } from '../../../components/article/Editor';
import { Article } from '../../../types/article';
import { articleService } from '../../../services';
import { Layout, Spin, message } from 'antd';

const { Content } = Layout;

const EditArticlePage: NextPage = () => {
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
        <ArticleEditor article={article} isEdit />
      </Content>
    </Layout>
  );
};

export default EditArticlePage;