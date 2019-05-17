// Copyright 2019 Pavel Petrov <itnelo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package response

import (
	"bytes"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gopkg.in/go-playground/validator.v8"
)

const (
	expectedResponseValidationErrors = "{\"status\":\"FAIL\",\"errors\":[{\"code\":3, \"description\":\"Invalid request param 'param1'.\", \"type\":\"request.binder.bad_request\"}]}"
)

type TestResponseWriter struct {
	bytes.Buffer
}

func (rw *TestResponseWriter) Header() http.Header {
	return make(http.Header)
}

func (rw *TestResponseWriter) Write(b []byte) (int, error) {
	return rw.Buffer.Write(b)
}

func (rw *TestResponseWriter) WriteHeader(statusCode int) {
}

func TestDefaultErrorDispatcher(t *testing.T) {
	var rw = &TestResponseWriter{}
	var context, _ = gin.CreateTestContext(rw)

	// Case 1: Validation errors.
	var errors = make(validator.ValidationErrors)
	errors["key1"] = &validator.FieldError{Name: "param1"}

	var ed = &DefaultErrorDispatcher{}
	ed.Dispatch(context, errors)

	assert.JSONEq(
		t,
		expectedResponseValidationErrors,
		rw.Buffer.String(),
		"Expecting response with validation errors.",
	)
}
