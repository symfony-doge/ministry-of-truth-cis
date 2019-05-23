// Copyright 2019 Pavel Petrov <itnelo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package analysis

// Purifier is an interface for components that performs word purification.
// This component removes all unwanted symbols to simplify semantic analysis,
// e.g. for string "Дина$,мич:ная;" the result string will become "Динамичная".
type Purifier interface {
	Purify(string) string
}
