package aperr

type Error struct {
	Code    string
	Message string
}

func (it *Error) Error() string {
	return it.Message
}

var unexpectedServerError = &Error{
	Code:    "C-00000",
	Message: "",
}

// -- error constructors --

func UnexpectedServerError() *Error {
	return unexpectedServerError
}

func InvalidRequest(msg string) *Error {
	return &Error{
		Code:    "C-00001",
		Message: msg,
	}
}

var TaskNotFound = &Error{
	Code:    "TASK-00001",
	Message: "task not found",
}
