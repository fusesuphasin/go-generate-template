package response

type SuccessResponse struct {
}

type ErrorResponse struct {
}

type ErrorResponses struct {
	Success bool        `json:"success"`
	Message interface{} `json:"message"`
}
