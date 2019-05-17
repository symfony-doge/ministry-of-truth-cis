// Copyright 2019 Pavel Petrov <itnelo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v8"
)

type DefaultErrorDispatcher struct{}

// Dispatches errors and returns error response for the client.
func (ed *DefaultErrorDispatcher) Dispatch(context *gin.Context, err error) {
	// Invalid request.
	if errFields, isValidationError := err.(validator.ValidationErrors); isValidationError {
		var errors Errors

		for key := range errFields {
			errors = append(errors, NewBadRequestError(errFields[key].Name))
		}

		resp := NewBadRequestErrorResponse(errors...)
		context.AbortWithStatusJSON(http.StatusBadRequest, resp)
	}
}

func NewDefaultErrorDispatcher() *DefaultErrorDispatcher {
	return &DefaultErrorDispatcher{}
}
