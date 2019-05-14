// Copyright 2019 Pavel Petrov <itnelo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/symfony-doge/ministry-of-truth-cis/response"
)

// Returns a middleware callback that recovers from any panic
// and writes a valid json response with negative status and error description.
func jsonRecoveryWith(logger *log.Logger) gin.HandlerFunc {
	return func(context *gin.Context) {
		defer func() {
			if err := recover(); nil != err {
				if nil != logger {
					logger.Printf("panic recovered: %s\n", err)
				}

				context.AbortWithStatusJSON(
					http.StatusOK,
					response.NewInternalErrorResponse(),
				)
			}
		}()

		context.Next()
	}
}
