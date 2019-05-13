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

// Represents expected request body content from client.
type Request struct {
	Locale `form:"locale" json:"locale" binding:"required"`
}

// Implements fmt.Stringer interface.
func (r Request) String() string {
	return fmt.Sprintf("Request{locale: '%s'}", r.Locale)
}
