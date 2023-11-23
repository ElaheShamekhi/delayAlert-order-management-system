package validation

import (
	"context"
	"fmt"
	"regexp"
	"time"
)

type ErrCode string

const (
	ErrInvalidDateFormat     ErrCode = "INVALID_DATE_FORMAT"
	ErrInvalidSpecialRequest ErrCode = "INVALID_SPECIAL_REQUEST"
	ErrUnknown               ErrCode = "UNKNOWN"
)

var specialCharacterRegex = regexp.MustCompile("[^a-zA-Z0-9 ]+")

type Err struct {
	Method  string
	Message string
	Code    ErrCode
	Unknown bool
}

func (e *Err) Error() string {
	return fmt.Sprintf("(%s) %s: %s", e.Method, e.Code, e.Message)
}

type Validation struct {
	Method     string
	Errors     []error
	MultiError bool
}

type Validator func(ctx context.Context) *Err

func New() *Validation {
	return &Validation{Method: "validation"}
}

func (v *Validation) Validate(validators ...Validator) *Validation {
	for _, validator := range validators {
		if err := validator(context.Background()); err != nil {
			err.Method = v.Method
			v.Errors = append(v.Errors, err)
			if !v.MultiError {
				return v
			}
		}
	}
	return v
}

func (v *Validation) SetMethod(method string) *Validation {
	v.Method = method
	return v
}

func (v *Validation) MultiErr() *Validation {
	v.MultiError = true
	return v
}

func StrDate(date string) Validator {
	return func(ctx context.Context) *Err {
		_, err := time.Parse(time.DateOnly, date)
		if err != nil {
			return &Err{Code: ErrInvalidDateFormat, Message: "Invalid date format"}
		}
		return nil
	}
}

func unknownError(err error) *Err {
	return &Err{
		Message: err.Error(),
		Code:    ErrUnknown,
		Unknown: true,
	}
}

func (v *Validation) Error() error {
	if len(v.Errors) > 0 {
		return v.Errors[0]
	}
	return nil
}
