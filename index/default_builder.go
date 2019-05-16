// Copyright 2019 Pavel Petrov <itnelo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package index

type DefaultBuilder struct {
	// TODO
}

func (b *DefaultBuilder) Build(context BuilderContext) *Index {
	// TODO
	// var description = context[BuilderContextDescription]

	return NewIndex()
}

func NewDefaultBuilder() *DefaultBuilder {
	return &DefaultBuilder{}
}
