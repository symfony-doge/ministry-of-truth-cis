// Copyright 2019 Pavel Petrov <itnelo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package request

import (
	"log"

	"github.com/gin-gonic/gin"
)

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