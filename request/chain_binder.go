// Copyright 2019 Pavel Petrov <itnelo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package request

import (
	"github.com/gin-gonic/gin"
)

// Converts HTTP request data into Request structure using nested binders.
type ChainBinder struct {
	binders []Binder
}

func (b *ChainBinder) Bind(context *gin.Context) *Request {
	for _, binder := range b.binders {
		if req := binder.Bind(context); nil != req {
			return req
		}
	}

	return nil
}

func (b *ChainBinder) AddBinder(childs ...Binder) {
	for _, child := range childs {
		b.binders = append(b.binders, child)
	}
}
