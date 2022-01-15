package response

type SuccessResponse struct {
	//A
}

type ErrorResponse struct {
	//B
}

type ErrorResponses struct {
	Success bool        `json:"success"`
	Message interface{} `json:"message"`
}
