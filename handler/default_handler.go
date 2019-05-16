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

// Callback signature for transforming a default request instance
// to the response.
type doFunc func(*request.DefaultRequest) interface{}

// Base handler implementation for HTTP requests.
type defaultHandler struct {
	// Request method to binder mappings.
	// Binder converts body of HTTP request
	// into appropriate request.Request structure.
	binderByMethod map[string]request.Binder

	// For error processing.
	errorDispatcher response.ErrorDispatcher
}

// A shortcut with common request processing logic for cases when
// a default request instance is enough for handler to create a response.
// Calls specified doFunc if default request is successfully binded and validated.
func (h *defaultHandler) handle(do doFunc) gin.HandlerFunc {
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

func (h *defaultHandler) MethodNotAllowed() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(http.StatusOK, response.NewMethodNotAllowedErrorResponse())
	}
}

func (h *defaultHandler) RouteNotFound() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.JSON(http.StatusOK, response.NewRouteNotFoundErrorResponse())
	}
}

func newDefaultHandler() defaultHandler {
	var handler defaultHandler

	// Default request binders.
	var jsonBinder, queryBinder = &request.JSONBinder{}, &request.QueryBinder{}
	jsonBinder.SetLogger(request.DefaultLogger)
	queryBinder.SetLogger(request.DefaultLogger)

	handler.binderByMethod = make(map[string]request.Binder)
	handler.binderByMethod["GET"] = queryBinder
	handler.binderByMethod["POST"] = jsonBinder

	// Default error handler.
	handler.errorDispatcher = response.NewDefaultErrorDispatcher()

	return handler
}
