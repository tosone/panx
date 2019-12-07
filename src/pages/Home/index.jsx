import React from 'react';
import { Layout } from 'antd';

import Banner from '../../components/Navbar';
import UserInfoCard from '../../components/UserInfoCard';

import './index.css';

class Home extends React.Component {
  render() {
    return (
      <Layout>
        <Banner />
        <div className="main">
          <div className="content">content</div>
          <div className="content-sider">
            <UserInfoCard />
          </div>
        </div>

        <Layout.Footer style={{ textAlign: 'center' }}>
          Code Share ©2019 Created by
          <a
            href="https://github.com/tosone"
            target="_blank"
            rel="noopener noreferrer">
              &nbsp;Tosone
          </a>
        </Layout.Footer>
      </Layout>
    );
  }
}

export default Home;
