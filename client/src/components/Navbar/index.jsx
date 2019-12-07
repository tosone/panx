import React from 'react';
import { withRouter } from 'react-router-dom'
import { Layout, Menu, Icon } from 'antd';

class Navbar extends React.Component {
  constructor(props) {
    super(props);

    this.handleClick = this.handleClick.bind(this);
  };

  state = {
    current: 'mail',
  };

  handleClick = e => {
    this.props.history.push(e.key);
  };

  render() {
    return (
      <Layout.Header>
        <Menu
          onClick={this.handleClick}
          selectedKeys={[this.state.current]}
          mode="horizontal"
          theme="dark"
          defaultSelectedKeys={['home']}
          style={{ lineHeight: '64px' }}>
          <Menu.Item key="home">
            <Icon type="home" />首页
          </Menu.Item>
          <Menu.Item key="explor">
            <Icon type="star" />探索
          </Menu.Item>
          <Menu.Item key="ranking">
            <Icon type="line-chart" />排行榜
          </Menu.Item>
          <Menu.Item key="signin" style={{ float: 'right' }}>
            登录
          </Menu.Item>
          <Menu.Item key="signup" style={{ float: 'right' }}>
            注册
          </Menu.Item>
        </Menu>
      </Layout.Header>
    );
  }
}

export default withRouter(Navbar);
