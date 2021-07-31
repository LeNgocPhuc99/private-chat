package handlers

import (
	"log"

	"github.com/gorilla/websocket"
)

func CreateUserWebSocket(hub *Hub, connection *websocket.Conn, userID string) {
	client := &Client{
		hub:       hub,
		conn:      connection,
		userID:    userID,
		send:      make(chan SocketEvent),
		writeKill: make(chan bool),
	}

	go client.read()
	go client.write()

	client.hub.register <- client

}

func HandleUserJoinEvent(hub *Hub, client *Client) {
	if !checkExistConnection(hub, client) {
		log.Println("Join Event")
		hub.clients[client] = true
		log.Println("Number Connection:", len(hub.clients))
		handleSocketEvent(client, SocketEvent{
			EventName:    "join",
			EventPayload: client.userID,
		})
	}
}

func HandleUserDisconnectEvent(hub *Hub, client *Client) {
	if _, ok := hub.clients[client]; ok {
		log.Println("Disconnect Event")
		delete(hub.clients, client)
		close(client.send)
		log.Println("Number Connection:", len(hub.clients))
		handleSocketEvent(client, SocketEvent{
			EventName:    "disconnect",
			EventPayload: client.userID,
		})
	}
}

func BroadcastToAll(hub *Hub, payload SocketEvent) {
	for client := range hub.clients {
		select {
		case client.send <- payload:
		default:
			close(client.send)
			delete(hub.clients, client)
		}
	}
}

func BroadcastToAllExceptMe(hub *Hub, payload SocketEvent, myUserID string) {
	for client := range hub.clients {
		if client.userID != myUserID {
			select {
			case client.send <- payload:
			default:
				close(client.send)
				delete(hub.clients, client)
			}
		}
	}
}

func SendToSpecificClient(hub *Hub, payload SocketEvent, userID string) {
	for client := range hub.clients {
		if client.userID == userID {
			if payload.EventName == "message-response" {
				log.Println("Send to: ", userID)
			}
			select {
			case client.send <- payload:
			default:
				close(client.send)
				delete(hub.clients, client)
			}
		}
	}
}

func handleSocketEvent(client *Client, eventPayload SocketEvent) {
	// <=> connection payload
	type chatlistResponse struct {
		Type     string      `json:"type"`
		Chatlist interface{} `json:"chatlist"`
	}
	switch eventPayload.EventName {
	case "join":
		// check user's status
		userID := (eventPayload.EventPayload).(string)
		userDetail, err := GetUserByUserID(userID)
		if err != nil {
			log.Println("An invalid userID")
			return
		} else {
			if userDetail.Online == "N" {
				log.Println("A logout user tried to connect")
			} else {
				// broadcast new user online to all client
				newUserOnlinePayload := SocketEvent{
					EventName: "chatlist-response",
					EventPayload: chatlistResponse{
						Type: "new-user-joined",
						Chatlist: UserResponse{
							Username: userDetail.Username,
							UserID:   userDetail.ID,
							Online:   userDetail.Online,
						},
					},
				}
				BroadcastToAllExceptMe(client.hub, newUserOnlinePayload, userDetail.ID)

				chatList, err := GetAllOnlineUser(userDetail.ID)
				if err != nil {
					log.Println("Get all online user error")
					return
				}

				allOnlineUserPayload := SocketEvent{
					EventName: "chatlist-response",
					EventPayload: chatlistResponse{
						Type:     "my-chat-list",
						Chatlist: chatList,
					},
				}
				// send list online client to user
				SendToSpecificClient(client.hub, allOnlineUserPayload, userDetail.ID)
			}
		}

	case "disconnect":
		userID := (eventPayload.EventPayload).(string)
		userDetail, err := GetUserByUserID(userID)
		if err != nil {
			log.Println("An invalid userID")
			return
		}
		// update user status
		UpdateUserStatus(userID, "N")

		// broad cast user status to all
		userDisconnectedPayload := SocketEvent{
			EventName: "chatlist-response",
			EventPayload: chatlistResponse{
				Type: "user-disconnected",
				Chatlist: UserResponse{
					Username: userDetail.Username,
					UserID:   userDetail.ID,
					Online:   "N",
				},
			},
		}
		BroadcastToAll(client.hub, userDisconnectedPayload)

	case "message":

		// creat message payload
		message := (eventPayload.EventPayload.(map[string]interface{})["message"]).(string)
		fromUserID := (eventPayload.EventPayload.(map[string]interface{})["fromUserID"]).(string)
		toUSerID := (eventPayload.EventPayload.(map[string]interface{})["toUserID"]).(string)
		log.Println("Receive message from: ", fromUserID)
		if message != "" && fromUserID != "" && toUSerID != "" {
			// store message
			messagePayload := MessagePayload{
				Message:    message,
				ToUserID:   toUSerID,
				FromUserID: fromUserID,
			}
			StoreMessage(messagePayload)

			// send message to toUser
			socketPayload := SocketEvent{
				EventName:    "message-response",
				EventPayload: messagePayload,
			}
			SendToSpecificClient(client.hub, socketPayload, toUSerID)
		}

	}
}

func checkExistConnection(hub *Hub, client *Client) bool {
	for c := range hub.clients {
		if client.userID == c.userID {
			return true
		}
	}
	return false
}
