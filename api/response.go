package api

import (
	"encoding/json"
	"fmt"
)

// Error is a custom error, defined to facilitate JSON marshalling
type Error struct {
	e string
}

// MarshalJSON converts error to json, which is not otherwise possible in golang
func (v Error) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.e)
}

func (v Error) Error() string {
	return v.e
}

// NewError creates an error
func NewError(s string) Error {
	return Error{s}
}

// NewErrors returns a list of errors from a list of string
func NewErrors(s ...string) []Error {
	errors := make([]Error, 0)
	for _, e := range s {
		errors = append(errors, Error{e})
	}
	return errors
}

// Response is the standard API response interface
type Response interface {
	Success() bool
	GetErrors() []Error
	String() string
	AddError(err error)
	Equals(r interface{}) bool
}

// StdResponse implements Response
type StdResponse struct {
	Successful bool    `json:"successful"`
	Errors     []Error `json:"errors"`
}

// Equals checks if the response is same to any other response
func (r *StdResponse) Equals(other interface{}) bool {
	if otherResp, ok := other.(Response); ok {
		return r.Successful == otherResp.Success()
	}
	return false
}

// AddError adds an error to response
func (r *StdResponse) AddError(err error) {
	errors := append(r.Errors, Error{err.Error()})
	r.Errors = errors
}

// AddNewError adds a string error to response
func (r *StdResponse) AddNewError(err string) {
	errors := append(r.Errors, Error{err})
	r.Errors = errors
}

// Success returns success status of response
func (r StdResponse) Success() bool {
	return r.Successful
}

// GetErrors returns all errror of the response
func (r StdResponse) GetErrors() []Error {
	return r.Errors
}

func (r StdResponse) String() string {
	errs := "["
	for _, e := range r.Errors {
		errs = errs + fmt.Sprintf("\"%v\", ", e.e)
	}
	errs = errs + "]"
	return fmt.Sprintf(`{"successful":%v, "errors":%v}`, r.Success(), errs)
}

// NewStdResponse creates a new StdResponse object
func NewStdResponse() *StdResponse {
	return &StdResponse{Successful: true}
}

// NewErrorResponse creates an error response object, with the desired error
func NewErrorResponse(err error) *StdResponse {
	return &StdResponse{
		Successful: false,
		Errors:     []Error{Error{err.Error()}},
	}
}
