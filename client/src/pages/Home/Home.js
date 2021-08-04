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
import Conversation from "./Conversation/Conversation";

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
    <div className="App">
      <header className="app-header">
        <nav className="navbar navbar-expand-md">
          <h4>Hello {getUsername(userDetail)}</h4>
        </nav>
        <ul className="nav justify-content-end">
          <li className="nav-item">
            <button
              class="btn btn-outline-primary"
              onClick={() => {
                logoutUser(props, userDetail);
              }}
            >
              Logout
            </button>
          </li>
        </ul>
      </header>

      <main role="main" className="container content">
        <div className="row chat-content">
          <div className="col-3 chat-list-container">
            <ChatList
              setSelectedUser={(user) => {
                setSelectedUser(user);
              }}
              userDetail={userDetail}
            />
          </div>
          <div className="col-8 message-container">
            <Conversation userDetail={userDetail} selectedUser={selectedUser} />
          </div>
        </div>
      </main>
    </div>
  );
}

export default withRouter(Home);
