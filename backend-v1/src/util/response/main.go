package response

import "fmt"

type CustomResponse struct {
	Code string `json:"code"`
	Payload *interface{} `json:"payload,omitempty"`
	Message *string `json:"message,omitempty"`
}

func Success(code int, payload interface{}) CustomResponse {
	return CustomResponse{
		fmt.Sprintf("%d", code),
		&payload,
		nil,
	}
}

func Error(code int, payload *interface{}, message *string) CustomResponse {
	return CustomResponse{
		fmt.Sprintf("%d", code),
		payload,
		message,
	}
}


