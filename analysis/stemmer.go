// Copyright 2019 Pavel Petrov <itnelo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package analysis

// Stemmer is an interface for components that performs word stemming,
// e.g. for string "Динамичная" the result string will become "динамичн".
type Stemmer interface {
	Stem(string) (string, error)
}
