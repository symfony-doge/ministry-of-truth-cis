// Copyright 2019 Pavel Petrov <itnelo@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT license
// that can be found in the LICENSE file.

package rule

// Determines which rules have match for a given text
// and returns a result set.
type Processor interface {
	FindMatch(MatchTask) (Rules, error)
}
