import React, { useState, useEffect } from "react";

import { withRouter } from "react-router";

import { getItemFormSS, removeItemInSS } from "../../services/storage-service";

import "./Home.css";

const useFetch = (props) => {
  const [internalErr, setInternalErr] = useState(null);
  const userDetail = getItemFormSS("userDetail");

  return [userDetail, internalErr];
};

const getUsernameInitial = (userDetail) => {
  if (userDetail && userDetail.username) {
    return userDetail.username;
  }
  return "_";
};

const getUsername = (userDetail) => {
  if (userDetail && userDetail.username) {
    return userDetail.username;
  }

  return "___";
};

const logoutUser = (props, userDetail) => {
  if (userDetail.userID) {
    removeItemInSS("userDetail");
    // websocket event
    props.history.push(`/`);
  }
};

function Home(props) {
  const [userDetail, internalErr] = useFetch(props);
  if (internalErr !== null) {
    return <h1>{internalErr}</h1>;
  }

  return (
    <div className="app__home-container">
      <header className="app__header-container">
        <nav className="app__header-user">
          <div className="user-detail">
            <h4>{getUsername(userDetail)}</h4>
          </div>
        </nav>
        <button
          className="logout"
          onClick={() => {
            logoutUser(props, userDetail);
          }}
        >
          Logout
        </button>
      </header>

      <div className="app__content-container">
        <div className="app__hone-chatlist"></div>

        <div className="app__home-message"></div>
      </div>
    </div>
  );
}

export default withRouter(Home);
