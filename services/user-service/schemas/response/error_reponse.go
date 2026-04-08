package response

type ErrorResponse struct {
	Message string `json:"message"`
}

func ToErrorResponse(err error) *ErrorResponse {
	return &ErrorResponse{Message: err.Error()}
}
