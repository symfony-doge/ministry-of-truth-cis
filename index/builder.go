// Copyright 2019 Pavel Petrov <itnelo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package index

type BuilderContext map[string]string

const (
	BuilderContextDescription string = "description"
)

// Builds sanity index for specified context
// according to predefined rule.Rule set.
type Builder interface {
	Build(BuilderContext) *Index
}
