// Copyright 2019 Pavel Petrov <itnelo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package response

// Response status.
type Status uint8

const (
	StatusFail Status = iota
	StatusOk
)

var responseStatuses = [...]string{
	"FAIL",
	"OK",
}

// Implements fmt.Stringer interface.
func (s Status) String() string {
	if s > StatusOk {
		panic("response: undefined status.")
	}

	return responseStatuses[s]
}

// Will be used if no other is available,
// e.g. while recovering from unexpected situations.
type DefaultResponse struct {
	Status string `json:"status"`
	Errors `json:"errors"`
}

// Returns positive response without any errors.
func NewOkResponse() DefaultResponse {
	return NewResponse(StatusOk, Errors{}...)
}

// Returns negative response with internal error.
func NewInternalErrorResponse() DefaultResponse {
	return NewResponse(StatusFail, NewInternalError())
}

// Returns response with specified status and errors.
func NewResponse(status Status, errors ...Error) DefaultResponse {
	return DefaultResponse{status.String(), errors}
}
