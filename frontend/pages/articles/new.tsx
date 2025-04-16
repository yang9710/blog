import { NextPage } from 'next';
import { ArticleEditor } from '../../components/article/Editor';
import { Layout } from 'antd';

const { Content } = Layout;

const NewArticlePage: NextPage = () => {
  return (
    <Layout>
      <Content>
        <ArticleEditor />
      </Content>
    </Layout>
  );
};

export default NewArticlePage;