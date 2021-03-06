// Copyright 2019 Pavel Petrov <itnelo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package response

import (
	"github.com/gin-gonic/gin"
)

// Dispatches various errors during request processing.
type ErrorDispatcher interface {
	Dispatch(*gin.Context, error)
}
