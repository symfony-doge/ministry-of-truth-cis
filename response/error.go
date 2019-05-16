// Copyright 2019 Pavel Petrov <itnelo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package response

import (
	"fmt"
)

// Represents code of application-layer error.
type ErrorCode uint64

const (
	CodeRouteNotFoundError ErrorCode = 1 + iota
	CodeMethodNotAllowedError
	CodeBadRequestError
	CodeInternalError
)

// Represents type of application-layer error (domain error identifier).
type ErrorType uint16

const (
	RouteNotFoundError ErrorType = iota
	MethodNotAllowedError
	BadRequestError
	InternalError
)

var responseErrorTypes = [...]string{
	"main.handler_not_found",
	"main.method_not_allowed",
	"request.binder.bad_request",
	"main.internal_error",
}

// Implements fmt.Stringer interface.
func (etype ErrorType) String() string {
	if etype > InternalError {
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
	Code ErrorCode `json:"code"`

	// Package/component path and other context for debugging.
	Type string `json:"type"`

	// Human-friendly error description.
	Description string `json:"description"`
}

func NewBadRequestError(param string) Error {
	return NewError(CodeBadRequestError, BadRequestError, fmt.Sprintf("Invalid request param '%s'.", param))
}

func NewMethodNotAllowedError() Error {
	return NewError(CodeMethodNotAllowedError, MethodNotAllowedError, "Method is not allowed.")
}

func NewRouteNotFoundError() Error {
	return NewError(CodeRouteNotFoundError, RouteNotFoundError, "Handler is not found.")
}

func NewInternalError() Error {
	return NewError(CodeInternalError, InternalError, "Internal error.")
}

func NewError(code ErrorCode, etype ErrorType, description string) Error {
	return Error{code, etype.String(), description}
}
