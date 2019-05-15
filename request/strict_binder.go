// Copyright 2019 Pavel Petrov <itnelo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package request

import (
	"github.com/gin-gonic/gin"
)

// StrictBinder is a wrapper that starts panic if nested binder is unable
// to bind a Request instance.
type StrictBinder struct {
	nested Binder
}

func (b *StrictBinder) Bind(context *gin.Context) *Request {
	var req *Request = b.nested.Bind(context)

	if nil != req {
		return req
	}

	panic("request: unable to bind request data.")
}

func (b *StrictBinder) SetNested(nested Binder) {
	b.nested = nested
}
