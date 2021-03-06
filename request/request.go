// Copyright 2019 Pavel Petrov <itnelo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package request

import (
	"fmt"
	"log"
	"os"
)

// Package-level logger.
var DefaultLogger *log.Logger = log.New(os.Stdout, "[request] ", log.Ldate|log.Ltime|log.Lshortfile)

type Locale string

// Supported locales.
var localeSupported = map[Locale]bool{
	"ru": true,
}

// Represents base request body of each client request.
type DefaultRequest struct {
	Locale `form:"locale" json:"locale" binding:"required,locale"`
}

// Implements fmt.Stringer interface.
func (r DefaultRequest) String() string {
	return fmt.Sprintf("DefaultRequest(locale='%s')", r.Locale)
}
