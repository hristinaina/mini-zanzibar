package errors

import "fmt"

type CustomError struct {
	Code    int
	Message string
}

// Error implements the error interface for CustomError.
func (e CustomError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}
