package handlers

import (
	"encoding/json"
	"net/http"
)

type APIResponse struct {
	Code     int         `json:"code"`
	Status   string      `json:"status"`
	Message  string      `json:"message"`
	Response interface{} `json:"response"`
}

func Response(rw http.ResponseWriter, r *http.Request, apiResponse APIResponse) {
	var (
		responseMessage, responseStatusText string
		responseHTTPCode                    int
	)

	if apiResponse.Code == 0 {
		responseHTTPCode = http.StatusOK
	} else {
		responseHTTPCode = apiResponse.Code
	}

	if apiResponse.Status != "" {
		responseStatusText = apiResponse.Status
	} else {
		responseStatusText = http.StatusText(http.StatusOK)
	}

	if apiResponse.Message != "" {
		responseMessage = apiResponse.Message
	} else {
		responseMessage = SuccessfulResponse
	}

	httpResponse := APIResponse{
		Code:     responseHTTPCode,
		Status:   responseStatusText,
		Message:  responseMessage,
		Response: apiResponse.Response,
	}

	responsePayload, err := json.Marshal(httpResponse)
	if err != nil {
		panic(err)
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(httpResponse.Code)
	rw.Write(responsePayload)
}
