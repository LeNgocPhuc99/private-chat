package handlers

type Hub struct {
	clients    map[*Client]bool
	register   chan *Client
	unregister chan *Client
}

func NewHub() *Hub {
	return &Hub{
		clients:    make(map[*Client]bool),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

func (hub *Hub) Run() {
	for {
		select {
		case client := <-hub.register:
			HandleUserJoinEvent(hub, client)
		case client := <-hub.unregister:
			HandleUserDisconnectEvent(hub, client)
		}
	}
}
