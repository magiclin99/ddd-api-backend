package aperr

type Error struct {
	Code    int
	Message string
}

func (it *Error) Error() string {
	return it.Message
}

var unexpectedServerError = &Error{
	Code:    1000000, // convention =  component code(100) + error code(0000)
	Message: "",
}

// -- error constructors --

func UnexpectedServerError() *Error {
	return unexpectedServerError
}

func InvalidRequest(msg string) *Error {
	return &Error{
		Code:    2000001,
		Message: msg,
	}
}
