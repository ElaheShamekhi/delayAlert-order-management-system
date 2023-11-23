package serr

import (
	"fmt"
	"gorm.io/gorm"
	"net/http"
)

type ErrorCode string

const (
	ErrInternal ErrorCode = "INTERNAL"
)

type ServiceError struct {
	Method    string
	Cause     error
	Message   string
	ErrorCode ErrorCode
	Code      int
}

func (e ServiceError) Error() string {
	return fmt.Sprintf(
		"%s (%d) - %s: %s",
		e.Method, e.Code, e.Message, e.Cause,
	)
}

func ValidationErr(method, message string, code ErrorCode) error {
	return &ServiceError{
		Method:    method,
		Message:   message,
		Code:      http.StatusBadRequest,
		ErrorCode: code,
	}
}

func ServiceErr(method, message string, cause error, code int) error {
	return &ServiceError{
		Method:  method,
		Cause:   cause,
		Message: message,
		Code:    code,
	}
}

func DBError(method, repo string, cause error) error {
	err := &ServiceError{
		Method: fmt.Sprintf("%s.%s", repo, method),
		Cause:  cause,
	}
	switch cause {
	case gorm.ErrRecordNotFound:
		err.Code = http.StatusNotFound
		err.Message = fmt.Sprintf("%s not found", repo)
	default:
		err.Code = http.StatusInternalServerError
		err.Message = fmt.Sprintf("could not perform action on %s", repo)
	}
	return err
}

func IsDBNotFound(err error) bool {
	if err == nil {
		return false
	}
	_err := err
	switch err.(type) {
	case *ServiceError:
		_err = err.(*ServiceError).Cause
	}
	switch _err {
	case gorm.ErrRecordNotFound:
		return true
	}
	return false
}
