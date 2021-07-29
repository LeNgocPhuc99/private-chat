const events = require("events");

const CHAT_SERVER_ENDPOINT = "127.0.0.1:8080";
let webSocketConnection = null;

export const eventEmitter = new events.EventEmitter();

export function connectToWebSocket(userID) {
  if (userID === "" || userID === null || userID === undefined) {
    return {
      message: "You need user ID to connect to chat server",
      webSocketConnection: null,
    };
  }

  if (window["WebSocket"]) {
    webSocketConnection = new WebSocket(
      "ws://" + CHAT_SERVER_ENDPOINT + "/ws/" + userID
    );
    return {
      message: "You are connected to Chat Server",
      webSocketConnection: webSocketConnection,
    };
  } else {
    return {
      message: "Your Browser don't support Web Socket",
      webSocketConnection: null,
    };
  }
}

export function listenToWebSocketEvents() {
  if (webSocketConnection === null) {
    return;
  }

  webSocketConnection.onclose = (event) => {
    eventEmitter.emit("disconnect", event);
  };

  webSocketConnection.onmessage = (event) => {
    try {
      const socketPayload = JSON.parse(event.data);
      switch (socketPayload.eventName) {
        case "chatlist-response":
          if (!socketPayload.eventPayload) {
            return;
          }
          eventEmitter.emit("chatlist-response", socketPayload.eventPayload);
          break;

        case "disconnect":
          if (!socketPayload.eventPayload) {
            return;
          }
          eventEmitter.emit("chatlist-response", socketPayload.eventPayload);
          break;
        default:
          console.log("default case");
      }
    } catch (error) {
      console.log(error);
      console.warn("Something went wrong while decoding the Message Payload");
    }
  };
}

export function sendMessage(messagePayload) {
  if (webSocketConnection === null) {
    return;
  }

  webSocketConnection.send(
    JSON.stringify({
      eventName: "message",
      eventPayload: messagePayload,
    })
  );
}

export function emitLogoutEvent() {
  if (webSocketConnection === null) {
    return;
  }

  webSocketConnection.close();
}
