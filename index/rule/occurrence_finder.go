// Copyright 2019 Pavel Petrov <itnelo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package rule

// Represents a component that can find applicable rules for the given word
// and context marker; second retval will be positive if rules are found
// for the given word and third - a processed word form.
type OccurrenceFinder interface {
	FindApplicableRules(word, contextMarker string) (Rules, bool, string)
}
