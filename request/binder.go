// Copyright 2019 Pavel Petrov <itnelo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package request

import (
	"github.com/gin-gonic/gin"
)

// Converts request body into Request structure using gin.Context
type Binder interface {
	Bind(*gin.Context) (*Request, error)
}
