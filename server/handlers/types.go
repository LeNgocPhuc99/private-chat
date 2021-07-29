package handlers

// user documents
type User struct {
	ID       string `bson:"_id,omitempty"`
	Username string
	Password string
	Online   string
	SocketID string
}

// conversation documents
type Conversation struct {
	ID         string `json:"id" bson:"_id,omitempty"`
	Message    string `json:"message"`
	ToUserID   string `json:"toUserID"`
	FromUserID string `json:"formUserID"`
}

type UserLoginRequest struct {
	Username string
	Password string
}

type UserRegistrationRequest struct {
	Username string
	Password string
}

type UserResponse struct {
	Username string `json:"username"`
	UserID   string `json:"userID"`
	Online   string `json:"online"`
}

type SocketEvent struct {
	EventName    string      `json:"eventName"`
	EventPayload interface{} `json:"eventPayload"`
}

type MessagePayload struct {
	FromUserID string `json:"fromUserID"`
	ToUserID   string `json:"toUserID"`
	Message    string `json:"message"`
}
