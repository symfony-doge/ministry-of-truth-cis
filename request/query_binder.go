// Copyright 2019 Pavel Petrov <itnelo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package request

import (
	"log"

	"github.com/gin-gonic/gin"
)

// Converts URI query parameters into DefaultRequest structure.
type QueryBinder struct {
	logger *log.Logger
}

func (b *QueryBinder) Bind(context *gin.Context) (*DefaultRequest, error) {
	var requestFromQuery DefaultRequest
	var isLoggerProvided = nil != b.logger

	if err := context.ShouldBindQuery(&requestFromQuery); nil != err {
		if isLoggerProvided {
			b.logger.Printf("QueryBinder.Bind: %T %v\n", err, err)
		}

		return nil, err
	}

	if isLoggerProvided && gin.IsDebugging() {
		b.logger.Printf("QueryBinder.Bind: %v\n", requestFromQuery)
	}

	return &requestFromQuery, nil
}

func (b *QueryBinder) SetLogger(logger *log.Logger) {
	b.logger = logger
}
