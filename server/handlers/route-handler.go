package handlers

import (
	"encoding/json"
	"net/http"
	"regexp"

	"github.com/gorilla/mux"
)

func Login(rw http.ResponseWriter, r *http.Request) {
	var requestPayload UserLoginRequest

	// decode request data
	decoder := json.NewDecoder(r.Body)
	decodeErr := decoder.Decode(&requestPayload)

	if decodeErr != nil {
		response := APIResponse{
			Code:     http.StatusBadRequest,
			Status:   http.StatusText(http.StatusOK),
			Message:  UsernameAndPasswordCantBeEmpty,
			Response: nil,
		}
		Response(rw, r, response)
	}

	// query database
	userResponse, loginErr := LoginQuery(requestPayload)
	if loginErr != nil {
		response := APIResponse{
			Code:     http.StatusNotFound,
			Status:   http.StatusText(http.StatusNotFound),
			Message:  loginErr.Error(),
			Response: nil,
		}
		Response(rw, r, response)
	} else {
		response := APIResponse{
			Code:     http.StatusOK,
			Status:   http.StatusText(http.StatusOK),
			Message:  UserLoginCompleted,
			Response: userResponse,
		}
		Response(rw, r, response)
	}
}

func Registration(rw http.ResponseWriter, r *http.Request) {
	var requestPayload UserRegistrationRequest

	// decode request data
	decoder := json.NewDecoder(r.Body)
	decodeErr := decoder.Decode(&requestPayload)
	if decodeErr != nil {
		response := APIResponse{
			Code:     http.StatusBadRequest,
			Status:   http.StatusText(http.StatusBadRequest),
			Message:  ServerFailedResponse,
			Response: nil,
		}
		Response(rw, r, response)
	}

	// check existence username

	// query database
	userObjectID, registrationErr := RegisterQuery(requestPayload)
	if registrationErr != nil {
		response := APIResponse{
			Code:     http.StatusInternalServerError,
			Status:   http.StatusText(http.StatusInternalServerError),
			Message:  registrationErr.Error(),
			Response: nil,
		}
		Response(rw, r, response)
	} else {
		response := APIResponse{
			Code:    http.StatusOK,
			Status:  http.StatusText(http.StatusOK),
			Message: UserRegistrationCompleted,
			Response: UserResponse{
				Username: requestPayload.Username,
				UserID:   userObjectID,
			},
		}
		Response(rw, r, response)
	}
}

func UserLoginCheck(rw http.ResponseWriter, r *http.Request) {
	var IsAlphaNumeric = regexp.MustCompile(`^[A-Za-z0-9]([A-Za-z0-9_-]*[A-Za-z0-9])?$`).MatchString
	userID := mux.Vars(r)["userID"]

	if !IsAlphaNumeric(userID) {
		response := APIResponse{
			Code:     http.StatusBadRequest,
			Status:   http.StatusText(http.StatusBadRequest),
			Message:  UserIdIsNotAvailable,
			Response: false,
		}
		Response(rw, r, response)
	} else {
		_, err := GetUserByUserID(userID)
		if err != nil {
			response := APIResponse{
				Code:     http.StatusOK,
				Status:   http.StatusText(http.StatusOK),
				Message:  YouAreNotLoggedIN,
				Response: false,
			}
			Response(rw, r, response)
		} else {
			response := APIResponse{
				Code:     http.StatusOK,
				Status:   http.StatusText(http.StatusOK),
				Message:  YouAreLoggedIN,
				Response: true,
			}
			Response(rw, r, response)
		}
	}
}

func GetAllUserAllOnline(rw http.ResponseWriter, r *http.Request) {
	userID := mux.Vars(r)["userID"]
	listOnlineUser, err := GetAllOnlineUser(userID)
	if err != nil {
		response := APIResponse{
			Code:     http.StatusInternalServerError,
			Status:   http.StatusText(http.StatusInternalServerError),
			Message:  "Internal Server Error",
			Response: nil,
		}
		Response(rw, r, response)
	}

	response := APIResponse{
		Code:     http.StatusOK,
		Status:   http.StatusText(http.StatusOK),
		Message:  "Get all online user success",
		Response: listOnlineUser,
	}

	Response(rw, r, response)
}

func GetMessages(rw http.ResponseWriter, r *http.Request) {
	fromUserID := mux.Vars(r)["fromUserID"]
	toUserID := mux.Vars(r)["toUserID"]

	conversations, err := GetConversationBetweenTwoUsers(fromUserID, toUserID)
	if err != nil {
		response := APIResponse{
			Code:     http.StatusInternalServerError,
			Status:   http.StatusText(http.StatusInternalServerError),
			Message:  "Internal Server Error",
			Response: nil,
		}
		Response(rw, r, response)
	}

	response := APIResponse{
		Code:     http.StatusOK,
		Status:   http.StatusText(http.StatusOK),
		Message:  "Get Conversaion susscess",
		Response: conversations,
	}
	Response(rw, r, response)
}
