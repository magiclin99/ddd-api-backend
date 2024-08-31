package dto

type CreateTaskRequest struct {
	Name string `json:"name" binding:"required"`
}

func NewApiError(code int, msg string) *ApiResponse {
	return &ApiResponse{
		Code:         code,
		ErrorMessage: msg,
	}
}

func NewApiOK(data any) *ApiResponse {
	return &ApiResponse{
		Code: 0,
		Data: data,
	}

}

type ApiResponse struct {
	Code         int    `json:"code"` // 0 - success, otherwise error
	ErrorMessage string `json:"errorMsg,omitempty"`
	Data         any    `json:"data,omitempty"`
}
