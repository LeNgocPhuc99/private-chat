import React, { useState, useEffect } from "react";

import { withRouter } from "react-router";

import { userLoginCheckRequest } from "../../services/api-service";
import { getItemFormSS, removeItemInSS } from "../../services/storage-service";
import {
  connectToWebSocket,
  listenToWebSocketEvents,
  emitLogoutEvent,
} from "../../services/socket-service";
import ChatList from "./ChatList/ChatList";

import "./Home.css";

const useFetch = (props) => {
  const [internalErr, setInternalErr] = useState(null);
  const userDetail = getItemFormSS("userDetail");

  // after render home page
  useEffect(() => {
    (async () => {
      // check userDetail Info
      if (userDetail === null || userDetail === "") {
        props.history.push(`/`);
      } else {
        // check user's session
        const isUserLogged = await userLoginCheckRequest(userDetail.userID);
        if (!isUserLogged.response) {
          props.history.push(`/`);
        } else {
          // connect to web socket
          console.log("Connect to server");
          const connection = connectToWebSocket(userDetail.userID);
          if (connection.webSocketConnection === null) {
            setInternalErr(connection.message);
          } else {
            console.log("Listing event");
            listenToWebSocketEvents();
          }
        }
      }
    })();
  }, [props, userDetail]);

  return [userDetail, internalErr];
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
    emitLogoutEvent();
    props.history.push(`/`);
  }
};

function Home(props) {
  const [userDetail, internalErr] = useFetch(props);
  const [selectedUser, setSelectedUser] = useState(null);
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
        <div className="app__hone-chatlist">
          <ChatList
            setSelectedUser={(user) => {
              setSelectedUser(user);
            }}
            userDetail={userDetail}
          />
        </div>

        <div className="app__home-message"></div>
      </div>
    </div>
  );
}

export default withRouter(Home);
