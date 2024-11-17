package errors

type ErrorDetails struct {
	ErrorDetails Error `json:"Error"`
}

type Error struct {
	Code    string `json:"Code"`
	Message string `json:"Message"`
}

func (e ErrorDetails) Error() string {
	return "CODE : " + e.ErrorDetails.Code + " Message : " + e.ErrorDetails.Message
}

func (e Error) Error() ErrorDetails {
	return ErrorDetails{ErrorDetails: e}
}

var (
	INTERNAL_ERROR   = &Error{Code: "INTERNAL_ERROR", Message: "Internal Server Error"}
	ID_REQUIRED      = &Error{Code: "ID_REQUIRED", Message: "ID is Required"}
	MESSAGE_REQUIRED = &Error{Code: "MESSAGE_REQUIRED", Message: "Message is Required"}
	USER_NOT_FOUND = &Error{Code: "USER_NOT_FOUND", Message: "User not Found"}

)
