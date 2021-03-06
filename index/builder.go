// Copyright 2019 Pavel Petrov <itnelo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package index

import (
	"github.com/symfony-doge/ministry-of-truth-cis/request"
)

type BuilderContext map[string]string

// Builds sanity index for specified context
// according to predefined rules set (see "rule" package).
type Builder interface {
	Build(BuilderContext, request.Locale) *Index
}
