package api

import ()

// Request can be anything
type Request interface {
	String() string
	Validate() error
}
