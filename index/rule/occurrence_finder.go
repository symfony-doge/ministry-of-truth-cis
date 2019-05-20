// Copyright 2019 Pavel Petrov <itnelo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package rule

// Represents a component that can find applicable rules for the given word
// It is not performs any specific context validations, etc.;
// just a word occurrence checks.
type OccurrenceFinder interface {
	FindApplicableRules(word, contextMarker string) Rules
}
