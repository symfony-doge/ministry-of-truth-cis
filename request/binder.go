// Copyright 2019 Pavel Petrov <itnelo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package request

import (
	"github.com/gin-gonic/gin"
)

// Converts request body into Request structure using gin.Context
// or can start a panic if any error occurred. This method doesn't return error,
// because it is a critical part of request processing and can't be
// normally recovered by a caller.
type Binder interface {
	Bind(context *gin.Context) *Request
}
