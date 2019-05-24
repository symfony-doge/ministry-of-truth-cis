// Copyright 2019 Pavel Petrov <itnelo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package index

import (
	"log"
	"os"

	"github.com/symfony-doge/ministry-of-truth-cis/tag"
)

// Package-level logger.
var DefaultLogger *log.Logger = log.New(os.Stdout, "[index] ", log.Ldate|log.Ltime|log.Lshortfile)

// Represents sanity index.
type Index struct {
	// Index value (0.0 - 100.0).
	Value float64 `json:"value"`

	// Related tags, with name of tag.Group as the aggregation key.
	Tags map[string]tag.Tags `json:"tags"`
}

func newIndex() *Index {
	return &Index{
		Value: 0.0,
		Tags:  make(map[string]tag.Tags),
	}
}
