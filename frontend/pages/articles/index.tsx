import { NextPage } from 'next';
import { ArticleList } from '../../components/article/List';
import { Layout } from 'antd';

const { Content } = Layout;

const ArticlesPage: NextPage = () => {
  return (
    <Layout>
      <Content>
        {/* <ArticleList /> */}
      </Content>
    </Layout>
  );
};

export default ArticlesPage;