package dto

type CreateTaskRequest struct {
	Name string `json:"name" binding:"required"`
}

func NewApiError(code string, msg string) *ApiResponse {
	return &ApiResponse{
		Code:         code,
		ErrorMessage: msg,
	}
}

func NewApiOK(data any) *ApiResponse {
	return &ApiResponse{
		Code: "",
		Data: data,
	}

}

type ApiResponse struct {
	Code         string `json:"code"` // empty for success, otherwise error
	ErrorMessage string `json:"errorMsg,omitempty"`
	Data         any    `json:"data,omitempty"`
}
