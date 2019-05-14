// Copyright 2019 Pavel Petrov <itnelo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
	applog "github.com/symfony-doge/ministry-of-truth-cis/log"
)

// Package-level logger for error events.
var defaultErrorLogger *log.Logger

// Configures and returns application-level recovery middleware for Gin.
func Recovery() gin.HandlerFunc {
	if nil == defaultErrorLogger {
		defaultErrorLoggerInit()
	}

	return jsonRecoveryWith(defaultErrorLogger)
}

func defaultErrorLoggerInit() {
	if errLogger, err := applog.NewErrorLogger("[middleware] "); nil != err {
		log.Println(err)

		panic("middleware: cannot init logger.")
	} else {
		defaultErrorLogger = errLogger
	}
}
