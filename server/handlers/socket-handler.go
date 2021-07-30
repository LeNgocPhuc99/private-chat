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

	client.hub.register <- client

	go client.read()
	go client.write()

}

func HandleUserJoinEvent(hub *Hub, client *Client) {
	hub.clients[client] = true
	handleSocketEvent(client, SocketEvent{
		EventName:    "join",
		EventPayload: client.userID,
	})
}

func HandleUserDisconnectEvent(hub *Hub, client *Client) {
	if _, ok := hub.clients[client]; ok {
		delete(hub.clients, client)
		close(client.send)
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
		log.Println("Join Event")
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
		log.Println("Disconnect Event")

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
		log.Println("Message Event")
		// creat message payload
		message := (eventPayload.EventPayload.(map[string]interface{})["message"]).(string)
		fromUserID := (eventPayload.EventPayload.(map[string]interface{})["fromUserID"]).(string)
		toUSerID := (eventPayload.EventPayload.(map[string]interface{})["toUserID"]).(string)

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
