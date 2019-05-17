// Copyright 2019 Pavel Petrov <itnelo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package rule

import (
	"log"
	"os"
)

// Package-level logger.
var DefaultLogger *log.Logger = log.New(os.Stdout, "[rule] ", log.Ldate|log.Ltime|log.Lshortfile)

type Rules []*Rule

// Rule for sanity index calculation.
type Rule struct {
	// TODO
}
