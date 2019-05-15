// Copyright 2019 Pavel Petrov <itnelo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package response

import (
	"fmt"
)

// Represents type of application-layer error (domain error identifier).
type ErrorType uint16

const (
	InternalError ErrorType = iota
	BadRequestError
)

var responseErrorTypes = [...]string{
	"main.internal_error",
	"request.binder.bad_request",
}

// Implements fmt.Stringer interface.
func (etype ErrorType) String() string {
	if etype > BadRequestError {
		panic("response: undefined error type.")
	}

	return responseErrorTypes[etype]
}

type Errors []Error

// Error is a safe description of problem that prevents request processing.
// Will be provided for client and should not contain any sensitive information.
type Error struct {
	// Not an HTTP code, but a code of application-layer error.
	// Any package or component can provide its error codes.
	Code int `json:"code"`

	// Package/component path and other context for debugging.
	Type string `json:"type"`

	// Human-friendly error description.
	Description string `json:"description"`
}

func NewBadRequestError(param string) Error {
	return NewError(400, BadRequestError, fmt.Sprintf("Invalid request param '%s'.", param))
}

func NewInternalError() Error {
	return NewError(500, InternalError, "Internal error.")
}

func NewError(code int, etype ErrorType, description string) Error {
	return Error{code, etype.String(), description}
}
