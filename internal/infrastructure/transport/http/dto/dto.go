package dto

func NewApiError(code string, msg string) *ApiError {
	return &ApiError{
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
	Data any `json:"data,omitempty"`
}

type ApiError struct {
	Code         string `json:"code"`
	ErrorMessage string `json:"errorMsg"`
}

type CreateTaskRequest struct {
	Name string `json:"name" binding:"min=1,max=100"`
}
