import React, { useState, useEffect } from "react";

import { eventEmitter } from "../../../services/socket-service";

import "./ChatList.css";

function ChatList(props) {
  const [selectedUser, setSelectedUser] = useState(null);
  const [chatList, setChatList] = useState([]);

  const chatListSubscription = (socketPayload) => {
    let newChatList = chatList;

    if (socketPayload.type === "new-user-joined") {
      const incomingChatList = socketPayload.chatlist;
      if (incomingChatList) {
        newChatList = newChatList.filter(
          (obj) => obj.userID !== incomingChatList.userID
        );
      }

      /* Adding new online user into chat list array */
      newChatList = [...newChatList, ...[incomingChatList]];
    } else if (socketPayload.type === "user-disconnected") {
      const outGoingUser = socketPayload.chatlist;
      const loggedOutUserIndex = newChatList.findIndex(
        (obj) => obj.userID === outGoingUser.userID
      );
      if (loggedOutUserIndex >= 0) {
        newChatList[loggedOutUserIndex].online = "N";
      }
    } else {
      newChatList = socketPayload.chatlist;
    }

    // slice is used to create aa new instance of an array
    if(newChatList !== null) {
      setChatList(newChatList.slice());
    } else {
      setChatList([]);
    }
    
  };

  useEffect(() => {
    eventEmitter.on("chatlist-response", chatListSubscription);
    return () => {
      eventEmitter.removeListener("chatlist-response", chatListSubscription);
    };
  });

  const updateSelectedUser = (user) => {
    if (user) {
      setSelectedUser(user);
      props.setSelectedUser(user);
    }
  };

  if (chatList && chatList.length === 0) {
    return (
      <div className="alert">
        {chatList.length === 0
          ? "Loading your chat list."
          : "No User Available to chat."}
      </div>
    );
  }

  return (
    <div className="app__chatlist-container">
      <div className="user__chat-list">
        {chatList.map((user, index) => (
          <div
            key={index}
            className={`user-name ${
              selectedUser !== null && selectedUser.userID === user.userID
                ? "selected-username"
                : ""
            }`}
            onClick={() => updateSelectedUser(user)}
          >
            {user.username}
            <span className={user.online === "Y" ? "online" : "offline"}></span>
          </div>
        ))}
      </div>
    </div>
  );
}

export default ChatList;
