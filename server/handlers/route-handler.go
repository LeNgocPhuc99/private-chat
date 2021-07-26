package handlers

import (
	"encoding/json"
	"net/http"
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
