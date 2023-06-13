package errors

import "fmt"

// ResponseError struct
type ResponseError struct {
	Field     string
	Message   string
	MainError error
}

func (e *ResponseError) Error() string {
	return fmt.Sprintf("Field: %s, Message: %s, MainError: %s", e.Field, e.Message, e.MainError.Error())
}
