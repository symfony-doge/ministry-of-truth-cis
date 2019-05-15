// Copyright 2019 Pavel Petrov <itnelo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/symfony-doge/ministry-of-truth-cis/request"
	"github.com/symfony-doge/ministry-of-truth-cis/response"
)

// Callback signature for transforming a request instance
// to the response.
type DoFunc func(*request.Request) interface{}

// Base handler implementation for HTTP requests.
type DefaultHandler struct {
	// Request method to binder mappings.
	// Binder converts body of HTTP request
	// into appropriate request.Request structure.
	binderByMethod map[string]request.Binder

	// For error processing.
	errorDispatcher response.ErrorDispatcher
}

// A shortcut with common request processing logic for simple cases.
// Calls DoFunc if request is successfully binded and validated.
func (h *DefaultHandler) handle(do DoFunc) gin.HandlerFunc {
	return func(context *gin.Context) {
		var binder = h.binderByMethod[context.Request.Method]

		req, bindErr := binder.Bind(context)
		if nil != bindErr {
			h.errorDispatcher.Dispatch(context, bindErr)

			return
		}

		resp := do(req)

		context.JSON(http.StatusOK, resp)
	}
}

func NewDefaultHandler() DefaultHandler {
	var handler DefaultHandler

	// Default request binders.
	var jsonBinder, queryBinder = &request.JSONBinder{}, &request.QueryBinder{}
	jsonBinder.SetLogger(request.DefaultLogger)
	queryBinder.SetLogger(request.DefaultLogger)

	handler.binderByMethod = make(map[string]request.Binder)
	handler.binderByMethod["GET"] = queryBinder
	handler.binderByMethod["POST"] = jsonBinder

	// Default error handler.
	handler.errorDispatcher = &response.DefaultErrorDispatcher{}

	return handler
}
