// Package errors provides custom made errors for aggart
package errors

import (
	"server/utils"
)

// Error codes
const (
	argumentMismatch errorCode = iota
	invalidType
	objectNotFound
	objectAlreadyExists
)

var (
	errorCodesToStringMap map[errorCode]string

	// ErrorObjectNotFound when you don't find an entity from dao
	ErrorObjectNotFound = AggError{Code: objectNotFound}
	// ErrorObjectAlreadyExists is when you try to persist in dao, but it already ezists
	ErrorObjectAlreadyExists = AggError{Code: objectAlreadyExists}

	// ErrorArgumentMismatch means expected argument was not found
	ErrorArgumentMismatch = AggError{Code: argumentMismatch}
	// ErrorInvalidType is
	ErrorInvalidType = AggError{Code: invalidType}
)

func init() {
	errorCodesToStringMap = map[errorCode]string{
		argumentMismatch:    "ArgumentMismatch",
		invalidType:         "InvalidType",
		objectNotFound:      "ObjectNotFound",
		objectAlreadyExists: "ObjectAlreadyExists",
	}
}

type errorCode int

// AggError is customized error
type AggError struct {
	Code    errorCode
	Message string
}

// Error implements error interface
func (e AggError) Error() string {
	s := "Code: [" + e.StringCode() + "]"
	if !utils.IsBlank(e.Message) {
		s = s + ", Message: [" + e.Message + "]"
	}
	return s
}

// StringCode returns string code
func (e AggError) StringCode() string {
	return errorCodesToStringMap[e.Code]
}
