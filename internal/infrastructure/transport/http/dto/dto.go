package dto

func NewApiError(code string, msg string) *ApiResponse {
	return &ApiResponse{
		Code:         code,
		ErrorMessage: msg,
	}
}

func NewApiOK(data any) *ApiResponse {
	return &ApiResponse{
		Data: data,
	}

}

type ApiResponse struct {
	Code         string `json:"code"` // empty for success, otherwise error
	ErrorMessage string `json:"errorMsg,omitempty"`
	Data         any    `json:"data,omitempty"`
}

type CreateTaskRequest struct {
	Name string `json:"name" binding:"min=1,max=100"`
}
