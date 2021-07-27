

const CHAT_SERVER_ENDPOINT = "127.0.0.1:8080";
let webSocketConnection = null;

export function connectToWebSocket(userID) {
    if(userID === "" || userID === null || userID === undefined) {
        return {
            message: "You need user ID to connect to chat server",
            webSocketConnection: null
        }
    } 

    if(window["WebSocket"]) {
        webSocketConnection = new WebSocket("ws://" + CHAT_SERVER_ENDPOINT + "/ws/" + userID);
        return {
            message: "You are connected to Chat Server",
            webSocketConnection: webSocketConnection
        }
    } else {
        return {
            message: "Your Browser don't support Web Socket",
            webSocketConnection: null
        }
    }
}