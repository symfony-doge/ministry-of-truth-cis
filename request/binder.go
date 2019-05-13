// Copyright 2019 Pavel Petrov <itnelo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package request

import (
	"log"

	"github.com/gin-gonic/gin"
)

// Converts request body into Request structure using gin.Context
type Binder interface {
	Bind(context *gin.Context) *Request
}

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

	panic("request: unable to extract request data.")
}

func (b *StrictBinder) SetNested(nested Binder) {
	b.nested = nested
}

// Converts json from request body into Request structure.
type JSONBinder struct {
	logger *log.Logger
}

func (b *JSONBinder) Bind(context *gin.Context) *Request {
	var requestFromJson Request
	var isLoggerProvided = nil != b.logger

	if err := context.ShouldBindJSON(&requestFromJson); nil != err {
		if isLoggerProvided {
			b.logger.Printf("JSONBinder.Bind: %v\n", err)
		}

		return nil
	}

	if isLoggerProvided && gin.IsDebugging() {
		b.logger.Printf("JSONBinder.Bind: %v\n", requestFromJson)
	}

	return &requestFromJson
}

func (b *JSONBinder) SetLogger(logger *log.Logger) {
	b.logger = logger
}

// Converts URI query parameters into Request structure.
type QueryBinder struct {
	logger *log.Logger
}

func (b *QueryBinder) Bind(context *gin.Context) *Request {
	var requestFromQuery Request
	var isLoggerProvided = nil != b.logger

	if err := context.ShouldBindQuery(&requestFromQuery); nil != err {
		if isLoggerProvided {
			b.logger.Printf("QueryBinder.Bind: %v\n", err)
		}

		return nil
	}

	if isLoggerProvided && gin.IsDebugging() {
		b.logger.Printf("QueryBinder.Bind: %v\n", requestFromQuery)
	}

	return &requestFromQuery
}

func (b *QueryBinder) SetLogger(logger *log.Logger) {
	b.logger = logger
}
