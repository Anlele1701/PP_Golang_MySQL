package common

import (
	"net/http"
	"time"
)

type responseBody struct {
	Paging interface{} `json:"paging,omitempty"`
	Filter interface{} `json:"filter,omitempty"`
	Data   interface{} `json:"data"`
}

type successResponse struct {
	TimeStamp    time.Time    `json:"timestamp"`
	ResponseCode int          `json:"responseCode"`
	Message      string       `json:"message"`
	ResponseBody responseBody `json:"responseBody"`
}

func NewSuccessResponse(data, paging, filter interface{}) *successResponse {
	return &successResponse{
		TimeStamp:    time.Now(),
		ResponseCode: http.StatusOK,
		Message:      "Success",
		ResponseBody: responseBody{
			Paging: paging,
			Filter: filter,
			Data:   data,
		},
	}
}
func SimepleSuccessResponse(data interface{}) *successResponse {
	return NewSuccessResponse(data, nil, nil)
}
