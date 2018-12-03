import React from 'react';

import './index.css';

class UserInfoCard extends React.Component {
  render() {
    return (
      <div className="UserInfoCard">
        <div className="UserInfoCard-bg" />
        <img
          alt="face"
          src="/panx/img/face.png"
          className="UserInfoCard-face-img"
        />
        <div className="UserInfoCard-username">Tosone</div>
        <div className="UserInfoCard-info">
          <div className="UserInfoCard-info-left">
            <div className="UserInfoCard-info-number">217</div>
            <div className="UserInfoCard-info-intro">Snippets</div>
          </div>
          <div className="UserInfoCard-info-mid">
            <div className="UserInfoCard-info-number">323</div>
            <div className="UserInfoCard-info-intro">Following</div>
          </div>
          <div className="UserInfoCard-info-right">
            <div className="UserInfoCard-info-number">479</div>
            <div className="UserInfoCard-info-intro">Followers</div>
          </div>
        </div>
      </div>
    );
  }
}

export default UserInfoCard;
