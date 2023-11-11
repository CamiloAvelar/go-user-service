package errors

import "fmt"

type ApplicationError struct {
	Type       string
	Message    string
	StatusCode int
}

func (e *ApplicationError) Error() string {
	return fmt.Sprintf(e.Message)
}
