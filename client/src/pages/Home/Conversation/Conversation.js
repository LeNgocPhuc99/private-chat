import React, { useState, useEffect, useRef } from "react";
import { Form } from "react-bootstrap";
import {
  eventEmitter,
  sendWebSocketMessage,
} from "../../../services/socket-service";
import { getConversationBetweenUsers } from "../../../services/api-service";

import "./Conversation.css";

const alignMessage = (userDetail, toUserID) => {
  const { userID } = userDetail;
  return userID !== toUserID;
};

const scrollMessageContainer = (messageContainer) => {
  // message UI
  if (messageContainer.current !== null) {
    try {
      setTimeout(() => {
        messageContainer.current.scrollTop =
          messageContainer.current.scrollHeight;
      }, 100);
    } catch (error) {
      console.warn(error);
    }
  }
};

const getMessageUI = (messageContainer, userDetail, conversations) => {
  return (
    <>
      <ul ref={messageContainer} className="message-thread">
        {conversations.map((conversation, index) => (
          <li
            className={`${
              alignMessage(userDetail, conversation.toUserID)
                ? "align-right"
                : ""
            }`}
            key={index}
          >
            {conversation.message}
          </li>
        ))}
      </ul>
    </>
  );
};

const getInitiateConversationUI = (userDetail) => {
  if (userDetail !== null) {
    return (
      <div className="message-thread start-chatting-banner">
        <p className="heading">
          You haven 't chatted with {userDetail.username} in a while,
          <span className="sub-heading"> Say Hi.</span>
        </p>
      </div>
    );
  }
};

function Conversation(props) {
  const selectedUser = props.selectedUser;
  const userDetail = props.userDetail;

  const messageContainer = useRef(null);
  const [conversations, updateConversation] = useState([]);
  const [messageLoading, updateMessageLoading] = useState(true);

  // get message conversation
  useEffect(() => {
    if (userDetail && selectedUser) {
      (async () => {
        const conversationResponse = await getConversationBetweenUsers(
          userDetail.userID,
          selectedUser.userID
        );
        updateMessageLoading(false);
        if (conversationResponse.response) {
          updateConversation(conversationResponse.response);
          // console.log("Load:", conversationResponse.response);
        } else if (conversationResponse.response === null) {
          updateConversation([]);
        }
      })();
    }
  }, [userDetail, selectedUser]);

  // receive message from selected user
  useEffect(() => {
    const newMessage = (messagePayload) => {
      if (
        selectedUser !== null &&
        selectedUser.userID === messagePayload.fromUserID
      ) {
        updateConversation([...conversations, messagePayload]);
        scrollMessageContainer(messageContainer);
      }
    };

    eventEmitter.on("message-response", newMessage);
    return () => {
      eventEmitter.removeListener("message-response", newMessage);
    };
  }, [conversations, selectedUser]);

  // send message
  const sendMessage = (event) => {
    if (event.key === "Enter") {
      const message = event.target.value;
      // check
      if (message === "" || message === undefined || message === null) {
        alert(`Message can't be empty.`);
      } else if (selectedUser === undefined) {
        alert(`Select a user to chat.`);
      } else {
        event.target.value = "";
        const messagePayload = {
          message: message.trim(),
          toUserID: selectedUser.userID,
          fromUserID: userDetail.userID,
        };
        sendWebSocketMessage(messagePayload);
        updateConversation([...conversations, messagePayload]);
        scrollMessageContainer(messageContainer);
      }
    }
  };

  if (messageLoading) {
    return (
      <div className="message-overlay">
        <h3>
          {selectedUser !== null && selectedUser.username
            ? "Loading Message"
            : "Select a User to chat"}
        </h3>
      </div>
    );
  }

  return (
    <div
      className={`message-wrapper ${messageLoading ? "visibility-hidden" : ""}`}
    >
      <div className="message-container">
        {conversations.length > 0
          ? getMessageUI(messageContainer, userDetail, conversations)
          : getInitiateConversationUI(selectedUser)}
      </div>

      <div className="message-typer">
        <Form.Control
          as="textarea"
          placeholder={`${
            selectedUser !== null ? "" : "Select a user and"
          } Type your message here`}
          onKeyPress={sendMessage}
        />
        
      </div>
    </div>
  );
}

export default Conversation;
