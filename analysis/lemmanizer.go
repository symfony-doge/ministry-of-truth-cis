// Copyright 2019 Pavel Petrov <itnelo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package analysis

// Lemmanizer is an interface for components that performs
// semantical analysis for input text and returns a string with
// all words replaced by their lemmas, e.g. for string "динамичная развивающаяся компания"
// the result string will become "динамичный развивающийся компания".
type Lemmanizer interface {
	Lemmanize(string) string
}
